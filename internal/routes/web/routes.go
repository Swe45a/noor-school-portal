package web

import (
	"net/http"

	"school-portal/internal/handlers/web/parents"
	"school-portal/internal/handlers/web/teachers"
	webmw "school-portal/internal/middlewares/web"
)

// NewRouter wires the parent and teacher login endpoints. Each is guarded by its own per-IP
// rate limiter since both are brute-forceable by ID enumeration; the parent flow has no OTP
// step, so its limiter is the only brute-force defense in front of it.
func NewRouter(
	parentHandler *parents.Handler,
	teacherHandler *teachers.Handler,
	parentLoginLimiter *webmw.IPRateLimiter,
	teacherLoginLimiter *webmw.IPRateLimiter,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("POST /api/parents/login", parentLoginLimiter.Middleware(http.HandlerFunc(parentHandler.Login)))
	mux.Handle("POST /api/teachers/login", teacherLoginLimiter.Middleware(http.HandlerFunc(teacherHandler.Login)))

	return mux
}
