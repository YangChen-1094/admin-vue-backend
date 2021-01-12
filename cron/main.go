package main

import (
	"github.com/robfig/cron"
	"my_gin/models"
	"my_gin/pkg/file"
	"my_gin/pkg/setting"
	"my_gin/pkg/util"
)
func init() {//初始化
	setting.Setup()
	models.Setup()
}

func main(){
	data := [][]string{
		{"1", "test1", "中文1"},
		{"2", "test2", "我时一个安抚"},
		{"3", "test3", "让我发"},
	}
	testFile := util.NewFile()
	path := setting.GetExportPath()
	_ = testFile.ExportToCsv(path, "test.csv", data)
	crontab := cron.New()


	ModelCron := models.NewModelCron()
	ModelCron.WeeklyIncr(crontab)


	crontab.Start()
	defer crontab.Stop()
	//select{}
}
