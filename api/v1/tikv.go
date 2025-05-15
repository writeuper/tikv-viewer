package v1

import (
	"github.com/writeuper/tikv-viewer/internal/service"
	"github.com/writeuper/tikv-viewer/pkg/response"
)
import "github.com/gin-gonic/gin"

type TikvAPI struct {
	service *service.TiKVService
}

func NewTikvAPI(service *service.TiKVService) *TikvAPI {
	return &TikvAPI{service: service}
}

func (t *TikvAPI) GetValueHandler(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		response.BadRequest(c, "key is empty")
		return
	}
	value, err := t.service.GetValue(c.Request.Context(), []byte(key))
	if err != nil {
		response.BadRequest(c, "get value failed", err)
		return
	}
	response.Success(c, map[string]interface{}{
		"key":   key,
		"value": string(value),
	})
}

func (t *TikvAPI) SetValueHandler(c *gin.Context) {
	var req struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err)
		return
	}
	if err := t.service.SetValue(c.Request.Context(), []byte(req.Key), []byte(req.Value)); err != nil {
		response.InternalServerError(c, "set value failed", err)
		return
	}
	response.Success(c, map[string]interface{}{
		"key":     req.Key,
		"message": "set value success",
	})
}

func (t *TikvAPI) DeleteValueHandler(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		response.BadRequest(c, "key is empty")
		return
	}
	if err := t.service.DeleteValue(c.Request.Context(), []byte(key)); err != nil {
		response.InternalServerError(c, "delete value failed", err)
		return
	}
	response.Success(c, map[string]interface{}{
		"key":     key,
		"message": "delete value success",
	})
}

func (t *TikvAPI) BatchGetHandler(c *gin.Context) {
	var req struct {
		Keys []string `json:"keys" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err)
		return
	}
	keys := make([][]byte, len(req.Keys))
	for i, key := range req.Keys {
		keys[i] = []byte(key)
	}
	values, err := t.service.BatchGet(c.Request.Context(), keys)
	if err != nil {
		response.InternalServerError(c, "batch get failed", err)
		return
	}
	results := make([]map[string]string, len(values))
	for i := range keys {
		results[i] = map[string]string{
			"key":   string(keys[i]),
			"value": string(values[i]),
		}
	}
	response.Success(c, map[string]interface{}{
		"results": results,
		"count":   len(results), // TODO: count should be the number of valid keys in the request
	})
}

func (t *TikvAPI) ScanHandler(c *gin.Context) {
	var req struct {
		StartKey string `json:"start_key" binding:"required"`
		EndKey   string `json:"end_key" binding:"required"`
		Limit    int    `json:"limit"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body", err)
		return
	}
	keys := make([][]byte, req.Limit)
	values := make([][]byte, req.Limit)
	keys, values, err := t.service.Scan(c.Request.Context(), []byte(req.StartKey), []byte(req.EndKey), req.Limit)
	if err != nil {
		response.InternalServerError(c, "scan failed", err)
		return
	}

	results := make([]map[string]string, len(values))
	for i := range keys {
		results[i] = map[string]string{
			"key":   string(keys[i]),
			"value": string(values[i]),
		}
	}
	response.Success(c, map[string]interface{}{
		"results": results,
		"count":   len(results), // TODO: count should be the number of valid keys in the request
	})
}
