package enum

//go:generate stringer -type BizCode -linecomment -output=biz_code_string.go
type BizCode int

const (
	BizCodeOK      BizCode = iota // 正常
	BizCodeFail                   // 失败
	BizCodeTimeout                // 超时
)
