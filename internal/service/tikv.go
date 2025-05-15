package service

import (
	"context"
	"github.com/writeuper/tikv-viewer/internal/repository/tikv"
	"github.com/writeuper/tikv-viewer/pkg/logger"
)

type TiKVService struct {
	repo *tikv.Client
}

func NewTiKVService(repo *tikv.Client) *TiKVService {
	return &TiKVService{
		repo: repo,
	}
}

func (s *TiKVService) GetValue(ctx context.Context, key []byte) ([]byte, error) {
	l := logger.GetLogger()
	l.Debugf("GetValue for key: %s", string(key))
	return s.repo.Get(ctx, key)
}

func (s *TiKVService) SetValue(ctx context.Context, key []byte, value []byte) error {
	l := logger.GetLogger()
	l.Debugf("SetValue for key: %s", string(key))
	return s.repo.Put(ctx, key, value)
}

func (s *TiKVService) DeleteValue(ctx context.Context, key []byte) error {
	l := logger.GetLogger()
	l.Debugf("DeleteValue for key: %s", string(key))
	return s.repo.Delete(ctx, key)
}

func (s *TiKVService) BatchGet(ctx context.Context, keys [][]byte) ([][]byte, error) {
	l := logger.GetLogger()
	l.Debugf("BatchGet for keys: %v", keys)
	return s.repo.BatchGet(ctx, keys)
}

func (s *TiKVService) BatchPut(ctx context.Context, keys [][]byte, value [][]byte) error {
	l := logger.GetLogger()
	l.Debugf("BatchPut for keys: %v", keys)
	return s.repo.BatchPut(ctx, keys, value)
}

func (s *TiKVService) Scan(ctx context.Context, prefix []byte, endKey []byte, limit int) ([][]byte, [][]byte, error) {
	l := logger.GetLogger()
	l.Debugf("Scan for prefix: %s", string(prefix))
	return s.repo.Scan(ctx, prefix, endKey, limit)
}

func (s *TiKVService) Close() {}
