package transaction

import (
	"context"

	"gorm.io/gorm"
)

type contextKey string

const tokenContextKey contextKey = "key"

// WithContext ...
func WithContext(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(tokenContextKey).(*gorm.DB)
	return tx, ok
}

// NewContext ...
func NewContext(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, tokenContextKey, tx)
}
