package apperrors

import (
	"errors"
	"log"
	"net/http"
	"todo/api/middlewares"

	"github.com/labstack/echo/v4"
)

// エラーが発生した時のレスポンス処理を一括実施
func ErrorHandler(ctx echo.Context, err error) error {
	// 受け取ったエラーを独自エラーへ変換
	var appErr *TodoAppError
	if !errors.As(err, &appErr) {
		appErr = &TodoAppError{
			ErrCode: UnKnown,
			Message: "internal process failed",
			Err: err,
		}
	}

	// Todo: middlewareが動かないとハンドラが動かなくなってしまう
	// エラーと一緒にトレースIDをログ出力
	traceID := middlewares.GetTraceID(ctx.Request().Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	return ctx.JSON(statusCode, appErr)
}