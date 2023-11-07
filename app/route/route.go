package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	admin := e.Group("/admins")
	user := e.Group("/users")
	report := e.Group("/reports")
	RouteAdmin(admin, db)
	RouteUser(user, db)
	RouteReport(report, db)
}
