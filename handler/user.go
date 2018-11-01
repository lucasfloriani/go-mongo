package handler

import (
	"net/http"

	"github.com/lucasfloriani/go-mongo/helper"
	"github.com/lucasfloriani/go-mongo/model"

	"github.com/labstack/echo"
)

type (
	// userService specifies the interface for the user service needed by userResource.
	userService interface {
		Get(id string) (*model.User, error)
		Query(offset, limit int) ([]model.User, error)
		Count() (int, error)
		Create(model *model.User) (*model.User, error)
		Update(model *model.User) (*model.User, error)
		Delete(id string) (*model.User, error)
	}

	// userResource defines the handlers for the CRUD APIs.
	userResource struct {
		service userService
	}
)

// ServeUserResource sets up the routing of user endpoints and the corresponding handlers (routes)
func ServeUserResource(e *echo.Group, service userService) {
	at := &userResource{service}
	userGroup := e.Group("/user")
	{
		userGroup.GET("/:userID", at.get)
		userGroup.GET("/", at.query)
		userGroup.POST("/", at.create)
		userGroup.PUT("/:userID", at.update)
		userGroup.DELETE("/:userID", at.delete)
	}
}

// get verify rest params, call service method to execute business logic
// and return JSON data
func (r *userResource) get(c echo.Context) error {
	response, err := r.service.Get(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.NewErrorResponse(err))
	}
	return c.JSON(http.StatusFound, helper.NewSuccessResponse(*response))
}

// query verify rest params, call service method to execute business logic
// and return JSON data
func (r *userResource) query(c echo.Context) error {
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
func (r *userResource) create(c echo.Context) error {
	var model model.User
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
func (r *userResource) update(c echo.Context) error {
	id := c.Param("userID")
	model, err := r.service.Get(id)
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
func (r *userResource) delete(c echo.Context) error {
	response, err := r.service.Delete(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.NewErrorResponse(err))
	}

	return c.JSON(http.StatusOK, helper.NewSuccessResponse(*response))
}
