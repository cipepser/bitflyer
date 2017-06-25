package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"./myUtil"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// URL is a end point of bitflyer api.
	URL             = "https://api.bitflyer.jp"
	database string = "bitflyer"
	table    string = "executions"
)

func main() {
	bu := myUtil.BarUnit{
		T: 1,
		// Unit: myUtil.FormatSecond,
		Unit: myUtil.FormatMinute,
		// Unit: myUtil.FormatHour,

		// Unit: myUtil.FormatDay,
		// Unit: myUtil.FormatMonth,
		// Unit: myUtil.FormatYear,
	}

	start := "2017-06-01T12:00"
	end := "2017-06-01T12:20"

	s, err := time.Parse(bu.Unit, start)
	if err != nil {
		log.Fatal(err)
	}
	e, err := time.Parse(bu.Unit, end)
	if err != nil {
		log.Fatal(err)
	}
	d := e.Sub(s)

	n, err := myUtil.GetNofCandle(d, bu)
	if err != nil {
		log.Fatal(err)
	}

	btw, ts := myUtil.GetTimeDuration(s, n, bu)

	db, err := sql.Open("mysql", "root:"+os.Getenv("MYSQL_PASSWD")+"@/"+database)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data [][]float64
	fmt.Println(len(btw) - 1)
	for i := 0; i < len(btw)-1; i++ {
		// query
		q := "select Price from executions where ExecDate > " + "\"" + btw[i] + "\" and ExecDate < \"" + btw[i+1] + "\";"
		rows, err := db.Query(q)
		if err != nil {
			log.Fatal(err)
		}

		// result
		var d []float64
		for rows.Next() {
			var price float64
			err = rows.Scan(&price)
			if err != nil {
				log.Fatal(err)
			}
			d = append(d, price)
		}
		data = append(data, d)
		fmt.Println(i)
	}

	if len(data[0]) == 0 {
		data[0] = append(data[0], 0)
		// log.Fata("the length of data[0] is 0.")
	}

	for i := 1; i < len(data); i++ {
		if len(data[i]) == 0 {
			data[i] = append(data[i], data[i-1][len(data[i-1])-1])
		}
	}

	myUtil.MyCandleChart(ts, data, bu)
}
