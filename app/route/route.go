package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginPage := e.Group("/")
	user := e.Group("/users")
	report := e.Group("/reports")
	admin:=e.Group("/admins")

	RouteLoginPage(loginPage,db)
	RouteUser(user, db)
	RouteReport(report, db)
	RouteAdmin(admin, db)
	RouteArticle(admin,db)
}
