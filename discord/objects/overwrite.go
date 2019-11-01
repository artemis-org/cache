package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
)

type Overwrite struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Allow int `json:"allow"`
	Deny  int `json:"deny"`
}

func (o *Overwrite) KeyName() string {
	return fmt.Sprintf("cache:overwrite:%s", o.Id)
}

func (o *Overwrite) Serialize() map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})
	utils.CopyNonNil(fields, o.KeyName(), o)
	return fields
}
