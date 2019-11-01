package events

import (
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
)

type GuildMembersChunkEvent func(*GuildMembersChunk)

type GuildMembersChunk struct {
	GuildId string           `json:"guild_id"`
	Members []objects.Member `json:"members"`
}

func (cc GuildMembersChunkEvent) Type() EventType {
	return GUILD_MEMBERS_CHUNK
}

func (cc GuildMembersChunkEvent) New() interface{} {
	return &GuildMembersChunk{}
}

func (cc GuildMembersChunkEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildMembersChunk); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildMembersChunkEvent() {
	w.EventBus.Register(GuildMembersChunkEvent(func(g *GuildMembersChunk) {
		go func() {
			fields := make(map[string]map[string]interface{})

			for _, m := range g.Members {
				fields = utils.Append(fields, m.Serialize(g.GuildId))
			}

			w.Cache(fields)
		}()
	}))
}
