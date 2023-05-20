package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetResponse() *Response {
	return &Response{
		httpCode: http.StatusOK,
		Result: &Result{
			Code:    0,
			Message: "",
			Data:    nil,
			Cost:    "",
		},
	}
}
func BadRequest(ctx *gin.Context, data ...any) {
	GetResponse().withDataAndHttpCode(http.StatusBadRequest, ctx, data)
}

// Success 业务成功响应
func Success(ctx *gin.Context, data ...any) {
	if data != nil {
		GetResponse().WithDataSuccess(ctx, data[0])
		return
	}
	GetResponse().Success(ctx)
}

// Fail 业务失败响应
func Fail(ctx *gin.Context, code int, message *string, data ...any) {
	var msg string
	if message == nil {
		msg = ""
	} else {
		msg = *message
	}
	if data != nil {
		GetResponse().WithData(data[0]).FailCode(ctx, code, msg)
		return
	}
	GetResponse().FailCode(ctx, code, msg)
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Cost    string      `json:"cost"`
}

type Response struct {
	httpCode int
	Result   *Result
}

// Fail 错误返回
func (r *Response) Fail(ctx *gin.Context) *Response {
	r.SetCode(http.StatusInternalServerError)
	r.json(ctx)
	return r
}

// FailCode 自定义错误码返回
func (r *Response) FailCode(ctx *gin.Context, code int, msg ...string) *Response {
	r.SetCode(code)
	if msg != nil {
		r.WithMessage(msg[0])
	}
	r.json(ctx)
	return r
}

// Success 正确返回
func (r *Response) Success(ctx *gin.Context) *Response {
	r.SetCode(http.StatusOK)
	r.json(ctx)
	return r
}

// WithDataSuccess 成功后需要返回值
func (r *Response) WithDataSuccess(ctx *gin.Context, data interface{}) *Response {
	r.SetCode(http.StatusOK)
	r.WithData(data)
	r.json(ctx)
	return r
}

func (r *Response) withDataAndHttpCode(code int, ctx *gin.Context, data interface{}) *Response {
	r.SetHttpCode(code)
	if data != nil {
		r.WithData(data)
	}
	r.json(ctx)
	return r
}

// SetCode 设置返回code码
func (r *Response) SetCode(code int) *Response {
	r.Result.Code = code
	return r
}

// SetHttpCode 设置http状态码
func (r *Response) SetHttpCode(code int) *Response {
	r.httpCode = code
	return r
}

type defaultRes struct {
	Result any `json:"result"`
}

// WithData 设置返回data数据
func (r *Response) WithData(data interface{}) *Response {
	switch data.(type) {
	case string, int, bool:
		r.Result.Data = &defaultRes{Result: data}
	default:
		r.Result.Data = data
	}
	return r
}

// WithMessage 设置返回自定义错误消息
func (r *Response) WithMessage(message string) *Response {
	r.Result.Message = message
	return r
}

// json 返回 gin 框架的 HandlerFunc
func (r *Response) json(ctx *gin.Context) {
	r.Result.Cost = time.Since(ctx.GetTime("requestStartTime")).String()
	ctx.AbortWithStatusJSON(r.httpCode, r.Result)
}
