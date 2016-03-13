package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/haya14busa/hatebu-ranking-entries/category"
	parser "github.com/haya14busa/hatebu-ranking-entries/parser"
	"github.com/haya14busa/hatebu-ranking-entries/url"
)

var (
	year      int
	month     int
	day       int
	period    string
	noentries bool
)

func init() {
	flag.IntVar(&year, "year", 0, "year")
	flag.IntVar(&month, "month", 0, "month")
	flag.IntVar(&day, "day", 0, "day")
	flag.StringVar(&period, "period", "daily", "period[daily|weekly|monthly]")
	flag.BoolVar(&noentries, "noentries", false, "noentries")
	flag.Parse()
}

func main() {
	now := time.Now()
	date := now.AddDate(-year, -month, -day)

	var hotentry string
	if period == "monthly" {
		lastmonth := now.AddDate(0, -1, 0)
		if date.After(lastmonth) {
			date = lastmonth
		}
		hotentry = url.MonthlyFromCategory(category.Hotentry, date)
	} else if period == "weekly" {
		lastweek := now.AddDate(0, 0, -7)
		if date.After(lastweek) {
			date = lastweek
		}
		hotentry = url.WeeklyFromCategory(category.Hotentry, date)
	} else {
		yesterday := now.AddDate(0, 0, -1)
		if date.After(yesterday) {
			date = yesterday
		}
		hotentry = url.DailyFromCategory(category.Hotentry, date)
	}

	if noentries {
		fmt.Printf("%s\n", hotentry)
		return
	}

	es, err := parser.NewEntries(hotentry)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for _, e := range es {
		fmt.Printf("rank:%d", e.Rank)
		fmt.Printf("\tusers:%d", e.Bookmarkcount)
		fmt.Printf("\ttitle:%s", e.Title)
		fmt.Printf("\turl:%s", e.Url)
		fmt.Printf("\tsource:%s", hotentry)
		fmt.Println("")
	}
}
