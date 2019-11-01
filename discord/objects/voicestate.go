package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
	"github.com/fatih/structs"
	"reflect"
)

type VoiceState struct {
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
	UserId    string `json:"user_id"`
	Member    Member `json:"member"`
	SessionId string `json:"session_id"`
	Deaf      bool   `json:"deaf"`
	Mute      bool   `json:"mute"`
	SelfDeaf  bool   `json:"self_deaf"`
	SelfMute  bool   `json:"self_mute"`
	Suppress  bool   `json:"suppress"`
}

func (vs *VoiceState) KeyName() string {
	return fmt.Sprintf("cache:voicestate:%s:%s", vs.GuildId, vs.UserId)
}

func (vs *VoiceState) Serialize() map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})
	utils.Initialise(fields, vs.KeyName())

	for k, v := range structs.Map(vs) {
		if k == "Member" {
			m := v.(Member)
			fields = utils.Append(fields, m.Serialize(vs.GuildId))
			// We already have the user ID + guild ID so no need to cache anything for member
		} else {
			if !utils.IsZero(reflect.ValueOf(v)) {
				fields[vs.KeyName()][k] = v
			}
		}
	}

	return fields
}
