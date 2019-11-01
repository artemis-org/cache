package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
)

type GuildEmojisUpdateEvent func(*GuildEmojisUpdate)

// No need to completely re-hash the entire guild object
type GuildEmojisUpdate struct {
	GuildId string          `json:"guild_id"`
	Emojis  []objects.Emoji `json:"emoji"`
}

func (cc GuildEmojisUpdateEvent) Type() EventType {
	return GUILD_EMOJIS_UPDATE
}

func (cc GuildEmojisUpdateEvent) New() interface{} {
	return &GuildEmojisUpdate{}
}

func (cc GuildEmojisUpdateEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildEmojisUpdate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildEmojisUpdateEvent() {
	w.EventBus.Register(GuildEmojisUpdateEvent(func(g *GuildEmojisUpdate) {
		go func() {
			// ID array
			var ids []string
			for _, e := range g.Emojis {
				ids = append(ids, e.Id)
			}

			// Delete deleted emojis, we want to delete them in 1 Redis command
			currentEmojis := w.SetGet(fmt.Sprintf("cache:guild:%s:Emojis", g.GuildId))
			var toDelete []string
			for _, e := range currentEmojis {
				if !utils.Contains(ids, e) {
					toDelete = append(toDelete, e)
				}
			}

			// Remove from emoji set on guild object
			w.SetDelete(fmt.Sprintf("cache:guild:%s:Emojis", g.GuildId), toDelete...)

			// Remove actual emoji object
			var deleteKeys []string
			for _, key := range toDelete {
				deleteKeys = append(deleteKeys, fmt.Sprintf("cache:emoji:%s", key))
			}
			w.Delete(deleteKeys...)

			// Update
			fields := make(map[string]map[string]interface{})

			// Set emoji objects
			for _, emoji := range g.Emojis {
				utils.Initialise(fields, emoji.KeyName())
				fields = utils.Append(fields, emoji.Serialize())
			}

			// Add emojis to emoji set on guild objects
			w.SetAdd(fmt.Sprintf("cache:guild:%s:Emojis", g.GuildId), ids...)

			// Update emoji list on guild object

			w.Cache(fields)
		}()
	}))
}
