package models

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

type ModelCron struct {
}

func (this *ModelCron)WeeklyIncr(crontab *cron.Cron){


	_ = crontab.AddFunc("0 * * * * *", func() {
		now := time.Now()
		year := now.Year()
		mon := now.Month()
		day := now.Day()
		hour := now.Hour()
		min := now.Minute()
		sec := now.Second()
		d := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", year, mon, day, hour, min, sec)
		fmt.Println(d, "每分钟执行")
	})

	_ = crontab.AddFunc("0 14-59/3 * * * *", func() {//从14分钟开始，之后的每3分钟（14，17，20，23。。。）
		now := time.Now()
		year := now.Year()
		mon := now.Month()
		day := now.Day()
		hour := now.Hour()
		min := now.Minute()
		sec := now.Second()
		d := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", year, mon, day, hour, min, sec)
		fmt.Println(d, "每3分钟执行")
	})
}
