package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
)

type GuildRoleCreateEvent func(*GuildRoleCreate)

type GuildRoleCreate struct {
	GuildId string       `json:"guild_id"`
	Role    objects.Role `json:"role"`
}

func (cc GuildRoleCreateEvent) Type() EventType {
	return GUILD_ROLE_CREATE
}

func (cc GuildRoleCreateEvent) New() interface{} {
	return &GuildRoleCreate{}
}

func (cc GuildRoleCreateEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildRoleCreate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildRoleCreateEvent() {
	w.EventBus.Register(GuildRoleCreateEvent(func(g *GuildRoleCreate) {
		go func() {
			fields := make(map[string]map[string]interface{})

			// Create role object
			fields = utils.Append(fields, g.Role.Serialize())

			// Append role ID to set on guild object
			w.SetAdd(fmt.Sprintf("cache:guild:%s:Roles", g.GuildId), g.Role.Id)

			w.Cache(fields)
		}()
	}))
}
