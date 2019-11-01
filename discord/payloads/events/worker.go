package events

import (
	"fmt"
	"github.com/apex/log"
	"github.com/artemis-org/cache/redis"
	"github.com/artemis-org/cache/utils"
	"time"
)

type Worker struct {
	RedisClient *redis.RedisClient
	EventBus *EventBus
}

func NewWorker(client *redis.RedisClient, eb *EventBus) Worker {
	return Worker{
		RedisClient: client,
		EventBus: eb,
	}
}

func (w *Worker) Cache(fields map[string]map[string]interface{}) {
	hash := make(map[string]map[string]interface{})

	for k, v := range fields {
		utils.Initialise(hash, k)
		for k1, v1 := range v {
			if arr, ok := v1.([]string); ok {
				w.RedisClient.SAdd(fmt.Sprintf("%s:%s", k, k1), arr)
			} else {
				hash[k][k1] = v1
			}
		}
	}

	for k, v := range hash {
		_, _ = w.RedisClient.HMSet(k, v).Result()
	}
}

func (w *Worker) CacheField(key string, field string, value interface{}) {
	w.RedisClient.HSet(key, field, value)
}

func (w *Worker) Delete(key ...string) {
	w.RedisClient.Del(key...)
}

func (w *Worker) SetDelete(key string, elem ...string) {
	w.RedisClient.SRem(key, elem)
}

func (w *Worker) SetAdd(key string, elem ...string) {
	w.RedisClient.SAdd(key, elem)
}

func (w *Worker) SetGet(key string) []string {
	if res, err := w.RedisClient.SMembers(key).Result(); err == nil {
		return res
	}
	return make([]string, 0)
}

func (w *Worker) DeletePattern(pattern string) {
	if keys, err := w.RedisClient.Keys(pattern).Result(); err == nil {
		w.RedisClient.Del(keys...)
	}
}

func (w *Worker) SetExpiry(key string, seconds int) {
	w.RedisClient.Expire(key, time.Duration(seconds) * time.Second)
}

func (w *Worker) GetHashMap(key string) map[string]map[string]string {
	fields := make(map[string]map[string]string)
	fields[key] = make(map[string]string)

	array, err := w.RedisClient.HGetAll(key).Result(); if err != nil {
		log.Error(err.Error())
		return fields
	}

	for k, v := range array {
		fields[key][k] = v
	}

	return fields
}

func (w *Worker) GetHash(key string, field string) string {
	res, err := w.RedisClient.HGet(key, field).Result(); if err != nil {
		return ""
	}

	return res
}
