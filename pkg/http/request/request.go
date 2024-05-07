package request

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"

	"ErgoGo/pkg/http/response"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

type Request struct {
}

func (r Request) Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	// 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		return false
	}

	// 表单验证
	err := handler(obj, c)

	// 判断验证是否通过
	if len(err) > 0 {
		response.ValidationError(c, err)
		return false
	}

	return true
}

func (r Request) ValidateWithMessage(data interface{}, rules, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

func (r Request) ValidateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}

	// 调用 govalidator 的 Validate 方法验证文件
	return govalidator.New(opts).Validate()
}
