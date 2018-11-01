package router

import (
	"github.com/lucasfloriani/go-mongo/dao"
	"github.com/lucasfloriani/go-mongo/handler"
	"github.com/lucasfloriani/go-mongo/service"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Setup creates routes from application with middlwares and handlers.
func Setup(db *mongo.Database) *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")

	userDAO := dao.NewUserDAO(db)
	handler.ServeUserResource(v1, service.NewUserService(userDAO))

	return e
}
