package main

import (
	"github.com/artemis-org/cache/config"
	"github.com/artemis-org/cache/discord/payloads/events"
	"github.com/artemis-org/cache/redis"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	provider := config.GetConfigProvider()
	provider.LoadConfig()

	redisClient := redis.NewRedisClient()
	redisClient.Connect(redis.CreateRedisURI(config.Conf.RedisUri))

	// Create event bus
	eb := events.NewEventBus()

	// Create worker
	w := events.NewWorker(redisClient, &eb)

	// Event i18n
	w.RegisterChannelCreateEvent()
	w.RegisterChannelUpdateEvent()
	w.RegisterChannelDeleteEvent()
	w.RegisterGuildCreateEvent()
	w.RegisterGuildUpdateEvent()
	w.RegisterGuildDeleteEvent()
	w.RegisterGuildEmojisUpdateEvent()
	w.RegisterGuildMemberAddEvent()
	w.RegisterGuildMemberRemoveEvent()
	w.RegisterGuildMemberUpdateEvent()
	w.RegisterGuildMembersChunkEvent()
	w.RegisterGuildRoleCreateEvent()
	w.RegisterGuildRoleUpdateEvent()
	w.RegisterGuildRoleDeleteEvent()
	w.RegisterVoiceStateUpdateEvent()

	// Listen for new objects
	go w.Listen()

	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, syscall.SIGINT, syscall.SIGKILL)

	<-killChan
}
