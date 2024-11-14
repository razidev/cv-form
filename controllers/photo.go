package controllers

import (
	exception "cv-form/exceptions"
	"cv-form/services"
	"cv-form/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoProfileController struct {
	Service  services.ProfileService
	Validate *validator.Validate
}

func NewPhotoProfileController(service services.ProfileService, validate *validator.Validate) *PhotoProfileController {
	return &PhotoProfileController{Service: service, Validate: validate}
}

func (c *PhotoProfileController) PostPhoto(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	var payload utils.PhotoProfile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := c.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	profile, err := c.Service.UpdateImageProfile(uint(number), payload.Base64img)
	fmt.Println("error updating profile", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to update image profile"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"profileCode": uint(number),
		"photoUrl":    profile,
	})
}

func (c *PhotoProfileController) GetPhoto(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	data, err := c.Service.GetImageProfile(uint(number))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Profile not found"})
		return
	}

	ctx.FileAttachment(data, "downloaded_image.jpg")
}
