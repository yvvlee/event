package event

import (
	"context"
	"errors"
	"reflect"
)

type Manager struct {
	binders map[reflect.Type][]Binder
}

func NewManager() *Manager {
	return &Manager{
		binders: make(map[reflect.Type][]Binder),
	}
}

func (m *Manager) Register(binders ...Binder) {
	for _, binder := range binders {
		m.binders[binder.EventType()] = append(m.binders[binder.EventType()], binder)
	}
}

func (m *Manager) Trigger(ctx context.Context, event any) error {
	eventType := reflect.TypeOf(event)
	if b, ok := m.binders[eventType]; ok {
		for _, binder := range b {
			if err := binder.Handle(ctx, event); err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("unknown event:" + eventType.Name())
}
