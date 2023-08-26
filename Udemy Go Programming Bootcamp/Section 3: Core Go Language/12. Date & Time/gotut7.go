// https://www.udemy.com/course/go-language/learn/lecture/35216218#notes
//
package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	// ----- TIME -----
	// Get day, month, year and time data
	// Get current time
	now := time.Now()
	pl("------- Here in South Africa--------")
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	// Set a location to get time
	pl("---------In New York---------")
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		pl(err)
	}
	fmt.Printf("Time in New York %s\n", now.In(loc))

	// Change location to Shanghai
	loc, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Printf("Time in Shanghai %s\n", now.In(loc))

	// Get times using different time standards
	pl("--------Time Standards---------")
	locEST, _ := time.LoadLocation("EST")
	locUTC, _ := time.LoadLocation("UTC")
	locMST, _ := time.LoadLocation("MST")
	fmt.Printf("EST : %s\n", now.In(locEST))
	fmt.Printf("UTC : %s\n", now.In(locUTC))
	fmt.Printf("MST : %s\n", now.In(locMST))

	// Calculate time since birthdate
	// Year, month, day, hour, minute, second
	// nanosecond and time zone
	pl("-------Time Calculations-------")
	birtdate := time.Date(1951, time.June, 2, 12, 30, 10, 0, time.Local)
	pl("Birthday : ", birtdate)

	// Get difference between past date and now
	diff := now.Sub(birtdate)

  // Difference in days
	fmt.Printf("Days alive : %d days\n", int(diff.Hours()/24))
  
  // Difference in hours
	fmt.Printf("Hours alive : %d hours\n", int(diff.Hours()))
}
