package routers

import (
	"cv-form/controllers"
	"cv-form/repository"
	"cv-form/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func profileRoute(group *gin.RouterGroup) {
	var validator = validator.New()

	profileRepo := repository.NewProfileRepository()
	profileService := services.NewProfileService(profileRepo)
	profileController := controllers.NewProfileController(profileService, validator)

	group.GET("/:profilecode", profileController.GetProfile)
	group.POST("/", profileController.PostProfile)
	group.PUT("/:profilecode", profileController.PutProfile)
}
