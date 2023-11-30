package route

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func New(e *echo.Echo, db *gorm.DB) {
	loginPage := e.Group("/")
	user := e.Group("/users")
	admin := e.Group("/admins")
	report := e.Group("/reports")
	faq := e.Group("/faq")
	recybot := e.Group("/recybot")
	mission := e.Group("")
	article := e.Group("")

	RouteLoginPage(loginPage, db)
	RouteUser(user, db)
	RouteReport(report, db)
	RouteAdmin(admin, db)
	RouteArticle(article, db)
	RouteDropPoint(admin, db)
	RouteFaqs(faq, db)
	RouteRecybot(recybot, db)
	RouteAchievement(admin, db)
	RouteVoucher(admin, db)
	RouteMissions(mission,db)
}
