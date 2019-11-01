package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
)

type GuildDeleteEvent func(*GuildDelete)

type GuildDelete struct {
	*objects.Guild
}

func (cc GuildDeleteEvent) Type() EventType {
	return GUILD_DELETE
}

func (cc GuildDeleteEvent) New() interface{} {
	return &GuildDelete{}
}

func (cc GuildDeleteEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildDelete); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildDeleteEvent() {
	w.EventBus.Register(GuildDeleteEvent(func(g *GuildDelete) {
		go func() {
			// Delete roles
			w.Delete(fmt.Sprintf("cache:guild:%s:Roles", g.Id))
			for _, role := range w.SetGet("cache:guild:%s:Roles") {
				w.Delete(fmt.Sprintf("cache:role:%s", role))
			}

			// Delete emojis
			w.Delete(fmt.Sprintf("cache:guild:%s:Emojis", g.Id))
			for _, emoji := range w.SetGet("cache:guild:%s:Emojis") {
				w.Delete(fmt.Sprintf("cache:emoji:%s", emoji))
			}

			// Delete members
			w.Delete(fmt.Sprintf("cache:guild:%s:Members", g.Id))
			for _, member := range w.SetGet("cache:guild:%s:Members") {
				w.Delete(fmt.Sprintf("cache:member:%s:%s", g.Id, member))
			}

			// Delete voicestates
			w.Delete(fmt.Sprintf("cache:guild:%s:VoiceStates", g.Id))
			for _, emoji := range w.SetGet("cache:guild:%s:VoiceStates") {
				w.Delete(fmt.Sprintf("cache:voicestate:%s", emoji))
			}

			// Delete channels
			for _, channel := range w.SetGet("cache:guild:%s:Channels") {
				// Delete overwrites
				w.Delete(w.SetGet(fmt.Sprintf("cache:channel:%s:PermissionsOverwrites", channel))...)

				// Delete channel object
				w.Delete(fmt.Sprintf("cache:channel:%s", channel))
				w.Delete(fmt.Sprintf("cache:guild:%s:Channels", g.Id))
			}

			// Delete main object
			w.Delete(g.KeyName())
		}()
	}))
}
