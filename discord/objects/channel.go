package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
	"github.com/fatih/structs"
	"reflect"
)

type Channel struct {
	Id                    string      `json:"id"`
	Type                  int         `json:"type"`
	GuildId               string      `json:"guild_id"`
	Position              int         `json:"position"`
	PermissionsOverwrites []Overwrite `json:"permission_overwrites"`
	Name                  string      `json:"name"`
	Topic                 string      `json:"topic"`
	Nsfw                  bool        `json:"nsfw"`
	LastMessageId         string      `json:"last_message_id"`
	Bitrate               int         `json:"bitrate"`
	UserLimit             int         `json:"user_limit"`
	RateLimitPerUser      int         `json:"rate_limit_per_user"`
	Recipients            []User      `json:"recipients"`
	Icon                  string      `json:"icon"`
	OwnerId               string      `json:"owner_id"`
	ApplicationId         string      `json:"application_id"`
	ParentId              string      `json:"parent_id"`
	LastPinTimestamp      string      `json:"last_pin_timestamp"`
}

type ExampleFuckingChannel struct {
	*Channel
}

func (c *Channel) KeyName() string {
	return fmt.Sprintf("cache:channel:%s", c.Id)
}

func (c *Channel) Serialize() map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})
	utils.Initialise(fields, c.KeyName()) // Only 1 key is used, we can run this here

	for k, v := range structs.Map(c) {
		if k == "PermissionsOverwrites" { // Must store arrays separately
			if len(c.PermissionsOverwrites) == 0 { // No point storing an empty field
				continue
			}

			var overwrites []string // Slice of IDs
			for _, overwrite := range c.PermissionsOverwrites {
				fields = utils.Append(fields, overwrite.Serialize()) // Store individual overwrite objects separately
				overwrites = append(overwrites, overwrite.Id)
			}
			fields[c.KeyName()][k] = overwrites
		} else if k == "Recipients" {
			if len(c.Recipients) == 0 { // No point storing an empty field
				continue
			}

			var ids []string
			for _, user := range c.Recipients {
				fields = utils.Append(fields, user.Serialize()) // Store individual user objects separately
				ids = append(ids, user.Id)
			}
			fields[c.KeyName()][k] = ids
		} else {
			if !utils.IsZero(reflect.ValueOf(v)) { // No point storing empty fields
				fields[c.KeyName()][k] = v
			}
		}
	}

	return fields
}
