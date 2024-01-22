package curd

import (
	"github.com/google/uuid"
	"github.com/segmentio/ksuid"
	"github.com/zgwit/iot-master/v4/lib"
	"reflect"
)

func GenerateRandomKey(l int) func(value interface{}) error {
	return func(data interface{}) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			field.SetString(lib.RandomString(l))
		}
		return nil
	}
}

func GenerateRandomId[T any](l int) func(data *T) error {
	return func(data *T) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			field.SetString(lib.RandomString(l))
		}
		return nil
	}
}

func GenerateUuidKey(data interface{}) error {
	value := reflect.ValueOf(data).Elem()
	field := value.FieldByName("id")
	//使用UUId作为Id
	//field.IsZero() 如果为空串时，生成UUID
	if field.Len() == 0 {
		field.SetString(uuid.NewString())
	}
	return nil
}

func GenerateKSUID[T any]() func(data *T) error {
	return func(data *T) error {
		value := reflect.ValueOf(data).Elem()
		field := value.FieldByName("Id")
		//使用UUId作为Id
		//field.IsZero() 如果为空串时，生成UUID
		if field.Len() == 0 {
			key := ksuid.New().String()
			field.SetString(key)
		}
		return nil
	}
}
