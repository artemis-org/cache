package events

import (
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
)

type GuildRoleUpdateEvent func(*GuildRoleUpdate)

type GuildRoleUpdate struct {
	GuildId string       `json:"guild_id"`
	Role    objects.Role `json:"role"`
}

func (cc GuildRoleUpdateEvent) Type() EventType {
	return GUILD_ROLE_UPDATE
}

func (cc GuildRoleUpdateEvent) New() interface{} {
	return &GuildRoleUpdate{}
}

func (cc GuildRoleUpdateEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildRoleUpdate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildRoleUpdateEvent() {
	w.EventBus.Register(GuildRoleUpdateEvent(func(g *GuildRoleUpdate) {
		go func() {
			fields := make(map[string]map[string]interface{})

			// Create role object
			fields = utils.Append(fields, g.Role.Serialize())

			// ID should already been in role ID set on guild object

			w.Cache(fields)
		}()
	}))
}
