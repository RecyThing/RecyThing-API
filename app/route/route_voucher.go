package route

import (
	"recything/features/voucher/handler"
	"recything/features/voucher/repository"
	"recything/features/voucher/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteVoucher(e *echo.Group, db *gorm.DB) {
	voucherRepository := repository.NewVoucherRepository(db)
	voucherService := service.NewVoucherService(voucherRepository)
	voucherHandler := handler.NewVoucherHandler(voucherService)

	admin := e.Group("/admins/manage/vouchers", jwt.JWTMiddleware())
	admin.POST("", voucherHandler.CreateVoucher)
	admin.GET("", voucherHandler.GetAllVoucher)
	admin.GET("/:id", voucherHandler.GetVoucherById)
	admin.PUT("/:id", voucherHandler.UpdateVoucher)
	admin.DELETE("/:id", voucherHandler.DeleteVoucherById)

	user := e.Group("/vouchers", jwt.JWTMiddleware())
	user.GET("", voucherHandler.GetAllVoucher)
	user.GET("/:id", voucherHandler.GetVoucherById)
}
