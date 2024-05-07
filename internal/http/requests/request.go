package requests

import (
	BaseRequest "ErgoGo/pkg/http/request"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type Request struct {
	BaseRequest.Request
}

var request Request

func Validate(c *gin.Context, obj interface{}, handler BaseRequest.ValidatorFunc) bool {
	return request.Validate(c, obj, handler)
}

func validate(data interface{}, rules, messages govalidator.MapData) map[string][]string {
	return request.ValidateWithMessage(data, rules, messages)
}

func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	return request.ValidateFile(c, data, rules, messages)
}
