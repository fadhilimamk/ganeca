package global

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ObjectToJSON convert object to json
func ObjectToJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

// GetCurrentTime return current time in unix format
func GetCurrentTime() int64 {
	t := time.Now()
	return t.Unix()
}

// DateStringToTime convert date string in Bahasa Indonesia into time.
// string that can be converted is simple date with with format dd MMM YYY
// example : 25 September 1997
func DateStringToTime(date string) time.Time {
	date = strings.ToLower(date)
	date = strings.TrimSpace(date)
	date = RemoveDuplicateSpaceInString(date)
	data := strings.Split(date, " ")

	day, _ := strconv.Atoi(data[0])
	year, _ := strconv.Atoi(data[2])
	month := time.January
	switch data[1] {
	case "januari":
		month = time.January
	case "februari":
		month = time.February
	case "maret":
		month = time.March
	case "april":
		month = time.April
	case "mei":
		month = time.May
	case "juni":
		month = time.June
	case "juli":
		month = time.July
	case "agustus":
		month = time.August
	case "september":
		month = time.September
	case "oktober":
		month = time.October
	case "november":
		month = time.November
	case "desember":
		month = time.December
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

// CompleteDateStringToTime convert date string in Bahasa Indonesia into time.
// example : Kamis, 25 September 1997, 18:08:52
func CompleteDateStringToTime(date string) time.Time {
	rawCompleteDate := strings.Split(date, ",")
	result := DateStringToTime(rawCompleteDate[1])

	rawClock := strings.TrimSpace(rawCompleteDate[2])
	rawClockData := strings.Split(rawClock, ":")

	hour, _ := strconv.Atoi(rawClockData[0])
	minute, _ := strconv.Atoi(rawClockData[1])
	second, _ := strconv.Atoi(rawClockData[2])

	result.Add(time.Hour*time.Duration(hour) +
		time.Minute*time.Duration(minute) +
		time.Second*time.Duration(second))

	result = time.Date(result.Year(), result.Month(), result.Day(), hour, minute, second, 0, result.Location())

	return result
}

// RemoveDuplicateSpaceInString normalize string with redundant space.
// Example to convert "Ini     adalah    kalimat" into "Ini adalah kalimat"
func RemoveDuplicateSpaceInString(input string) string {
	reLeadcloseWhtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := reLeadcloseWhtsp.ReplaceAllString(input, "")
	final = reInsideWhtsp.ReplaceAllString(final, " ")
	return final
}

// Paginate : chunk interface data with offset and limit
func Paginate(x []interface{}, offset int, limit int) []interface{} {
	if offset > len(x) {
		offset = len(x)
	}

	end := offset + limit
	if end > len(x) {
		end = len(x)
	}

	return x[offset:end]
}
