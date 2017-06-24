package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// URL is a end point of bitflyer api.
	URL             = "https://api.bitflyer.jp"
	database string = "bitflyer"
	table    string = "executions"
)

func main() {
	// c, _ := sdk.NewClient(URL, "user", "passwd", nil)
	//
	// es := c.GetExecutions("FX_BTC_JPY", "", "", "")
	// lid := es[0].ID
	//
	// for {
	// 	for i := 0; i < len(es); i++ {
	// 		fmt.Println(strconv.FormatFloat(es[len(es)-1-i].ID, 'f', 0, 64), ": ", es[len(es)-1-i].Price, "円 x ", es[len(es)-1-i].Size)
	// 	}
	//
	// 	time.Sleep(5 * time.Second)
	// 	lid = es[0].ID
	// 	es = c.GetExecutions("FX_BTC_JPY", "", "", strconv.FormatFloat(lid, 'f', 0, 64))
	// }

	// tmp := [4]float64{1, 2, 3, 4}

	// x := [][4]float64{tmp, tmp}
	// ts := []string{"2017-06-17 15:34:47.930000", "2017-06-17 15:35:47.930000", "2017-06-17 15:36:47.930000", "aa", "aaa", "aaa"}

	db, err := sql.Open("mysql", "root:"+os.Getenv("MYSQL_PASSWD")+"@/"+database)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var data [][]float64
	n := 1

	for i := 0; i < n; i++ {
		// query
		// q := "select Price from element where n = " + strconv.Itoa(n) + ";"
		q := "select Price from executions where ExecDate > " + "\"2017-06-24T13:18:22\";"
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
			// fmt.Println(price)
		}
		data = append(data, d)
	}

	fmt.Println(data)
	// time.Sleep(1 * time.Second)

	// ts := []string{"15:34", "35", "36", "37", "38", "39"}

	// data := [][]float64{{2, 2, 3, 3, 2, 1, 4, 2, 5, 3}, {4, 2, 6, 0, 2, 1}, {5, 2, 3, 3, 6, 5, 4},
	// {3, 2, 1, 4}, {2, 0.7, 3}, {2, 4, 1.5}}

	// myUtil.MyCandleChart(ts, data)
	// myUtil.MyBoxPlot()

}
