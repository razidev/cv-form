package routers

import (
	"cv-form/controllers"
	"cv-form/repository"
	"cv-form/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func employmentRoute(group *gin.RouterGroup) {
	var validator = validator.New()

	employmentRepo := repository.NewEmploymentRepository()
	profileRepo := repository.NewProfileRepository()
	employmentService := services.NewEmploymentService(employmentRepo, profileRepo)
	employmentController := controllers.NewEmploymentController(employmentService, validator)

	group.POST("/:profilecode", employmentController.PostEmployment)
	group.GET("/:profilecode", employmentController.GetEmployment)
	group.DELETE("/:profilecode", employmentController.DeleteEmployment)
}
