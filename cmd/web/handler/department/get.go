package department

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Get(c *gin.Context) {
	id := c.Param("id")
	log := zerolog.Ctx(c.Request.Context())
	log.Info().Str("id", id).Send()
	c.JSON(200, id)
}
