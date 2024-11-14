package controllers

import (
	exception "cv-form/exceptions"
	"cv-form/services"
	"cv-form/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

type ProfileController struct {
	Service  services.ProfileService
	Validate *validator.Validate
}

func NewProfileController(service services.ProfileService, validate *validator.Validate) *ProfileController {
	return &ProfileController{Service: service, Validate: validate}
}

func (c *ProfileController) GetProfile(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	profile, err := c.Service.GetProfile(uint(number))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
		return
	}

	ctx.JSON(http.StatusCreated, utils.ProfileResponse(profile))
}

func (c *ProfileController) PostProfile(ctx *gin.Context) {
	var payload utils.PayloadProfile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := c.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	newProfile, err := c.Service.CreateProfile(payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create profile"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"profileCode": newProfile.ProfileCode,
	})
}

func (c *ProfileController) PutProfile(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	var payload utils.PayloadProfile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := c.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	updatedProfile, err := c.Service.UpdateProfile(uint(number), payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update profile"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"profileCode": updatedProfile.ProfileCode,
	})
}
