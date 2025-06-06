package apperrors

type TodoAppError struct {
	ErrCode // ErrCode型のフィールド
	Message string
	Err error // エラーチェーンのための内部エラー
}

func (myErr *TodoAppError) Error() string {
	return myErr.Err.Error()
}

// 独自エラーに置き換える関数
func (code ErrCode) Wrap(err error, message string) error {
	return &TodoAppError{ErrCode: code, Message: message, Err: err}
}