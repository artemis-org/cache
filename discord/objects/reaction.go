package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
	"github.com/fatih/structs"
)

type Reaction struct {
	Count int
	Me bool
	Emoji Emoji
}

func (r *Reaction) KeyName(messageId string) string {
	return fmt.Sprintf("cache:reactions:%s:%s", messageId, r.Emoji.Id)
}

func (r *Reaction) Serialize(messageId string) map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})

	for k, v := range structs.Map(r) {
		if k == "Emoji" {
			fields = utils.Append(fields, r.Emoji.Serialize())
			fields[r.KeyName(messageId)][k] = r.Emoji.Id
		} else {
			fields[r.KeyName(messageId)][k] = v
		}
	}

	return fields
}
