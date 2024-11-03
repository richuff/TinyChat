package constant

import "github.com/gin-gonic/gin"

type Any interface{}

type Result struct {
	code    int    //编码：1成功，0和其它数字为失败
	message string //错误信息，或者提示信息
	data    Any    //数据
}

func NewResult(code int, message string) *Result {
	return &Result{
		code:    code,
		message: message,
	}
}

func (result *Result) Success() Any {
	return gin.H{
		"code":    result.code,
		"message": result.message,
	}
}

func (result *Result) SuccessByData(data Any) Any {
	return gin.H{
		"code":    result.code,
		"message": result.message,
		"data":    data,
	}
}

func (result *Result) Error() Any {
	return gin.H{
		"code":    result.code,
		"message": result.message,
	}
}

func (result *Result) ErrorByData(data Any) Any {
	return gin.H{
		"code":    result.code,
		"message": result.message,
		"data":    data,
	}
}
