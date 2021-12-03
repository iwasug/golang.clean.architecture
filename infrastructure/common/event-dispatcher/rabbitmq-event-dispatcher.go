package event_dispatcher

import (
	"reflect"

	"github.com/rozturac/rmqc"
	"golang.clean.architecture/application/consts"
	"golang.clean.architecture/domain/common"
)

type RabbitMQEventDispatcher struct {
	rbt     *rmqc.RabbitMQ
	appName string
}

func NewRabbitMQEventDispatcher(rbt *rmqc.RabbitMQ) common.IEventDispatcher {
	return &RabbitMQEventDispatcher{rbt: rbt}
}

func (handler RabbitMQEventDispatcher) Dispatch(events []common.IBaseEvent) {
	for _, event := range events {
		t := reflect.TypeOf(event)
		eventName := t.Elem().Name()
		handler.rbt.Publish(consts.AppName, eventName, event)
	}
}
