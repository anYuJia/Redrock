package t

import (
	"log"
	"sql/sql"
	"strconv"
	"time"
)

func getDate() string {
	year1 := time.Now().Year()
	month1 := time.Now().Format("01")
	day1 := time.Now().Day()
	return strconv.Itoa(year1) + "-" + month1 + "-" + strconv.Itoa(day1)
}
func GetDate() string {
	return getDate()
}
func getTime() string {
	hour := time.Now().Hour()     //小时
	minute := time.Now().Minute() //分钟
	second := time.Now().Second()
	return strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
}
func GetTime() string {
	return getTime()
}
func getAid(title string) int {
	row, err := sql.Getaid().Query(title)
	if err != nil {
		log.Println(err)

	}
	var aid int
	for row.Next() {
		err = row.Scan(&aid)
		if err != nil {
			log.Println(err)

		}
	}
	return aid
}
func GetAid(title string) int {
	return getAid(title)
}
