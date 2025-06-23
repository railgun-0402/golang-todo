package middlewares

import (
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

