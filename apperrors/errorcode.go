package apperrors

type ErrCode string

const (
	UnKnown ErrCode = "U0000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed ErrCode = "S002"
	NAData ErrCode = "S003"
	NoTargetData ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam ErrCode = "R002"
)