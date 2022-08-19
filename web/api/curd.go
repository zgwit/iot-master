package api

import (
	"github.com/google/uuid"
	"reflect"
)

type hook func(value interface{}) error

func generateUUID(data interface{}) error {
	value := reflect.ValueOf(data).Elem()
	field := value.FieldByName("Id")
	//使用UUId作为Id
	//field.IsZero() 如果为空串时，生成UUID
	if field.Len() == 0 {
		field.SetString(uuid.NewString())
	}
	return nil
}
