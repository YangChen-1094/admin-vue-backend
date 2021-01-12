package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"reflect"
)

func GetPage(c *gin.Context, pageSize int) int{
	ret :=0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		ret = (page - 1) * pageSize
	}
	return ret
}


//查找字符是否在数组中
func InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}