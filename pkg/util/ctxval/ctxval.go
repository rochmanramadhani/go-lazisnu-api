package ctxval

import (
	"context"
	"gorm.io/gorm"

	abstraction "github.com/rochmanramadhani/go-lazisnu-api/internal/model/abstraction"
)

type key string

var (
	keyAuth       key = "lazisnu-auth-context"
	keyUploadFile key = "lazisnu-upload-file-context"
	keyTrx        key = "lazisnu-trx-context"
)

func SetAuthValue(ctx context.Context, payload *abstraction.AuthContext) context.Context {
	return context.WithValue(ctx, keyAuth, payload)
}

func GetAuthValue(ctx context.Context) *abstraction.AuthContext {
	val, ok := ctx.Value(keyAuth).(*abstraction.AuthContext)
	if ok {
		return val
	}
	return nil
}

func SetTrxValue(ctx context.Context, trx *gorm.DB) context.Context {
	return context.WithValue(ctx, keyTrx, trx)
}

func GetTrxValue(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(keyTrx).(*gorm.DB)
	if ok {
		return val
	}
	return nil
}

func SetUploadFileValues(ctx context.Context, payloads *[]abstraction.UploadFileContext) context.Context {
	return context.WithValue(ctx, keyUploadFile, payloads)
}

func GetUploadFileValue(ctx context.Context) *[]abstraction.UploadFileContext {
	val, ok := ctx.Value(keyUploadFile).(*[]abstraction.UploadFileContext)
	if ok {
		return val
	}
	return nil
}
