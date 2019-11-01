package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
	"github.com/fatih/structs"
	"reflect"
)

type Emoji struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Roles         []string `json:"roles"`
	User          User     `json:"user"`
	RequireColons bool     `json:"require_colons"`
	Managed       bool     `json:"managed"`
	Animated      bool     `json:"animated"`
}

func (e *Emoji) KeyName() string {
	return fmt.Sprintf("cache:emoji:%s", e.Id)
}

func (e *Emoji) Serialize() map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})
	utils.Initialise(fields, e.KeyName())

	for k, v := range structs.Map(e) {
		if k == "Roles" {
			fields[e.KeyName()][k] = utils.JoinToString(e.Roles, ",")
		} else if k == "User" {
			if !utils.IsZero(reflect.ValueOf(e.User)) {
				fields = utils.Append(fields, e.User.Serialize())
				fields[e.KeyName()][k] = e.User.Id
			}
		} else {
			if !utils.IsZero(reflect.ValueOf(k)) {
				fields[e.KeyName()][k] = v
			}
		}
	}

	return fields
}
