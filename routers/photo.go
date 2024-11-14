package routers

import (
	"cv-form/controllers"
	"cv-form/repository"
	"cv-form/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func photoRoute(group *gin.RouterGroup) {
	var validator = validator.New()

	profileRepo := repository.NewProfileRepository()
	profileService := services.NewProfileService(profileRepo)
	photoController := controllers.NewPhotoProfileController(profileService, validator)

	group.PUT("/:profilecode", photoController.PostPhoto)
	group.GET("/:profilecode", photoController.GetPhoto)
}
