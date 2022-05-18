package handle

import (
	"errors"
	"reflect"
)

// 遍历对象字段，赋值
func CopyData(dst, src interface{}) error {
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)

	// dst必须结构体指针类型
	if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}
	dstType, dstValue = dstType.Elem(), dstValue.Elem()

	// src必须为结构体或者结构体指针
	if srcType.Kind() != reflect.Struct && srcType.Kind() != reflect.Ptr {
		return errors.New("src type should be a struct or a struct pointer")
	}
	if srcType.Kind() == reflect.Ptr {
		srcType, srcValue = srcType.Elem(), srcValue.Elem()
	}

	for m := 0; m < srcType.NumField(); m++ {
		// 属性
		property := dstType.Field(m)
		// 待填充属性值
		propertyValue := srcValue.FieldByName(property.Name)

		// 无效，说明src没有这个属性 || 属性同名但类型不同
		if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
			continue
		}

		if dstValue.Field(m).CanSet() {
			dstValue.Field(m).Set(propertyValue)
		}
	}
	return nil
}
