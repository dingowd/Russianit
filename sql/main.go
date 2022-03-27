package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// get date in "YYYY-MM-DD" format
func GetDate(y, m, d int) string {
	var mm, dd string
	if m < 10 {
		mm = fmt.Sprint("0", m)
	} else {
		mm = fmt.Sprint(m)
	}
	if d < 10 {
		dd = fmt.Sprint("0", d)
	} else {
		dd = fmt.Sprint(d)
	}
	return fmt.Sprint(y) + "-" + mm + "-" + dd
}

// compose query
func ComposeQuery(days, year, month int) string {
	query := "select Сотрудник, "
	for i := 0; i < days-1; i++ {
		query += "max(q" + strconv.Itoa(i) + ") as \"" + strconv.Itoa(i+1) + "." + strconv.Itoa(month) + "\", "
	}
	query += "max(q" + strconv.Itoa(days-1) + ") as \"" + strconv.Itoa(days) + "." +
		strconv.Itoa(month) + "\"\n" + "from\n(\nselect EmployeeID as Сотрудник,\n"
	for i := 0; i < days-1; i++ {
		query += "if(DATE(StartPeriod) = '" + GetDate(year, month, i+1) + "' or DATE(EndPeriod) = '" +
			GetDate(year, month, i+1) + "', '+', '') as q" + strconv.Itoa(i) + ",\n"

	}
	query += "if(DATE(StartPeriod) = '" + GetDate(year, month, days) + "' or DATE(EndPeriod) = '" +
		GetDate(year, month, days) + "', '+', '') as q" + strconv.Itoa(days-1) + "\n" +
		"from timework\n) as query1\ngroup by Сотрудник;"

	return query
}

func main() {

	// enter year and month to check timesheet
	var year, month int
	fmt.Print("Year:")
	fmt.Scan(&year)
	fmt.Print("Month:")
	fmt.Scan(&month)

	// determine the number of days per month
	days := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()

	// compose query
	query := ComposeQuery(days, year, month)

	// write query to file
	file, err := os.Create("query.sql")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(query)
}
