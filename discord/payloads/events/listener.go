package events

import (
	"encoding/json"
	"fmt"
	"github.com/apex/log"
)

func (w *Worker) Listen() {
	for {
		data := w.RedisClient.BLPop(0, "cache:update")
		res, err := data.Result(); if err != nil {
			log.Warn(err.Error())
			continue
		}

		if len(res) > 1 {
			raw := res[1]

			var p Payload
			if err = json.Unmarshal([]byte(raw), &p); err != nil {
				continue
			}

			var event Event
			if ev := w.EventBus.GetEventByName(p.Name); ev == nil {
				continue
			} else {
				event = *ev
			}

			p.Data = event.New()

			if err = json.Unmarshal(p.Raw, p.Data); err != nil {
				fmt.Println(err.Error())
				continue
			}

			go event.Handle(p.Data)
		}
	}
}
