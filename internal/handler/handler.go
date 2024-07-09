package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"summer_practice/internal/domain"
)

type CarService interface {
	CreateCarService(car domain.Car) (domain.Car, error)
	GetAllCarsService() ([]domain.Car, error)
	GetCarByIdService(id int) (domain.Car, error)
	PutCarByIdService(car domain.Car) (domain.Car, error)
	PatchCarByIdService(fieldsToUpdate map[string]interface{}) (domain.Car, error)
	DeleteCarService(id int) error
}

type CarHandler struct {
	Service CarService
}

func New(svc CarService) *gin.Engine {
	handler := &CarHandler{
		Service: svc,
	}

	g := gin.Default()

	cars := g.Group("/cars")
	{
		cars.POST("/", handler.CreateCarHandler)
		cars.GET("/", handler.GetAllCarsHandler)
		cars.GET("/:id", handler.GetCarByIdHandler)
		cars.PUT("/:id", handler.PutCarByIdHandler)
		cars.PATCH("/:id", handler.PatchCarByIdHandler)
		cars.DELETE("/:id", handler.DeleteCarHandler)
	}

	return g
}

func (h *CarHandler) CreateCarHandler(c *gin.Context) {
	var car domain.Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": domain.ErrBadParamInput.Error()})
		return
	}

	carOutput, err := h.Service.CreateCarService(car)

	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, carOutput)

}

func (h *CarHandler) GetAllCarsHandler(c *gin.Context) {
	cars, err := h.Service.GetAllCarsService()

	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cars)

}

func (h *CarHandler) GetCarByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(getStatusCode(domain.ErrBadParamInput), gin.H{"error": domain.ErrBadParamInput.Error()})
		return
	}

	car, err := h.Service.GetCarByIdService(idInt)

	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (h *CarHandler) PutCarByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(getStatusCode(domain.ErrBadParamInput), gin.H{"error": domain.ErrBadParamInput.Error()})
		return
	}

	var car domain.Car

	err = c.BindJSON(&car)
	if err != nil {
		c.JSON(getStatusCode(domain.ErrBadParamInput), gin.H{"error": domain.ErrBadParamInput.Error()})
		return
	}

	car.ID = idInt

	carOutput, err := h.Service.PutCarByIdService(car)

	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, carOutput)

}

func (h *CarHandler) PatchCarByIdHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(getStatusCode(domain.ErrBadParamInput), gin.H{"error": domain.ErrBadParamInput.Error()})
		return
	}

	var fieldsToUpdate map[string]interface{}

	err = c.BindJSON(&fieldsToUpdate)
	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	fieldsToUpdate["id"] = idInt

	carOutput, err := h.Service.PatchCarByIdService(fieldsToUpdate)

	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, carOutput)
}

func (h *CarHandler) DeleteCarHandler(c *gin.Context) {
	idStr := c.Param("id")
	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(getStatusCode(domain.ErrBadParamInput), gin.H{"error": domain.ErrBadParamInput.Error()})
		return
	}
	err = h.Service.DeleteCarService(idInt)

	if err != nil {
		c.JSON(getStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
