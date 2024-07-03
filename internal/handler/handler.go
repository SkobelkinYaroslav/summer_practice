package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"summer_practice/internal/domain"
)

type CarService interface {
	CreateCarService(car domain.Car) (domain.Car, error)
	GetAllCarsService() ([]domain.Car, error)
	GetCarByIdService(id int) (domain.Car, error)
	UpdateCarByIdService(car domain.Car) (domain.Car, error)
	PatchCarByIdService(car domain.Car) (domain.Car, error)
	DeleteCarService(id int) error
}

type CarHandler struct {
	Service CarService
}

func NewCarHandler(svc CarService) *gin.Engine {
	handler := &CarHandler{
		Service: svc,
	}

	g := gin.Default()

	g.POST("/cars", handler.CreateCarHandler)
	g.GET("/cars", handler.GetAllCarsHandler)
	g.GET("/cars/:id", handler.GetCarByIdHandler)
	g.PUT("/cars/:id", handler.UpdateCarByIdHandler)
	g.PATCH("/cars/:id", handler.UpdateCarByIdHandler)
	g.DELETE("/cars/:id", handler.DeleteCarHandler)

	return g
}

func (h *CarHandler) CreateCarHandler(c *gin.Context) {
	var car domain.Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	carOutput, err := h.Service.CreateCarService(car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, carOutput)

}

func (h *CarHandler) GetAllCarsHandler(c *gin.Context) {
	cars, err := h.Service.GetAllCarsService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	if len(cars) == 0 {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	c.JSON(http.StatusOK, cars)

}

func (h *CarHandler) GetCarByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	car, err := h.Service.GetCarByIdService(idInt)

	if errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) UpdateCarByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	var car domain.Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	car.ID = idInt

	carOutput, err := h.Service.UpdateCarByIdService(car)

	if errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, carOutput)

}

func (h *CarHandler) PatchCarByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	var car domain.Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	car.ID = idInt

	carOutput, err := h.Service.PatchCarByIdService(car)

	if errors.Is(err, domain.ErrNotFound) {
		c.JSON(http.StatusNotFound, domain.ErrNotFound)
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	c.JSON(http.StatusOK, carOutput)
}
func (h *CarHandler) DeleteCarHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput)
		return
	}

	if err := h.Service.DeleteCarService(idInt); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
