package util

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

type file struct {}

func NewFile()*file{
	return &file{}
}

//文件相关的操作

//获取文件大小
func (this *file)GetFileSize(file multipart.File)(int, error){
	content, err := ioutil.ReadAll(file)
	return len(content), err
}

//获取文件的扩展名
func (this *file)GetExt(file string)string{
	return path.Ext(file)
}

func (this *file)CheckExist(file string)bool{
	_, err := os.Stat(file)
	return os.IsExist(err)
}

//检查权限是否足够
func (this *file)CheckPermission(file string)bool{
	_, err := os.Stat(file)
	return os.IsPermission(err)
}

//不存在则创建
func (this *file)IsNotExistMkDir(file string) error{
	if exists := this.CheckExist(file); exists == false {
		if err := this.MkDir(file); err != nil{
			return err
		}
	}
	return nil
}

//创建指定目录
func (this *file)MkDir(file string) error{
	err := os.MkdirAll(file, os.ModePerm)
	if err != nil{
		return err
	}
	return nil
}

func (this *file)Open(fileName string, flag int, perm os.FileMode)(*os.File, error){
	f, err := os.OpenFile(fileName, flag, perm)
	if err != nil{
		return nil, err
	}
	return f, nil
}

//导出到csv
func (this *file)ExportToCsv(exportPath string, fileName string, data [][]string) error {
	if exportPath == ""{
		return fmt.Errorf("导出目录名错误")
	}
	err := this.IsNotExistMkDir(exportPath)
	if err != nil {
		return nil
	}
	f, err := os.Create(exportPath + fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, _ = f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	_ = w.WriteAll(data)
	return nil
}

//读取文件的字节数据
func (this *file)ReadContent(fileName string) ([]byte, error) {
	fd, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(fd)
	_ = fd.Close()
	return contents, err
}


//读取文件数据，返回字符串
func (this *file)GetContentString(fileName string)(string, error){
	contents, err := this.ReadContent(fileName)
	return string(contents), err
}
