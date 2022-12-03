package event

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	manager := NewManager()
	manager.Register(
		NewSimpleBinder(new(TestEvent), MyEventListener),
		NewAsyncBinder(new(TestEvent), MyAsyncEventListener),
	)
	err := manager.Trigger(context.Background(), &TestEvent{
		Name: "test",
	})
	t.Log("trigger done")
	t.Log(err)
	err = manager.Trigger(context.Background(), TestEvent{
		Name: "test",
	})
	t.Log(err)
	time.Sleep(10 * time.Second)
	t.Log("test done")
}

type TestEvent struct {
	Name string
}

func MyEventListener(ctx context.Context, e *TestEvent) {
	fmt.Println(e)
}

func MyAsyncEventListener(ctx context.Context, e *TestEvent) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}
