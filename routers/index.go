package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")

	profileGroup := apiGroup.Group("/profile")
	profileRoute(profileGroup)

	employmentGroup := apiGroup.Group("/employment")
	employmentRoute(employmentGroup)

	photoGroup := apiGroup.Group("/photo")
	photoRoute(photoGroup)

	return r
}
