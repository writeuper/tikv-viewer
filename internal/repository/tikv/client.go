package tikv

import (
	"context"
	"github.com/tikv/client-go/v2/config"
	tikverr "github.com/tikv/client-go/v2/error"
	"github.com/tikv/client-go/v2/rawkv"
	"github.com/writeuper/tikv-viewer/pkg/logger"
)

type Client struct {
	client *rawkv.Client
}

// NewTiKVClient 创建tikv客户端
func NewTiKVClient(pdAddrs []string) (*Client, error) {
	l := logger.GetLogger()
	l.Infof("Connecting to TiKV cluster: %v", pdAddrs)

	client, err := rawkv.NewClient(context.Background(), pdAddrs, config.Security{})
	if err != nil {
		return nil, err
	}
	l.Infof("Connected to TiKV cluster: %v", pdAddrs)
	return &Client{client: client}, nil
}

func (c *Client) Put(ctx context.Context, key []byte, value []byte) error {
	// 不允许设置为 nil 或者空字节，统一 row 和 txn 的表现形式
	if value == nil || len(value) == 0 {
		return tikverr.ErrCannotSetNilValue
	}
	return c.client.Put(ctx, key, value)
}

func (c *Client) Get(ctx context.Context, key []byte) ([]byte, error) {
	return c.client.Get(ctx, key)
}

func (c *Client) BatchGet(ctx context.Context, keys [][]byte) ([][]byte, error) {
	return c.client.BatchGet(ctx, keys)
}

func (c *Client) BatchPut(ctx context.Context, keys [][]byte, value [][]byte) error {
	return c.client.BatchPut(ctx, keys, value)
}

func (c *Client) Delete(ctx context.Context, key []byte) error {
	return c.client.Delete(ctx, key)
}

func (c *Client) Close() error {
	return c.client.Close()
}
