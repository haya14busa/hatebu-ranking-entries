package url

import (
	"fmt"
	"path"
	"time"

	"github.com/haya14busa/hatebu-ranking-entries/category"
)

const (
	PROTOCOL = "http"
	HOST     = "b.hatena.ne.jp"
)

// period
const (
	Daily   = "daily"
	Weekly  = "weekly"
	Monthly = "monthly"
)

const locationStr = "Asia/Tokyo"

var location *time.Location

func init() {
	location = loc()
	time.Local = location
}

func loc() *time.Location {
	loc, err := time.LoadLocation(locationStr)
	if err != nil {
		loc = time.FixedZone(locationStr, 9*60*60)
	}
	return loc
}

func DailyFromCategory(c category.Category, date time.Time) string {
	return genpath(Daily, toDay(date), c.Id())
}

func DailyFromCategoryLatest(c category.Category) string {
	return DailyFromCategory(c, time.Now().AddDate(0, 0, -1))
}

// date must be after 20110103
func WeeklyFromCategory(c category.Category, date time.Time) string {
	return genpath(Weekly, toWeek(date), c.Id())
}

func WeeklyFromCategoryLatest(c category.Category) string {
	return WeeklyFromCategory(c, time.Now().AddDate(0, 0, -7))
}

// date must be after 201101
func MonthlyFromCategory(c category.Category, date time.Time) string {
	return genpath(Monthly, toMonth(date), c.Id())
}

func MonthlyFromCategoryLatest(c category.Category) string {
	n := time.Now()
	d := time.Date(n.Year(), n.Month()-1, 1, 0, 0, 0, 0, location)
	return MonthlyFromCategory(c, d)
}

func toDay(date time.Time) string {
	return date.Format("20060102")
}

func toWeek(date time.Time) string {
	return toMonday(date).Format("20060102")
}

func toMonday(date time.Time) time.Time {
	if date.Weekday() == time.Sunday {
		return date.AddDate(0, 0, -6)
	}
	return date.AddDate(0, 0, int(-date.Weekday()+1))
}

func toMonth(date time.Time) string {
	return date.Format("200601")
}

func genpath(period, date, category string) string {
	path := path.Join(HOST, "ranking", period, date, category)
	return fmt.Sprintf("%s://%s", PROTOCOL, path)
}
