/*
 * @Author: tj
 * @Date: 2022-10-21 16:15:51
 * @LastEditors: tj
 * @LastEditTime: 2022-11-02 12:18:08
 * @FilePath: \createApiMarkdown\document\document.go
 */
package document

import (
	"os"
	"reflect"
	"strings"

	"createmd/common"

	"github.com/sirupsen/logrus"
)

const (
	requiredKey string = "required"
)

var (
	log = logrus.WithFields(logrus.Fields{
		"document": "",
	})
)

func NewDocument(url string, options ...Option) *Document {
	d := &Document{
		Title:   "接口文档",
		Version: "1.0",
		Url:     url,
		Items:   make([]*DocItem, 0, 8),
		mdKey:   "validate",
	}

	for _, option := range options {
		option(d)
	}

	return d
}

// 获取接口列表
func (d *Document) GetItems() []*DocItem {
	return d.Items
}

// 添加接口对象
func (d *Document) AddDocItem(item *DocItem) error {
	if item == nil {
		return os.ErrInvalid
	}

	reqFields, err := d.parseReqOrRsp(item.Request)
	if err != nil {
		log.Errorln("AddDocItem Request parseReqOrRsp error:", err)
		return err
	}
	item.ReqFields = reqFields

	rspFields, err := d.parseReqOrRsp(item.Response)
	if err != nil {
		log.Errorln("AddDocItem Response parseReqOrRsp error:", err)
		return err
	}
	item.RspFields = rspFields

	d.Items = append(d.Items, item)

	return nil
}

// 解析接口的请求对象或者返回对象
func (d *Document) parseReqOrRsp(param interface{}) ([]*Field, error) {
	if param == nil {
		return nil, nil
	}

	value, err := d.checkKind(param)
	if err != nil {
		return nil, err
	}

	return d.getFields(value)
}

func (d *Document) checkKind(param interface{}) (*reflect.Value, error) {
	value := reflect.ValueOf(param)
	if !value.IsValid() {
		return nil, nil
	}

	switch value.Kind() {
	case reflect.Slice:
		if value.Len() <= 0 {
			return nil, nil
		}
		value = value.Index(0)

		// []*data
		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				return nil, nil
			}
			value = value.Elem()
		}

	case reflect.Ptr:
		if value.IsNil() {
			return nil, nil
		}
		value = value.Elem()

	case reflect.Struct:
		break

	default:
		return nil, os.ErrInvalid
	}

	return &value, nil
}

func (d *Document) getFields(value *reflect.Value) ([]*Field, error) {
	valueType := value.Type()
	fields := make([]*Field, 0)
	// 遍历结构体所有成员
	for i := 0; i < value.NumField(); i++ {
		// 获取结构体字段值
		fieldValue := value.Field(i)

		structField := valueType.Field(i)
		// 内嵌结构体
		if structField.Anonymous {
			subFields, err := d.parseReqOrRsp(fieldValue.Interface())
			if err != nil {
				return nil, err
			}

			fields = append(fields, subFields...)
			continue
		}

		// 非内嵌结构体
		field := &Field{
			Name:        common.FirstLower(structField.Name),
			Kind:        structField.Type.Kind().String(),
			IsRequired:  d.isRequired(structField),
			Description: d.getDescription(structField),
			List:        nil,
		}

		err := d.getStructField(fieldValue, field)
		if err != nil {
			return nil, err
		}

		fields = append(fields, field)
	}

	return fields, nil
}

// 解析结构体字段的下级结构
func (d *Document) getStructField(fieldValue reflect.Value, field *Field) error {
	switch field.Kind {
	case reflect.Interface.String(), reflect.Struct.String():
		subFields, err := d.parseReqOrRsp(fieldValue.Interface())
		if err != nil {
			return err
		}

		field.List = subFields

	case reflect.Ptr.String():
		if fieldValue.IsNil() {
			break
		}

		subValue := fieldValue.Elem()

		subFields, err := d.parseReqOrRsp(subValue.Interface())
		if err != nil {
			return err
		}

		field.List = subFields

	case reflect.Slice.String():
		if fieldValue.Len() <= 0 {
			break
		}

		field.Kind += "["
		switch fieldValue.Index(0).Interface().(type) {
		case bool:
			field.Kind += reflect.Bool.String()
		case string:
			field.Kind += reflect.String.String()
		case int:
			field.Kind += reflect.Int.String()
		case uint:
			field.Kind += reflect.Uint.String()
		case uint8:
			field.Kind += reflect.Uint8.String()
		case uint16:
			field.Kind += reflect.Uint16.String()
		case uint32:
			field.Kind += reflect.Uint32.String()
		case uint64:
			field.Kind += reflect.Uint64.String()
		case int8:
			field.Kind += reflect.Int8.String()
		case int16:
			field.Kind += reflect.Int16.String()
		case int32:
			field.Kind += reflect.Int32.String()
		case int64:
			field.Kind += reflect.Int64.String()
		case float64:
			field.Kind += reflect.Float64.String()
		case float32:
			field.Kind += reflect.Float32.String()
		default: //结构体
			field.Kind += reflect.Struct.String()
			subFields, err := d.parseReqOrRsp(fieldValue.Interface())
			if err != nil {
				return err
			}

			field.List = subFields
		}
		field.Kind += "]"
	}

	return nil
}

func (d *Document) isRequired(field reflect.StructField) bool {
	tag := field.Tag.Get(d.mdKey)

	return strings.Contains(tag, requiredKey)
}

func (d *Document) getDescription(field reflect.StructField) string {
	tag := field.Tag.Get(d.mdKey)

	if strings.Contains(tag, requiredKey) {
		if strings.Contains(tag, ",") {
			return tag[len(requiredKey)+1:]
		} else {
			return tag[len(requiredKey):]
		}
	}

	return tag
}
