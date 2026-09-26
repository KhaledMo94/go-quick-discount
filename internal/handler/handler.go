package handler

import (
	"context"
	"net/http"

	"github.com/KhaledMo94/quick-discount/internal/handler/middleware"
	"github.com/KhaledMo94/quick-discount/internal/i18n"
	"github.com/gin-gonic/gin"
)

type Pinger interface {
	Ping(ctx context.Context) error
}

type Handler struct {
	db             Pinger
	redis          Pinger
	ms             Pinger
	MainCategories *MainCategoryHandler
	tr             *i18n.Translator
	defaultLocale  string
}

func New(
	dbConn, redisConn, msConn Pinger,
	mainCategories *MainCategoryHandler,
	tr *i18n.Translator,
	defaultLocale string,
) *Handler {
	return &Handler{
		db:             dbConn,
		redis:          redisConn,
		ms:             msConn,
		MainCategories: mainCategories,
		tr:             tr,
		defaultLocale:  defaultLocale,
	}
}

func (h *Handler) Routes() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.Locale(h.defaultLocale))
	router.GET("/health", h.health)

	api := router.Group("/api")
	{
		mc := api.Group("/main-categories")
		{
			mc.GET("", h.MainCategories.Index)
		}
	}

	return router
}

func (h *Handler) health(c *gin.Context) {
	ctx := c.Request.Context()
	dbError := h.db.Ping(ctx)
	redisError := h.redis.Ping(ctx)
	msError := h.ms.Ping(ctx)

	status := "ok"
	code := http.StatusOK

	if dbError != nil || redisError != nil || msError != nil {
		status = "degraded"
		code = http.StatusServiceUnavailable
	}

	c.JSON(code, gin.H{"status": status})
}
