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

	voucher := e.Group("/manage/vouchers", jwt.JWTMiddleware())
	voucher.POST("", voucherHandler.CreateVoucher)
	voucher.GET("", voucherHandler.GetAllVoucher)
	voucher.GET("/:id", voucherHandler.GetVoucherById)
	voucher.PUT("/:id", voucherHandler.UpdateVoucher)
	voucher.DELETE("/:id", voucherHandler.DeleteVoucherById)
}
