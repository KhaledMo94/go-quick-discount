package handler

import (
	"net/http"
	"time"

	"github.com/KhaledMo94/quick-discount/internal/handler/middleware"
	"github.com/KhaledMo94/quick-discount/internal/i18n"
	"github.com/KhaledMo94/quick-discount/internal/service"
	"github.com/gin-gonic/gin"
)

type MainCategoryHandler struct {
	s      *service.MainCategoryService
	locale string
	appURL string
	tr 		*i18n.Translator
}

type mainCategoryResource struct {
	ID          int64      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Image       *string    `json:"image"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

func NewMainCategoryHandler(s *service.MainCategoryService, locale, appURL string,tr *i18n.Translator) *MainCategoryHandler {
	return &MainCategoryHandler{s: s, locale: locale, appURL: appURL,tr: tr}
}

func (h *MainCategoryHandler) Register(r *gin.RouterGroup) {
	g := r.Group("/main-categories")
	{
		g.GET("", h.Index)
		g.GET("/:id", h.Show)
		g.POST("", h.Store)
		g.PUT("/:id", h.Update)
		g.DELETE("/:id", h.Destroy)
	}
}

func (h *MainCategoryHandler) Index(ctx *gin.Context) {
	locale := middleware.GetLocale(ctx)

	items, err := h.s.FindAll()
	if err != nil {
		CheckErrorAndReturnData(ctx, nil, &err, h.tr.T("ok", locale), http.StatusOK)
		return
	}

	out := make([]mainCategoryResource, 0, len(items))

	for _, item := range items {
		name, _ := NormalizeJSONByLocale(item.Name, locale)

		var description *string
		if item.Description != nil {
			if desc, err := NormalizeJSONByLocale(*item.Description, locale); err == nil {
				description = &desc
			}
		}

		imageURL := item.Image.URL(h.appURL)
		var image *string
		if imageURL != "" {
			image = &imageURL
		}

		out = append(out, mainCategoryResource{
			ID:          item.ID,
			Name:        name,
			Description: description,
			Image:       image,
			Status:      string(item.Status),
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}

	CheckErrorAndReturnData(ctx, out, &err, h.tr.T("ok", locale), http.StatusOK)
}

func (h *MainCategoryHandler) Show(ctx *gin.Context) {
	CheckErrorAndReturnData(ctx, nil, nil, "not implemented", http.StatusNotImplemented)
}

func (h *MainCategoryHandler) Store(ctx *gin.Context) {
	CheckErrorAndReturnData(ctx, nil, nil, "not implemented", http.StatusNotImplemented)
}

func (h *MainCategoryHandler) Update(ctx *gin.Context) {
	CheckErrorAndReturnData(ctx, nil, nil, "not implemented", http.StatusNotImplemented)
}

func (h *MainCategoryHandler) Destroy(ctx *gin.Context) {
	CheckErrorAndReturnData(ctx, nil, nil, "not implemented", http.StatusNotImplemented)
}
