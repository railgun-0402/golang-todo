package middlewares

import (
	"log"
	"net/http"
)

// 委譲によって、Header メソッド・Write メソッド・WriteHeader メソッドを持つ
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// コンストラクタ
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeaderメソッドのオーバーライド
func (rsw *resLoggingWriter) WriteHeader(code int) {
	// resLoggingWriter 構造体の code フィールドに、使うレスポンスコードを保存する
	rsw.code = code

	// HTTP レスポンスに使うレスポンスコードを指定
	rsw.ResponseWriter.WriteHeader(code)
}

// リクエスト・レスポンス情報をロギング
func LoggingMiddleware(next http.Handler) http.Handler {
	/**
	 * ハンドラ関数 func(w http.ResponseWriter, r *http.Request) を
	 * http.HandlerFunc 型にキャストすることで
	 * 戻り値である http.Handler インターフェースを満たすようにしている
	**/
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		// 前処理：リクエスト情報をログ記録
		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method)

		// リクエストに含まれるコンテキストに、トレースIDを付加
		ctx := SetTraceID(req.Context(), traceID)
		req = req.WithContext(ctx)

		// 返り値なしの ServeHTTP の中でどうレスポンスが作られたのかはわからない
		// →そこで自作ResponseWriter
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		// 後処理：自作 ResponseWriter からロギングしたいデータを出す
		log.Printf("[%d]res: %d", traceID, rlw.code)
	})
}
