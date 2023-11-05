package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	admin := e.Group("/admins")
	user := e.Group("/users")
	RouteAdmin(admin, db)
	RouteUser(user, db)
}
