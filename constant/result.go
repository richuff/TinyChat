package constant

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Any interface{}

type Result struct {
	code    int    //编码：1成功，0和其它数字为失败
	message string //错误信息，或者提示信息
}

func NewResult() *Result {
	return &Result{}
}

func (*Result) SetCode(code int) *Result {
	return &Result{
		code: code,
	}
}

func (*Result) SetMessage(message string) *Result {
	return &Result{
		message: message,
	}
}

func (result *Result) Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    result.code,
		"message": result.message,
	})
}

func (result *Result) SuccessByData(c *gin.Context, data Any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    result.code,
		"message": result.message,
		"data":    data,
	})
}

func (result *Result) Error(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    result.code,
		"message": result.message,
	})
}

//func (result *Result) ErrorByData(data Any) Any {
//	return gin.H{
//		"code":    result.code,
//		"message": result.message,
//		"data":    data,
//	}
//}
//func (result *Result) Success() Any {
//	return gin.H{
//		"code":    result.code,
//		"message": result.message,
//	}
//}
//func NewResult(code int, message string) *Result {
//	return &Result{
//		code:    code,
//		message: message,
//	}
//}
