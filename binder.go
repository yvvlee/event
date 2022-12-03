package event

import (
	"context"
	"errors"
	"fmt"
	"reflect"
)

type Binder interface {
	EventType() reflect.Type
	Handle(context.Context, any) error
}

type SimpleBinder[T any] struct {
	event     T
	listeners []func(context.Context, T)
}

func NewSimpleBinder[T any](event T, listeners ...func(context.Context, T)) *SimpleBinder[T] {
	return &SimpleBinder[T]{event: event, listeners: listeners}
}

func (b *SimpleBinder[T]) EventType() reflect.Type {
	return reflect.TypeOf(b.event)
}

func (b *SimpleBinder[T]) Handle(ctx context.Context, event any) error {
	if t, ok := event.(T); ok {
		for _, f := range b.listeners {
			f(ctx, t)
		}
		return nil
	}
	return errors.New(fmt.Sprintf("unsupported event. expect:%t got:%t", new(T), event))
}

type AsyncBinder[T any] struct {
	event     T
	listeners []func(context.Context, T)
}

func NewAsyncBinder[T any](event T, listeners ...func(context.Context, T)) *AsyncBinder[T] {
	return &AsyncBinder[T]{event: event, listeners: listeners}
}

func (b *AsyncBinder[T]) EventType() reflect.Type {
	return reflect.TypeOf(b.event)
}

func (b *AsyncBinder[T]) Handle(ctx context.Context, event any) error {
	if t, ok := event.(T); ok {
		go func() {
			for _, f := range b.listeners {
				f(ctx, t)
			}
		}()
		return nil
	}
	return errors.New(fmt.Sprintf("unsupported event. expect:%t got:%t", new(T), event))
}
