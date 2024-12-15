package sdk

import (
	"context"
	"testing"
	"time"
)

func TestNewIdGenClient(t *testing.T) {
	client := NewIdGenClient(WithBizTagName(""))
	for i := 0; i < 1000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		_, err := client.NextId(ctx)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestNewIdGenClientWithTagName(t *testing.T) {
	client := NewIdGenClient(WithBizTagName("tag3"))
	for i := 0; i < 1000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		_, err := client.NextId(ctx)
		cancel()
		if err != nil {
			t.Fatal(err)
		}
	}
	time.Sleep(30 * time.Second)

}
