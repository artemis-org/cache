package events

import (
	"github.com/artemis-org/cache/discord/objects"
)

type GuildUpdateEvent func(*GuildUpdate)

type GuildUpdate struct {
	*objects.Guild
}

func (cc GuildUpdateEvent) Type() EventType {
	return GUILD_UPDATE
}

func (cc GuildUpdateEvent) New() interface{} {
	return &GuildUpdate{}
}

func (cc GuildUpdateEvent) Handle(i interface{}) {
	if t, ok := i.(*GuildUpdate); ok {
		cc(t)
	}
}

func (w *Worker) RegisterGuildUpdateEvent() {
	w.EventBus.Register(GuildUpdateEvent(func(g *GuildUpdate) {
		go w.Cache(g.Serialize())
	}))
}
