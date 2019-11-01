package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
)

type ChannelDeleteEvent func(*ChannelDelete)

type ChannelDelete struct {
	*objects.Channel
}

func (cc ChannelDeleteEvent) Type() EventType {
	return CHANNEL_DELETE
}

func (cc ChannelDeleteEvent) New() interface{} {
	return &ChannelDelete{}
}

func (cc ChannelDeleteEvent) Handle(i interface{}) {
	if t, ok := i.(*ChannelDelete); ok {
		cc(t)
	}
}

func (w *Worker) RegisterChannelDeleteEvent() {
	w.EventBus.Register(ChannelDeleteEvent(func(c *ChannelDelete) {
		go func() {
			// Delete channel object
			w.Delete(c.KeyName())

			// Delete channel from channel set on guild object
			w.SetDelete(fmt.Sprintf("cache:guild:%s:Channels", c.GuildId), c.Id)

			// Delete overwrites
			var overwriteIds []string
			for _, overwrite := range c.PermissionsOverwrites {
				overwriteIds = append(overwriteIds, overwrite.Id)
			}
			w.Delete(overwriteIds...)
		}()
	}))
}
