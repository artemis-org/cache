package events

import (
	"fmt"
	"github.com/artemis-org/cache/discord/objects"
)

type VoiceStateUpdateEvent func(*VoiceStateUpdate)

type VoiceStateUpdate struct {
	*objects.VoiceState
}

func (vs VoiceStateUpdateEvent) Type() EventType {
	return VOICE_STATE_UPDATE
}

func (vs VoiceStateUpdateEvent) New() interface{} {
	return &VoiceStateUpdate{}
}

func (vs VoiceStateUpdateEvent) Handle(i interface{}) {
	if t, ok := i.(*VoiceStateUpdate); ok {
		vs(t)
	}
}

func (w *Worker) RegisterVoiceStateUpdateEvent() {
	w.EventBus.Register(VoiceStateUpdateEvent(func(vs *VoiceStateUpdate) {
		go func() {
			arrayKey := fmt.Sprintf("cache:guild:%s:VoiceStates", vs.GuildId)
			if vs.ChannelId == "" { // Left
				// Remove user ID from voice state set on guild object
				w.SetDelete(arrayKey, vs.UserId)
			} else { // Joined or moved
				// Update voice state set on guild object
				w.SetAdd(arrayKey, vs.UserId)
			}

			// Cache voice state object
			w.Cache(vs.Serialize())
		}()
	}))
}
