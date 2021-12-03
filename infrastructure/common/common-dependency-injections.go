package common_di

import (
	"sync"

	"github.com/rozturac/rmqc"
	"golang.clean.architecture/api/configs"
	"golang.clean.architecture/domain/common"
	event_dispatcher "golang.clean.architecture/infrastructure/common/event-dispatcher"
)

var (
	once = sync.Once{}
	rbt  *rmqc.RabbitMQ
)

func NewEventHandlerResolve(rbt *rmqc.RabbitMQ) common.IEventDispatcher {
	return event_dispatcher.NewRabbitMQEventDispatcher(rbt)
}

func NewRabbitMQResolve(config configs.Config) *rmqc.RabbitMQ {
	var err error

	once.Do(func() {
		rbt, err = rmqc.Connect(rmqc.RabbitMQConfig{
			Host:           config.RabbitMQ.Host,
			Username:       config.RabbitMQ.Username,
			Password:       config.RabbitMQ.Password,
			Port:           config.RabbitMQ.Port,
			VHost:          config.RabbitMQ.VHost,
			ConnectionName: config.RabbitMQ.ConnectionName,
			Reconnect: rmqc.Reconnect{
				MaxAttempt: config.RabbitMQ.Reconnect.MaxAttempt,
				Interval:   config.RabbitMQ.Reconnect.Interval,
			},
		})
	})

	if err != nil {
		panic(err)
	}

	return rbt
}
