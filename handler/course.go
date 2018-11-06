package handler

import (
	"net/http"

	"github.com/lucasfloriani/go-mongo/helper"
	"github.com/lucasfloriani/go-mongo/model"

	"github.com/labstack/echo"
)

type (
	// courseService specifies the interface for the course service needed by courseResource.
	courseService interface {
		Get(id string) (*model.Course, error)
		Query(offset, limit int) ([]model.Course, error)
		Count() (int, error)
		Create(model *model.Course) (*model.Course, error)
		Update(model *model.Course) (*model.Course, error)
		Delete(id string) (*model.Course, error)
	}

	// courseResource defines the handlers for the CRUD APIs.
	courseResource struct {
		service courseService
	}
)

// ServeCourseResource sets up the routing of course endpoints and the corresponding handlers (routes)
func ServeCourseResource(e *echo.Group, service courseService) {
	at := &courseResource{service}
	courseGroup := e.Group("/course")
	{
		courseGroup.GET("/:courseID", at.get)
		courseGroup.GET("/", at.query)
		courseGroup.POST("/", at.create)
		courseGroup.PUT("/:courseID", at.update)
		courseGroup.DELETE("/:courseID", at.delete)
	}
}

// get verify rest params, call service method to execute business logic
// and return JSON data
func (r *courseResource) get(c echo.Context) error {
	response, err := r.service.Get(c.Param("courseID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.NewErrorResponse(err))
	}
	return c.JSON(http.StatusFound, helper.NewSuccessResponse(*response))
}

// query verify rest params, call service method to execute business logic
// and return JSON data
func (r *courseResource) query(c echo.Context) error {
	count, err := r.service.Count()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	paginatedList := helper.GetPaginatedListFromRequest(c, count)
	items, err := r.service.Query(paginatedList.Offset(), paginatedList.Limit())
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}
	paginatedList.Items = items

	return c.JSON(http.StatusFound, helper.NewSuccessResponse(paginatedList))
}

// create call service method to execute business logic
// and return JSON data
func (r *courseResource) create(c echo.Context) error {
	var model model.Course
	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}
	response, err := r.service.Create(&model)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	return c.JSON(http.StatusCreated, helper.NewSuccessResponse(*response))
}

// update verify rest params, call service method to execute business logic
// and return JSON data
func (r *courseResource) update(c echo.Context) error {
	model, err := r.service.Get(c.Param("courseID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	response, err := r.service.Update(model)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	return c.JSON(http.StatusOK, helper.NewSuccessResponse(*response))
}

// delete verify rest params, call service method to execute business logic
// and return JSON data
func (r *courseResource) delete(c echo.Context) error {
	response, err := r.service.Delete(c.Param("courseID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	return c.JSON(http.StatusOK, helper.NewSuccessResponse(*response))
}
