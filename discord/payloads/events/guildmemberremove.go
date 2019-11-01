package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
	"strconv"
)

type GuildMemberRemoveEvent func(*GuildMemberRemove)

type GuildMemberRemove struct {
	GuildId string          `json:"guild_id"`
	*objects.User
}

func (cc GuildMemberRemoveEvent) Type() EventType {
	return GUILD_MEMBER_REMOVE
}

func (cc GuildMemberRemoveEvent) New() interface{} {
	return &GuildMemberRemove{}
}

func (cc GuildMemberRemoveEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildMemberRemove); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildMemberRemoveEvent() {
	w.EventBus.Register(GuildMemberRemoveEvent(func(g *GuildMemberRemove) {
		go func() {
			fields := make(map[string]map[string]interface{})

			guildKey := fmt.Sprintf("cache:guild:%s", g.GuildId)
			utils.Initialise(fields, guildKey)

			// Don't delete user object as it may be used in other guilds

			// Update guild object
			if memberCount, err := strconv.Atoi(w.GetHash(guildKey, "MemberCount")); err == nil {
				fields[guildKey]["MemberCount"] = memberCount - 1
			}

			// Remove member from member set on guild object
			w.SetDelete(fmt.Sprintf("cache:guild:%s:Members", g.GuildId), g.Id)

			// Delete member object
			w.Delete(fmt.Sprintf("cache:member:%s:%s", g.GuildId, g.Id))

			w.Cache(fields)
		}()
	}))
}
