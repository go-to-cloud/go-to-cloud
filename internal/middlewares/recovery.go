package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strings"
)

// GenericRecovery 通用错误 (panic) 拦截中间件、对可能发生的错误进行拦截、统一记录
func GenericRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即
		errStr := ""
		if conf.Environment.Debugger {
			errStr = fmt.Sprintf("%v", err)
		}
		response.GetResponse().SetHttpCode(http.StatusInternalServerError).FailCode(c, http.StatusInternalServerError, errStr)
	})
}

// PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	s1 := "An error occurred in the server's internal code："
	var build strings.Builder
	build.WriteString(s1)
	build.Write(b)
	errStr := build.String()
	return len(errStr), errors.New(errStr)
}
