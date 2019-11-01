package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
)

type GuildMemberUpdateEvent func(*GuildMemberUpdate)

type GuildMemberUpdate struct {
	GuildId string       `json:"guild_id"`
	Roles   []string     `json:"roles"`
	User    objects.User `json:"user"`
	Nick    string       `json:"nick"`
}

func (cc GuildMemberUpdateEvent) Type() EventType {
	return GUILD_MEMBER_UPDATE
}

func (cc GuildMemberUpdateEvent) New() interface{} {
	return &GuildMemberUpdate{}
}

func (cc GuildMemberUpdateEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildMemberUpdate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildMemberUpdateEvent() {
	w.EventBus.Register(GuildMemberUpdateEvent(func(g *GuildMemberUpdate) {
		go func() {
			fields := make(map[string]map[string]interface{})

			// Update user object
			fields = utils.Append(fields, g.User.Serialize())

			// Update roles
			rolesKey := fmt.Sprintf("cache:member:%s:%s:Roles", g.GuildId, g.User.Id)
			currentRoles := w.SetGet(rolesKey)

			// Remove any roles the user doesn't have
			var toDelete []string
			for _, r := range currentRoles {
				if !utils.Contains(g.Roles, r) {
					toDelete = append(toDelete, r)
				}
			}
			w.SetDelete(rolesKey, toDelete...)

			// Update roles
			w.SetAdd(rolesKey, g.Roles...)

			// Update nick
			memberKey := fmt.Sprintf("cache:member:%s:%s", g.GuildId, g.User.Id)
			utils.Initialise(fields, memberKey)
			fields[memberKey]["Nick"] = g.Nick

			w.Cache(fields)
		}()
	}))
}
