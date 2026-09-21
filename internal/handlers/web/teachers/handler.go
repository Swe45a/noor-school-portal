package teachers

import (
	web "school-portal/internal/services/web"
)

type Handler struct {
	service web.TeacherService
}

func NewHandler(service web.TeacherService) *Handler {
	return &Handler{service: service}
}
