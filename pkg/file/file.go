package file

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"my_gin/pkg/setting"
	"os"
	"path"
)

//获取文件大小
func GetFileSize(file multipart.File)(int, error){
	content, err := ioutil.ReadAll(file)
	return len(content), err
}

//获取文件的扩展名
func GetExt(file string)string{
	return path.Ext(file)
}

func CheckExist(file string)bool{
	_, err := os.Stat(file)
	return os.IsExist(err)
}

//检查权限是否足够
func CheckPermission(file string)bool{
	_, err := os.Stat(file)
	return os.IsPermission(err)
}

//不存在则创建
func IsNotExistMkDir(file string) error{
	if exists := CheckExist(file); exists == false {
		if err := MkDir(file); err != nil{
			return err
		}
	}
	return nil
}

//创建指定目录
func MkDir(file string) error{
	err := os.MkdirAll(file, os.ModePerm)
	if err != nil{
		return err
	}
	return nil
}

func Open(fileName string, flag int, perm os.FileMode)(*os.File, error){
	f, err := os.OpenFile(fileName, flag, perm)
	if err != nil{
		return nil, err
	}
	return f, nil
}

func GetExportPath() string{
	return fmt.Sprintf("%s", setting.AppSetting.ExportPath)
}

//导出到csv
func ExportToCsv(fileName string, data [][]string) error {
	exportPath := GetExportPath()
	if exportPath == ""{
		return fmt.Errorf("导出目录名错误")
	}
	err := IsNotExistMkDir(exportPath)
	if err != nil {
		return nil
	}
	f, err := os.Create(exportPath + fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.WriteAll(data)
	return nil
}

