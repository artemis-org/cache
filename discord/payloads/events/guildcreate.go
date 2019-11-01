package events

import (
	"github.com/artemis-org/cache/discord/objects"
)

type GuildCreateEvent func(*GuildCreate)

type GuildCreate struct {
	*objects.Guild
}

func (cc GuildCreateEvent) Type() EventType {
	return GUILD_CREATE
}

func (cc GuildCreateEvent) New() interface{} {
	return &GuildCreate{}
}

func (cc GuildCreateEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildCreate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildCreateEvent() {
	w.EventBus.Register(GuildCreateEvent(func(g *GuildCreate) {
		go w.Cache(g.Serialize())
	}))
}
