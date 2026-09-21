package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"school-portal/internal/core"
	"school-portal/internal/handlers/web/parents"
	"school-portal/internal/handlers/web/teachers"
	webmw "school-portal/internal/middlewares/web"
	"school-portal/internal/repos"
	webroutes "school-portal/internal/routes/web"
	webservices "school-portal/internal/services/web"
)

// Rate limits for the brute-forceable login endpoints. These are conservative per-IP fixed
// windows. The parent flow has no OTP step, so parentLoginLimiter is its only brute-force
// defense; the teacher flow additionally requires two matching values per attempt.
const (
	parentLoginRateLimit   = 5
	parentLoginRateWindow  = time.Minute
	teacherLoginRateLimit  = 10
	teacherLoginRateWindow = time.Minute
)

type App struct {
	cfg    *core.Config
	db     *pgxpool.Pool
	Router http.Handler
}

func New(ctx context.Context, cfg *core.Config) (*App, error) {
	db, err := core.NewDBPool(ctx, cfg.DBURL)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	parentRepo := repos.NewParentRepo(db)
	teacherRepo := repos.NewTeacherRepo(db)

	parentService := webservices.NewParentService(parentRepo)
	teacherService := webservices.NewTeacherService(teacherRepo)

	parentHandler := parents.NewHandler(parentService)
	teacherHandler := teachers.NewHandler(teacherService)

	parentLoginLimiter := webmw.NewIPRateLimiter(parentLoginRateLimit, parentLoginRateWindow)
	teacherLoginLimiter := webmw.NewIPRateLimiter(teacherLoginRateLimit, teacherLoginRateWindow)

	router := webroutes.NewRouter(parentHandler, teacherHandler, parentLoginLimiter, teacherLoginLimiter)

	return &App{
		cfg:    cfg,
		db:     db,
		Router: router,
	}, nil
}

func (a *App) Close() {
	a.db.Close()
}
