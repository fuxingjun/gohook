package util

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// toQueryStrWithoutEncode converts a map of parameters into a query string without URL encoding.
// It supports slices and arrays by repeating the key for each value.
//
// Parameters:
//
//	params: A map of key-value pairs to be converted.
//
// Returns:
//
//	A query string with key=value pairs joined by &.
func ToQueryStrWithoutEncode(params map[string]any) string {
	if params == nil {
		return ""
	}
	var parts []string

	for key, value := range params {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < v.Len(); i++ {
				parts = append(parts, fmt.Sprintf("%s=%v", key, v.Index(i).Interface()))
			}
		default:
			parts = append(parts, fmt.Sprintf("%s=%v", key, value))
		}
	}

	return strings.Join(parts, "&")
}

// ensureDir 确保指定的目录存在，如果不存在则创建它
func EnsureDir(path string) error {
	if path == "" {
		return os.ErrInvalid
	}
	// 统一转换为正斜杠格式，便于判断
	normalized := filepath.ToSlash(path)
	endsWithSlash := strings.HasSuffix(normalized, "/")
	var target string
	if endsWithSlash {
		// 视为目录路径，创建该目录
		target = path
	} else {
		// 视为文件路径，创建父目录
		target = filepath.Dir(path)
	}
	// 清理路径，避免冗余符号（如 ./, ../）
	target = filepath.Clean(target)
	// 递归创建目录
	return os.MkdirAll(target, 0755)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
