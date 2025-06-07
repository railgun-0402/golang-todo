package services

import "errors"

// 内部に何もエラーをラップしていない「起点」となるエラー
var ErrNoData = errors.New("get 0 record from db.Query")