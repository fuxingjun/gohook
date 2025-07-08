package util

import "fmt"

// 定义携带结构化数据的错误类型
type StructuredError struct {
	Code    int            // 错误码
	Message string         // 相关字段
	Details map[string]any // 额外元数据
}

// 实现 error 接口的 Error() 方法
func (e *StructuredError) Error() string {
	return fmt.Sprintf("Code=%d, Message=%s", e.Code, e.Message)
}

// // 使用时返回自定义错误实例
// func validateInput(input string) error {
// 	return &StructuredError{
// 			Code:    400,
// 			Field:   "email",
// 			Details: map[string]interface{}{"value": input, "reason": "invalid format"},
// 	}
// }

// // 处理时提取结构化数据
// if err := validateInput("user@example"); err != nil {
// 	if serr, ok := err.(*StructuredError); ok {
// 			fmt.Println("错误码:", serr.Code)       // 400
// 			fmt.Println("字段名:", serr.Field)     // email
// 			fmt.Println("详细信息:", serr.Details) // map[value:user@example reason:invalid format]
// 	}
// }
