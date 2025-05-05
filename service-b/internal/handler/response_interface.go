package handler

import "github.com/gin-gonic/gin"

type ContentType string

const (
	CONTENT_TYPE_JSON ContentType = "application/json"
	CONTENT_TYPE_XML  ContentType = "application/xml"
	CONTENT_TYPE_M3U8 ContentType = "application/vnd.apple.mpegurl"
	CONTENT_TYPE_TEXT ContentType = "text/plain"
)

type ResponseController interface {
	IsAbort() bool
	GetStatusCode() int
	GetErrors() []error
	GetResponse() interface{}
	SetStatusCode(statusCode int)
	SetErrors(errors []error)
	AddError(err error)
	SetResponse(response interface{})
	SetResult(status int, response interface{})
	Write(ctx *gin.Context)
	SetContentType(contentType ContentType)
}

type baseResponseController struct {
	contentType ContentType
	statusCode  int
	err         []error
	response    interface{}
}

func (b *baseResponseController) IsAbort() bool {
	return b.statusCode >= 400 || len(b.err) > 0
}

func (b *baseResponseController) GetStatusCode() int {
	return b.statusCode
}

func (b *baseResponseController) GetErrors() []error {
	return b.err
}

func (b *baseResponseController) GetResponse() interface{} {
	return b.response
}

func (b *baseResponseController) SetStatusCode(statusCode int) {
	b.statusCode = statusCode
}

func (b *baseResponseController) SetErrors(errors []error) {
	b.err = append(b.err, errors...)
}

func (b *baseResponseController) AddError(err error) {
	b.err = append(b.err, err)
}

func (b *baseResponseController) SetResponse(response interface{}) {
	b.response = response
}

func (b *baseResponseController) SetResult(status int, response interface{}) {
	b.statusCode = status
	b.response = response
}

func (b *baseResponseController) Write(ctx *gin.Context) {
	ctx.Header("Content-Type", string(b.contentType))
	switch b.contentType {
	case CONTENT_TYPE_M3U8:
		ctx.Writer.Write([]byte(b.response.(string)))
	case CONTENT_TYPE_TEXT:
		ctx.String(b.statusCode, b.response.(string))
	default:
		ctx.JSON(b.statusCode, b.response)
	}
}

func (b *baseResponseController) SetContentType(contentType ContentType) {
	b.contentType = contentType
}

func NewJsonResponseController() ResponseController {
	return &baseResponseController{
		contentType: CONTENT_TYPE_JSON,
	}
}
