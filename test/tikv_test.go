package test

import (
	"context"
	"fmt"
	"github.com/writeuper/tikv-viewer/internal/repository/tikv"
	"github.com/writeuper/tikv-viewer/internal/service"
	"log"
	"testing"
	"time"
)

func TestTikv(t *testing.T) {
	client, err := tikv.NewTiKVClient([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatalf("failed to connect to tikv: %v", err)
		return
	}
	tikvService := service.NewTiKVService(client)
	keyPrefix := fmt.Sprintf("%s-%d", t.Name(), time.Now().Nanosecond())
	genScanKeyValue := func(n int) ([]byte, []byte) {
		return []byte(fmt.Sprintf("%s-%d", keyPrefix, n)), []byte(fmt.Sprintf("%d", n))
	}
	startIndex := 100
	endIndex := 200
	for i := startIndex; i < endIndex; i++ {
		key, value := genScanKeyValue(i)
		err = tikvService.SetValue(context.Background(), key, value)
		if err != nil {
			t.Fatalf("failed to put value: %v", err)
		}
	}
	startKey := []byte(keyPrefix)
	endKey, _ := genScanKeyValue(endIndex)
	keys, values, err := tikvService.Scan(context.Background(), startKey, endKey, 100)
	if err != nil {
		return
	}
	for i := 0; i < len(keys); i++ {
		fmt.Printf("key: %s, value: %s\n", keys[i], values[i])
	}

}
