package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
	"github.com/artemis-org/cache/utils"
	"strconv"
)

type GuildMemberAddEvent func(*GuildMemberAdd)

type GuildMemberAdd struct {
	GuildId string          `json:"guild_id"`
	*objects.Member
}

func (cc GuildMemberAddEvent) Type() EventType {
	return GUILD_MEMBER_ADD
}

func (cc GuildMemberAddEvent) New() interface{} {
	return &GuildMemberAdd{}
}

func (cc GuildMemberAddEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildMemberAdd); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildMemberAddEvent() {
	w.EventBus.Register(GuildMemberAddEvent(func(g *GuildMemberAdd) {
		go func() {
			// Add member object
			fields := make(map[string]map[string]interface{})

			keyName := fmt.Sprintf("cache:guild:%s", g.GuildId)
			utils.Initialise(fields, keyName)

			fields = utils.Append(fields, g.Member.Serialize(g.GuildId))

			// Add user to member set on guild object
			w.SetAdd(fmt.Sprintf("cache:guild:%s:Members", g.GuildId), g.User.Id)

			if memberCount, err := strconv.Atoi(w.GetHash(keyName, "MemberCount")); err == nil {
				fields[keyName]["MemberCount"] = memberCount + 1
			}

			w.Cache(fields)
		}()
	}))
}
