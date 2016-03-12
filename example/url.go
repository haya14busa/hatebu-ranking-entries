package main

import (
	"fmt"
	"time"

	"github.com/haya14busa/hatebu-ranking-entries/category"
	"github.com/haya14busa/hatebu-ranking-entries/url"
)

func main() {
	now := time.Now()

	fmt.Println(url.DailyFromCategoryLatest(category.Hotentry))
	fmt.Println(url.DailyFromCategoryLatest(category.It))
	fmt.Println(url.DailyFromCategory(category.Hotentry, now.AddDate(0, 0, -7)))
	fmt.Println(url.DailyFromCategory(category.Hotentry, now.AddDate(0, -1, 0)))
	fmt.Println(url.DailyFromCategory(category.Hotentry, now.AddDate(-1, 0, 0)))
	fmt.Println("===weekly")
	fmt.Println(url.WeeklyFromCategory(category.Hotentry, now.AddDate(0, 0, -7)))
	fmt.Println(url.WeeklyFromCategoryLatest(category.It))
	fmt.Println(url.WeeklyFromCategory(category.It, time.Date(2011, 1, 8, 0, 0, 0, 0, time.UTC)))
	fmt.Println("===monthly")
	fmt.Println(url.MonthlyFromCategory(category.Hotentry, now.AddDate(-1, 0, 0)))
	fmt.Println(url.MonthlyFromCategoryLatest(category.It))
	fmt.Println(url.MonthlyFromCategory(category.It, now.AddDate(-5, 0, 0)))
}
