package parents

import (
	web "school-portal/internal/services/web"
)

type Handler struct {
	service web.ParentService
}

func NewHandler(service web.ParentService) *Handler {
	return &Handler{service: service}
}
