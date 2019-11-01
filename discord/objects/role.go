package objects

import (
	"fmt"
	"github.com/artemis-org/cache/utils"
)

type Role struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Color       int    `json:"color"`
	Hoist       bool   `json:"hoist"`
	Position    int    `json:"position"`
	Permissions int    `json:"permissions"`
	Managed     bool   `json:"managed"`
	Mentionable bool   `json:"mentionable"`
}

func (r *Role) KeyName() string {
	return fmt.Sprintf("cache:role:%s", r.Id)
}

func (r *Role) Serialize() map[string]map[string]interface{} {
	fields := make(map[string]map[string]interface{})
	utils.CopyNonNil(fields, r.KeyName(), r)
	return fields
}
