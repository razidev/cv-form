package controllers

import (
	exception "cv-form/exceptions"
	"cv-form/services"
	"fmt"
	"strconv"
	"time"

	"cv-form/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmploymentController struct {
	Service  services.EmploymentService
	Validate *validator.Validate
}

func NewEmploymentController(service services.EmploymentService, validate *validator.Validate) *EmploymentController {
	return &EmploymentController{Service: service, Validate: validate}
}

func isValidDate(date string) bool {
	// Layout untuk parsing: harus sesuai format `dd-mm-yyyy`
	layout := "02-01-2006"
	_, err := time.Parse(layout, date)
	return err == nil
}

func (c *EmploymentController) PostEmployment(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	var payload utils.PayloadEmployment
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid format request"})
		return
	}

	if err := c.Validate.Struct(payload); err != nil {
		errorsMap := exception.ValidationError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": errorsMap})
		return
	}

	if !isValidDate(payload.StartDate) || !isValidDate(payload.EndDate) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "date not valid"})
		return
	}

	employment, err := c.Service.CreateEmployment(uint(number), payload)
	fmt.Println("err", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create employment"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":          employment.ID,
		"profileCode": uint(number),
	})
}

func (c *EmploymentController) GetEmployment(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	employment, err := c.Service.ListEmployments(uint(number))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to list employments"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": utils.EmploymentResponses(employment),
	})
}

func (c *EmploymentController) DeleteEmployment(ctx *gin.Context) {
	profileCode := ctx.Param("profilecode")
	number, _ := strconv.ParseUint(profileCode, 10, 32)

	id, errId := strconv.Atoi(ctx.Query("id"))
	if errId != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id number"})
		return
	}

	err := c.Service.DeleteEmployment(uint(id), uint(number))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to Delete employments"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"profileCode": uint(number),
	})
}
