package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	fmt.Println("非嵌套map")
	validateMap()
	fmt.Println("嵌套map")
	validateNestedMap()
}

func validateMap() {
	user := map[string]interface{}{"name": "Arshiya Kiani", "email": "zytel3301@gmail.com"}

	// Every rule will be applied to the item of the data that the offset of rule is pointing to.
	// So if you have a field "email": "omitempty,required,email", the validator will apply these
	// rules to offset of email in user data
	rules := map[string]interface{}{"name": "required,min=8,max=32", "email": "omitempty,required,email"}

	// ValidateMap will return map[string]error.
	// The offset of every item in errs is the name of invalid field and the value
	// is the message of error. If there was no error, ValidateMap method will
	// return an EMPTY map of errors, not nil. If you want to check that
	// if there was an error or not, you must check the length of the return value

	// ValidateMap将返回map[string]error。errs中每个项的偏移量是无效字段的名称，值是错误消息。
	// 如果没有错误，ValidateMap方法将返回一个错误的EMPTY映射，而不是nil。
	// 如果要检查是否存在错误，则必须检查返回值的长度
	errs := validate.ValidateMap(user, rules)

	if len(errs) > 0 {
		fmt.Println(errs)
		// The user is invalid
	}

	// The user is valid
}

func validateNestedMap() {

	data := map[string]interface{}{
		"name":  "Arshiya Kiani",
		"email": "zytel3301@gmail.com",
		"details": map[string]interface{}{
			"family_members": map[string]interface{}{
				"father_name": "Micheal",
				"mother_name": "Hannah",
			},
			"salary": "1000",
			"phones": []map[string]interface{}{
				{
					"number": "11-111-1111",
					"remark": "home",
				},
				{
					"number": "22-222-2222",
					"remark": "work",
				},
			},
		},
	}

	// Rules must be set as the structure as the data itself. If you want to dive into the
	// map, just declare its rules as a map

	// 规则必须设置为数据本身的结构。
	// 如果你想深入map，只需将其规则声明为map即可
	rules := map[string]interface{}{
		"name":  "min=4,max=32",
		"email": "required,email",
		"details": map[string]interface{}{
			"family_members": map[string]interface{}{
				"father_name": "required,min=4,max=32",
				"mother_name": "required,min=4,max=32",
			},
			"salary": "number",
			"phones": map[string]interface{}{
				"number": "required,min=4,max=32",
				"remark": "required,min=1,max=32",
			},
		},
	}

	if len(validate.ValidateMap(data, rules)) == 0 {
		// Data is valid
	}

	// Data is invalid
}