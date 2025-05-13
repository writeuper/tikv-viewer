package v1

import "github.com/writeuper/tikv-viewer/internal/service"
import "github.com/gin-gonic/gin"

type TikvAPI struct {
	service service.TiKVService
}

func NewTikvAPI(service service.TiKVService) *TikvAPI {
	return &TikvAPI{service: service}
}

func (t *TikvAPI) GetValueHandler(c *gin.Context) {
	key := c.Param("key")
	if key == "" {

	}
}
