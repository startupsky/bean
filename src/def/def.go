package def

type BusinessException struct {
	ErrNo  int
	ErrMsg string
}

const (
	paramErrNo = 1000
)

const (
	UnExpectedErrNo  = 10001
	UnExpectedErrMsg = "意外错误"
)

//logic exception
var ParamException = &BusinessException{paramErrNo, "参数错误"}
