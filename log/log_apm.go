package log

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

func NewElasticAPM(c *gin.Context, name string, spanType string) (*apm.Span, context.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), name, spanType)

	return span, ctx
}
