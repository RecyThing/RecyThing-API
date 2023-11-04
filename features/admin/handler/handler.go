package handler

import (
	"recything/features/admin/entity"
)

type AdminHandler struct {
	AdminService entity.AdminServiceInterface
}

func NewAdminHandler(admin entity.AdminServiceInterface) *AdminHandler {
	return &AdminHandler{AdminService: admin}
}
