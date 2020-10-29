package file

import (
	"io/ioutil"
	"mime/multipart"
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



