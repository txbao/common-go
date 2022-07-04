package validators

import (
	"errors"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"regexp"
)

/**
验证基类
txbao
*/

//type Users struct {
//	Name   string `form:"name" json:"name" validate:"required"`
//	Age   uint8 `form:"age" json:"age" validate:"required,gt=18"`
//	Passwd   string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
//	Code   string `form:"code" json:"code" validate:"required,len=6"`
//  Email        string    `validate:"required,email"`
//  Mobile        string    `validate:"required,mobile"`
//}

//https://www.cnblogs.com/wangkun122/articles/11023964.html
//验证规则
//
//required ：必填
//email：验证字符串是email格式；例：“email”
//url：这将验证字符串值包含有效的网址;例：“url”
//max：字符串最大长度；例：“max=20”
//min:字符串最小长度；例：“min=6”
//excludesall:不能包含特殊字符；例：“excludesall=0x2C”//注意这里用十六进制表示。
//len：字符长度必须等于n，或者数组、切片、map的len值为n，即包含的项目数；例：“len=6”
//eq：数字等于n，或者或者数组、切片、map的len值为n，即包含的项目数；例：“eq=6”
//ne：数字不等于n，或者或者数组、切片、map的len值不等于为n，即包含的项目数不为n，其和eq相反；例：“ne=6”
//gt：数字大于n，或者或者数组、切片、map的len值大于n，即包含的项目数大于n；例：“gt=6”
//gte：数字大于或等于n，或者或者数组、切片、map的len值大于或等于n，即包含的项目数大于或等于n；例：“gte=6”
//lt：数字小于n，或者或者数组、切片、map的len值小于n，即包含的项目数小于n；例：“lt=6”
//lte：数字小于或等于n，或者或者数组、切片、map的len值小于或等于n，即包含的项目数小于或等于n；例：“lte=6”
//
//跨字段验证
//
//如想实现比较输入密码和确认密码是否一致等类似场景
//
//eqfield=Field: 必须等于 Field 的值；
//nefield=Field: 必须不等于 Field 的值；
//gtfield=Field: 必须大于 Field 的值；
//gtefield=Field: 必须大于等于 Field 的值；
//ltfield=Field: 必须小于 Field 的值；
//ltefield=Field: 必须小于等于 Field 的值；
//eqcsfield=Other.Field: 必须等于 struct Other 中 Field 的值；
//necsfield=Other.Field: 必须不等于 struct Other 中 Field 的值；
//gtcsfield=Other.Field: 必须大于 struct Other 中 Field 的值；
//gtecsfield=Other.Field: 必须大于等于 struct Other 中 Field 的值；
//ltcsfield=Other.Field: 必须小于 struct Other 中 Field 的值；
//ltecsfield=Other.Field: 必须小于等于 struct Other 中 Field 的值；

// 验证
func Valid(form interface{}) error {
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	validate.RegisterValidation("mobile", validateMobile)
	validate.RegisterValidation("date", validateDate)
	validate.RegisterValidation("integer", validateInteger)
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	//验证器注册翻译器
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(form)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			//return errors.New(err.Translate(trans))
			switch err.Tag() {
			case "mobile":
				return errors.New("手机号格式不对")
			case "integer":
				return errors.New("整数格式不对")
			case "date":
				return errors.New("日期格式不对")
			default:
				return errors.New(err.Translate(trans))
			}
		}
	}
	return nil
}

//验证手机号
func validateMobile(fl validator.FieldLevel) bool {
	if pass, _ := regexp.MatchString(
		`^(1[3456789][0-9]{9})$`, fl.Field().String(),
	); pass {
		return true
	}
	return false
}

//验证日期
func validateDate(fl validator.FieldLevel) bool {
	if pass, _ := regexp.MatchString(
		`^([\d]{4}-[\d]{2}-[\d]{2})$`, fl.Field().String(),
	); pass {
		return true
	}
	return false
}

//验证整数
func validateInteger(fl validator.FieldLevel) bool {
	if pass, _ := regexp.MatchString(
		`^([0-9]+)$`, fl.Field().String(),
	); pass {
		return true
	}
	return false
}
