package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginPage := e.Group("/")
	user := e.Group("/users")
	report := e.Group("/reports")
	admin := e.Group("/admins")
	faq:=e.Group("/faq")
	recybot := e.Group("/recybot")

	RouteLoginPage(loginPage, db)
	RouteUser(user, db)
	RouteReport(report, db)
	RouteAdmin(admin, db)
	RouteDropPoint(admin, db)
	RouteFaqs(faq, db)
	RouteRecybot(recybot,db)
	RouteAchievement(admin, db)
}
