package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const LocaleKey = "locale"

func Locale(defaultLocale string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.Query("lang")

		if lang == "" {
			if al := ctx.GetHeader("Accept-Language"); al != "" {
				lang = al
				if i := strings.IndexAny(lang, ",;"); i >= 0 {
					lang = lang[:i]
				}
				if i := strings.Index(lang, "-"); i >= 0 {
					lang = lang[:i]
				}
				lang = strings.TrimSpace(lang)
			}
		}
		if lang == "" {
			lang = defaultLocale
		}
		ctx.Set(LocaleKey, lang)
		ctx.Next()
	}
}

func GetLocale(ctx *gin.Context) string {
	if v, ok := ctx.Get(LocaleKey); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return "en"
}
