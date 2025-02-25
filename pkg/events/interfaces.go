package events

import (
	"sync"
	"time"
)

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup)
}

type EventDispatcherInterface interface {
	Register(eventName string, handle EventHandlerInterface) error
	Dispatch(event EventInterface) error
	Remove(eventName string, handle EventInterface) error
	Has(eventName string, handle EventHandlerInterface) bool
	Clear() error
}
