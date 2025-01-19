package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const ginContextKey = "Gin"

type ContextMiddleware struct{}

func (m *ContextMiddleware) Bind() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), ginContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func ContextToGinContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(ginContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not convert to gin.Context")
		return nil, err
	}

	converted, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return converted, nil
}

func NewContextMiddleware() (*ContextMiddleware, error) {
	return &ContextMiddleware{}, nil
}
