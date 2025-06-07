package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// エラーが発生した時のレスポンス処理を一括実施
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// 受け取ったエラーを独自エラーへ変換
	var appErr *TodoAppError
	if !errors.As(err, &appErr) {
		appErr = &TodoAppError{
			ErrCode: UnKnown,
			Message: "internal process failed",
			Err: err,
		}
	}

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}