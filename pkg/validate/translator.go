package validate

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zhoudm1743/go-frame/pkg/facades"
)

var (
	// Trans 全局翻译器
	Trans ut.Translator
)

// InitValidator 初始化验证器和中文翻译
func InitValidator() error {
	// 获取gin默认的validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		Validator = v

		// 注册自定义验证规则
		RegisterCustomValidations(v)

		// 初始化中文翻译器
		zhT := zh.New()
		uni := ut.New(zhT, zhT)
		Trans, _ = uni.GetTranslator("zh")

		// 注册默认翻译
		err := zhtranslations.RegisterDefaultTranslations(v, Trans)
		if err != nil {
			return err
		}

		// 注册自定义验证规则的中文翻译
		registerCustomTranslations()

		// 设置门面实例
		facades.SetValidator(v)
		facades.SetTranslator(Trans)

		// 记录初始化成功日志
		facades.Log.Infof("验证器初始化成功")
	}

	return nil
}

// 注册自定义验证规则的中文翻译
func registerCustomTranslations() {
	// 手机号验证翻译
	Validator.RegisterTranslation("phone", Trans, func(ut ut.Translator) error {
		return ut.Add("phone", "{0}必须是有效的中国手机号", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})

	// 身份证验证翻译
	Validator.RegisterTranslation("idcard", Trans, func(ut ut.Translator) error {
		return ut.Add("idcard", "{0}必须是有效的身份证号", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("idcard", fe.Field())
		return t
	})

	// 中文姓名验证翻译
	Validator.RegisterTranslation("chinese_name", Trans, func(ut ut.Translator) error {
		return ut.Add("chinese_name", "{0}必须是2-4个中文字符", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("chinese_name", fe.Field())
		return t
	})

	// 强密码验证翻译
	Validator.RegisterTranslation("strong_password", Trans, func(ut ut.Translator) error {
		return ut.Add("strong_password", "{0}必须至少8位，包含大小写字母、数字和特殊字符", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("strong_password", fe.Field())
		return t
	})

	// 中文字符验证翻译
	Validator.RegisterTranslation("chinese", Trans, func(ut ut.Translator) error {
		return ut.Add("chinese", "{0}必须是中文字符", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("chinese", fe.Field())
		return t
	})

	// 日期格式验证翻译
	Validator.RegisterTranslation("date", Trans, func(ut ut.Translator) error {
		return ut.Add("date", "{0}必须是有效的日期格式(YYYY-MM-DD)", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("date", fe.Field())
		return t
	})

	// URL格式验证翻译
	Validator.RegisterTranslation("url", Trans, func(ut ut.Translator) error {
		return ut.Add("url", "{0}必须是有效的URL地址", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("url", fe.Field())
		return t
	})

	// 邮政编码验证翻译
	Validator.RegisterTranslation("zipcode", Trans, func(ut ut.Translator) error {
		return ut.Add("zipcode", "{0}必须是有效的6位邮政编码", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("zipcode", fe.Field())
		return t
	})
}

// TranslateError 翻译验证错误为中文
func TranslateError(err error) map[string]string {
	if err == nil {
		return nil
	}

	result := make(map[string]string)
	errors := err.(validator.ValidationErrors)

	for _, fieldError := range errors {
		result[fieldError.Field()] = fieldError.Translate(Trans)
	}

	return result
}
