package events

import (
	"github.com/artemis-org/cache/discord/objects"
)

type ChannelUpdateEvent func(*ChannelUpdate)

type ChannelUpdate struct {
	*objects.Channel
}

func (cc ChannelUpdateEvent) Type() EventType {
	return CHANNEL_UPDATE
}

func (cc ChannelUpdateEvent) New() interface{} {
	return &ChannelUpdate{}
}

func (cc ChannelUpdateEvent) Handle(i interface{}) {
	if t, ok := i.(*ChannelUpdate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterChannelUpdateEvent() {
	w.EventBus.Register(ChannelUpdateEvent(func(c *ChannelUpdate) {
		go w.Cache(c.Serialize())
	}))
}
