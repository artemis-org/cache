package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
)

type GuildRoleDeleteEvent func(*GuildRoleDelete)

type GuildRoleDelete struct {
	GuildId string       `json:"guild_id"`
	Role    objects.Role `json:"role"`
}

func (cc GuildRoleDeleteEvent) Type() EventType {
	return GUILD_ROLE_DELETE
}

func (cc GuildRoleDeleteEvent) New() interface{} {
	return &GuildRoleDelete{}
}

func (cc GuildRoleDeleteEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildRoleDelete); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildRoleDeleteEvent() {
	w.EventBus.Register(GuildRoleDeleteEvent(func(g *GuildRoleDelete) {
		go func() {
			/*
			To save getting thousands of members and deleting the role from them individually,
			workers should remove invalid roles once encountered.
			 */

			// Delete role object
			w.Delete(fmt.Sprintf("cache:role:%s", g.Role.Id))

			// Delete role from role ID set on guild object
			w.SetDelete(fmt.Sprintf("cache:guild:%s:Roles", g.GuildId), g.Role.Id)
		}()
	}))
}
