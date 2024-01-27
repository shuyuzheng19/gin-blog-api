package config

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var validate *validator.Validate

var trans ut.Translator

var defaultLang = "zh" //验证翻译语言 英语是 en 中文是 zh 也可以是其他国家语言

// 初始化参数验证器
func init() {
	validate = validator.New()

	uniTrans := ut.New(zh.New(), en.New())

	trans, _ = uniTrans.GetTranslator(defaultLang)

	var err error
	switch defaultLang {
	case "zh":
		err = zh_translations.RegisterDefaultTranslations(validate, trans)
		break
	case "en":
		err = en_translations.RegisterDefaultTranslations(validate, trans)
		break
	default:
		err = en_translations.RegisterDefaultTranslations(validate, trans)
		break
	}
	helper.ErrorPanicAndMessage(err, "注册翻译器失败")
}

func ValidateStruct(object interface{}) map[string][]string {
	var err = validate.Struct(object)
	if err != nil {
		return Translate(err)
	}
	return nil
}

// Translate 翻译所有字段验证失败的信息
func Translate(err error) map[string][]string {
	var result = make(map[string][]string)

	errors := err.(validator.ValidationErrors)

	for _, err := range errors {
		result[err.Field()] = append(result[err.Field()], err.Translate(trans))
	}
	return result
}

// TranslateFirst 只翻译第一个字段的错误信息
func TranslateFirst(object interface{}) *string {

	var err = validate.Struct(object)

	if err == nil {
		return nil
	}

	errors := err.(validator.ValidationErrors)

	if len(errors) == 0 {
		return nil
	}

	var message = errors[0].Translate(trans)

	return &message
}

// ValidateError 验证结构体是否检查成功，如果失败则直接panic
func ValidateError(object interface{}) {
	var message = TranslateFirst(object)

	if message != nil {
		var result interface{} = common.AutoFail(common.FieldValidationFail)
		panic(result)
	}
}
