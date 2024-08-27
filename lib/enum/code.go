package enum

//go:generate stringer -type Code -linecomment
type Code int

const (
	CodeOK      Code = iota // 正常
	CodeFail                // 失败
	CodeTimeout             // 超时
)
