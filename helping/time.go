package helping

import (
	"fmt"
	"time"
)

func GetNowTime() string {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	return fmt.Sprintf("%d-%02d-%02d|%02d:%02d:%02d", year, month, day, hour, minute, second)
}
