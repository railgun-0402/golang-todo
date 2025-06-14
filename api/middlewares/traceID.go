package middlewares

import (
	"context"
	"sync"
)

type traceIDKeyType struct{}

var traceIDKey = traceIDKeyType{}

var (
	logNo int = 1
	mu sync.Mutex
)

// 現在のID数を記録する (トレースIDが競合しないようロックを使用)
func newTraceID() int {
	var no int

	// インクリメントをロックし、同じ番号のID作成防止
	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

// ContextにトレースIDを付加して返却
func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// トレースIDをContextから取り出す
func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey)

	// Valueはany型が返却されるので、int型にassertion
	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}
