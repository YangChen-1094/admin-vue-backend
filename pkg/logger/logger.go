package logger

import (
	"fmt"
	"log"
	"my_gin/pkg/setting"
	"my_gin/pkg/util"
	"os"
	"runtime"
	"strings"
	"time"
)
type Level int

var (
	FileHandle *os.File
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	DefaultPrefix = ""
	DefaultDepth = 2
)
const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)


func Info(name string, v ...interface{}){
	FileHandle, err := openLogFile(name)
	if err != nil {
		log.Printf("openLogFile() in os.Getwd err:%v", err)
	}
	logger := log.New(FileHandle, DefaultPrefix, log.LstdFlags | log.Lmicroseconds)
	setPrefix(logger, INFO)
	logger.Println(v)
	defer FileHandle.Close()
}

func Warn(name string, v ...interface{}){
	FileHandle, err := openLogFile(name)//log日志目录的相对路径之后的文件全名 如： logs/20060102-logName.log
	if err != nil {
		log.Printf("openLogFile() in os.Getwd err:%v", err)
	}
	logger := log.New(FileHandle, DefaultPrefix, log.LstdFlags | log.Lmicroseconds)
	setPrefix(logger, WARNING)
	logger.Println(v)
	defer FileHandle.Close()
}

func setPrefix(thisLog *log.Logger,level Level){
	_, file, line, ok := runtime.Caller(DefaultDepth)//调用记录日志的文件信息
	baseDir, _ := os.Getwd()//
	baseDir = strings.Replace(baseDir, "\\", "/", -1)//可执行文件之前的路径

	relativeFile := strings.Replace(file, baseDir + "/", "", -1) //只记录相对路径
	var logPrefix string
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d] ", levelFlags[level], relativeFile, line)
	}else{
		logPrefix = fmt.Sprintf("[%s] ", levelFlags[level])
	}
	thisLog.SetPrefix(logPrefix)
}


//log日志目录的相对路径 如： logs/
func getLogPath() string{
	return fmt.Sprintf("%s", setting.DeployConfig.App.LogSavePath)
}

//打开日志
func openLogFile(fileName string)(*os.File, error){
	file := util.NewFile()
	dir, err:= os.Getwd()//运行目录
	if err != nil {
		return nil, fmt.Errorf("openLogFile() in os.Getwd err:%v", err)
	}
	logsPath := getLogPath() // ...logs/
	Ymd := time.Now().Format(setting.DeployConfig.App.LogTimeFormat)
	logsPath = dir + "/" + logsPath + fileName + "/" //logs/logName/20200824.log

	err = file.IsNotExistMkDir(logsPath)
	if err != nil {
		return nil, fmt.Errorf("openLogFile() in IsNotExistMkDir err:%v", err)
	}

	finalName := logsPath + Ymd //logs/20060102/logName.log
	f, err := file.Open( finalName + "." + setting.DeployConfig.App.LogFileExt, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0664)
	if err != nil {
		return nil, fmt.Errorf("openLogFile() file.Open err:%v", err)
	}
	return f, nil
}