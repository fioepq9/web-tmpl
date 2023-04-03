package department

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.GET("/:id", Get)
}
