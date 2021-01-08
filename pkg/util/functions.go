package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/axgle/mahonia"
	"hash/crc32"
	"reflect"
	"strings"
)

/**
 * @brief  把当前字符串按照指定方式进行编码
 * @param[in]       src				   待进行转码的字符串
 * @param[in]       srcCode			   字符串当前编码
 * @param[in]       tagCode			   要转换的编码
 * @return   进行转换后的字符串
 */
func ConvertToString(src string, srcCode string, tagCode string) (string, error) {
	if len(src) == 0 || len(srcCode) == 0 || len(tagCode) == 0 {
		return "", errors.New("input arguments error")
	}
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)

	return result, nil
}

// GetBranchInsertSql 获取批量添加数据sql语句
func GetBranchInsertSql(objs []interface{}, tableName string) string {
	if len(objs) == 0 {
		return ""
	}
	fieldName := ""
	var valueTypeList []string
	fieldNum := reflect.TypeOf(objs[0]).NumField()
	fieldT := reflect.TypeOf(objs[0])

	for a := 0; a < fieldNum; a++ {
		gormName := fieldT.Field(a).Tag.Get("gorm")
		name := GetColumnName(gormName)
		// 添加字段名
		if a == fieldNum-1 {
			fieldName += fmt.Sprintf("`%s`", name)
		} else {
			fieldName += fmt.Sprintf("`%s`,", name)
		}
		// 获取字段类型
		if fieldT.Field(a).Type.Name() == "string" {
			valueTypeList = append(valueTypeList, "string")
		} else if strings.Index(fieldT.Field(a).Type.Name(), "uint") != -1 {
			valueTypeList = append(valueTypeList, "uint")
		} else if strings.Index(fieldT.Field(a).Type.Name(), "int") != -1 {
			valueTypeList = append(valueTypeList, "int")
		}
	}
	var valueList []string
	for _, obj := range objs {
		objV := reflect.ValueOf(obj)
		v := "("
		for index, i := range valueTypeList {
			if index == fieldNum-1 {
				v += GetFormatField(objV, index, i, "")
			} else {
				v += GetFormatField(objV, index, i, ",")
			}
		}
		v += ")"
		valueList = append(valueList, v)
	}
	insertSql := fmt.Sprintf("insert into `%s` (%s) values %s", tableName, fieldName, strings.Join(valueList, ",")+";")
	return insertSql
}

/**封装批量更新的sql
 * @param[in]       objs			   需要更新的数据
 * @param[in]       tableName		   表名
 * @param[in]       keys		       键值数组
 * @param[in]       notUpdateKeys	   不更新的字段
 * @param[in]       primaryKey		   主键key
 */
func GetBranchUpdateSql(objs []interface{}, tableName string, keys []interface{}, notUpdateKeys []string, primaryKey string) string {
	if len(objs) == 0 {
		return ""
	}
	if primaryKey == "" {
		primaryKey = "id"
	}
	fieldNum := reflect.TypeOf(objs[0]).NumField()
	fieldT := reflect.TypeOf(objs[0])
	sql := fmt.Sprintf(" UPDATE `%s` SET ", tableName)

	for a := 0; a < fieldNum; a++ {
		gormName := fieldT.Field(a).Tag.Get("gorm")
		name := GetColumnName(gormName)

		if InArray(name, notUpdateKeys) {
			continue
		}
		sql += fmt.Sprintf(" `%s` = CASE %s", name, primaryKey)
		for key, oneItem := range objs {
			oneItemValue := reflect.ValueOf(oneItem)
			thisValue:= oneItemValue.Field(a).Interface()

			primaryKeyVal := reflect.ValueOf(keys[key])
			keyType := primaryKeyVal.Kind()
			switch keyType.String() {
				case "int":
					val := primaryKeyVal.Int()
					sql += fmt.Sprintf(" WHEN %d THEN '%s' ", val, thisValue)
				case "string":
					val := primaryKeyVal.String()
					sql += fmt.Sprintf(" WHEN %s THEN '%s' ", val, thisValue)
				default:
					val := ""
					sql += fmt.Sprintf(" WHEN %s THEN '%s' ", val, thisValue)
			}
		}
		sql += " END,"
	}
	sql = strings.Trim(sql, ",")
	ids := Implode(keys, ",")
	sql += fmt.Sprintf(" WHERE %s IN (%s)", primaryKey, ids)
	return sql
}

func Implode(list interface{}, seq string) string {
	listValue := reflect.Indirect(reflect.ValueOf(list))
	if listValue.Kind() != reflect.Slice {
		return ""
	}
	count := listValue.Len()
	listStr := make([]string, 0, count)
	for i := 0; i < count; i++ {
		v := listValue.Index(i)
		if str, err := getValue(v); err == nil {
			listStr = append(listStr, str)
		}
	}
	return strings.Join(listStr, seq)
}


func getValue(value reflect.Value) (res string, err error) {
	switch value.Kind() {
	case reflect.Ptr:
		res, err = getValue(value.Elem())
	default:
		res = fmt.Sprint(value.Interface())
	}
	return
}

// GetFormatFeild 获取字段类型值转为字符串
func GetFormatField(objV reflect.Value, index int, t string, sep string) string {
	v := ""
	if t == "string" {
		v += fmt.Sprintf("'%s'%s", objV.Field(index).String(), sep)
	} else if t == "uint" {
		v += fmt.Sprintf("%d%s", objV.Field(index).Uint(), sep)
	} else if t == "int" {
		v += fmt.Sprintf("%d%s", objV.Field(index).Int(), sep)
	}
	return v

}
// GetColumnName 获取字段名
func GetColumnName(jsonName string) string {
	for _, name := range strings.Split(jsonName, ";") {
		if strings.Index(name, "column") == -1 {
			continue
		}
		return strings.Replace(name, "column:", "", 1)
	}
	return ""
}


func EncryptCRC32(str string) uint32{
	return crc32.ChecksumIEEE([]byte(str))
}

// 生成md5
func EncryptMD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func EncryptSHA1(str string) string{
	c:=sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}
