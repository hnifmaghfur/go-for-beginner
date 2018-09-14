package helper

import (
	"time"
	"strconv"
)

func ConvertToEmailDate(t time.Time, format string) string {
	var DayName string
	var Day string
	var Month string
	var Year string
	var HM string

	loc,_  := time.LoadLocation("Asia/Jakarta")
	t 		= t.In(loc)

	//Day Name in String
	switch d:= t.Weekday().String(); d {
		case "Monday":
			DayName = "Senin"
			break
		case "Tuesday":
			DayName = "Selasa"
			break
		case "Wednesday":
			DayName = "Rabu"
			break
		case "Thursday":
			DayName = "Kamis"
			break
		case "Friday":
			DayName = "Jumat"
			break
		case "Saturday":
			DayName = "Sabtu"
			break
		case "Sunday":
			DayName = "Minggu"
			break
	}

	// Day
	Day = strconv.Itoa(t.Day())

	// Month Name in String
	switch m := t.Month().String(); m {
	case "January":
		Month = "Januari"
		break
	case "February":
		Month = "Febuari"
		break
	case "March":
		Month = "Maret"
		break
	case "April":
		Month = "April"
		break
	case "May":
		Month = "Mei"
		break
	case "June":
		Month = "Juni"
		break
	case "July":
		Month = "Juli"
		break
	case "August":
		Month = "Agustus"
		break
	case "September":
		Month = "September"
		break
	case "October":
		Month = "Oktober"
		break
	case "November":
		Month = "November"
		break
	case "December":
		Month = "Desember"
		break
	}

	//Year
	Year = strconv.Itoa(t.Year())

	//Hour and Minute
	HM = t.Format("15:04")

	if format == "time" {
		return string(HM + " WIB")
	}
	if format == "date" {
		return string(DayName + ", " + Day + " " + Month + " " + Year)
	}
	if format == "onlydate" {
		return string(Day + " " + Month + " " + Year)
	}

	return string(DayName + ", " + Day + " " + Month + " " + Year + ", " + HM + " WIB")
}

// Golang unzip from timezone fail
func GetNowTime() time.Time {
	loc, err := GetLoc()
	if err != nil {
		return time.Now().Add(time.Hour * time.Duration(7)) // Add 7 hours from local time
	}

	t := time.Now().In(loc)
	return t
}

func GetLoc() (loc *time.Location, err error) {
	// 1st Try Indonesia
	loc, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		// 2nd Try Thailand
		loc, err = time.LoadLocation("Asia/Bangkok")
		if err != nil {
			// 3rd Try Vietnam
			loc, err = time.LoadLocation("Asia/Saigon")
		}
	}

	return loc, err
}

func TimestampISO8601() string {
	t := GetNowTime().Format("2006-01-02T15:04:03.000-0700")
	return t
}
