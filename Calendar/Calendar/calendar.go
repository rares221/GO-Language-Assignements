package calendar

import "fmt"

type DataCalendaristica struct {
	Day, Month, Year int
}

func Compare(d1, d2 DataCalendaristica) int {
	if d1.Year < d2.Year {
		return -1
	}
	if d1.Year > d2.Year {
		return 1
	}

	if d1.Month < d2.Month {
		return -1
	}

	if d1.Month > d2.Month {
		return 1
	}

	if d1.Day < d2.Day {
		return -1
	}

	if d1.Day > d2.Day {
		return 1
	}

	return 0
}

func (c DataCalendaristica) IsBisect() bool {
	if c.Year%4 == 0 {
		if c.Year%100 == 0 {
			if c.Year%400 == 0 {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func (c DataCalendaristica) PlusOneDay() DataCalendaristica {
	numberOfDays := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if c.Day < numberOfDays[c.Month] {
		c.Day++
	} else if c.Day == numberOfDays[c.Month] {
		if c.IsBisect() && c.Month == 2 {
			c.Day++
		} else if c.Month == 12 {
			c.Day = 1
			c.Month = 1
			c.Year++
		} else {
			c.Day = 1
			c.Month++
		}
	} else if c.Month == 2 && c.IsBisect() {
		c.Day = 1
		c.Month++
	}

	return c
}

func (c DataCalendaristica) MinusOneDay() DataCalendaristica {
	numberOfDays := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if c.Day == 1 {
		if c.IsBisect() && c.Month == 3 {
			c.Day = 29
			c.Month = 2
		} else if c.Month == 1 {
			c.Day = 31
			c.Month = 12
			c.Year--
		} else {
			c.Month--
			c.Day = numberOfDays[c.Month]

		}
	} else if c.Day < numberOfDays[c.Month] {
		c.Day--
	}

	return c
}

func (c DataCalendaristica) PlusOneMonth() DataCalendaristica {
	numberOfDays := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if c.Month < 12 {
		c.Month++
		if c.Day > numberOfDays[c.Month] {

			if c.Month == 2 {

				if c.IsBisect() {
					c.Day = 29
				} else {
					c.Day = 28
				}

			} else {
				c.Day--
			}

		}
	} else {
		c.Month = 1
		c.Year++
	}

	return c
}


func (c DataCalendaristica) MinusOneMonth() DataCalendaristica {
	numberOfDays := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if c.Month > 1 {
		c.Month--
		if c.Day > numberOfDays[c.Month] {

			if c.Month == 2 {

				if c.IsBisect() {
					c.Day = 29
				} else {
					c.Day = 28
				}

			} else {
				c.Day =numberOfDays[c.Month]
			}

		}
	} else {
		c.Month = 12
		c.Year--
	}

	return c
}



func (c DataCalendaristica) PlusOneYear() DataCalendaristica {
	c.Year++
	return c
}

func (c DataCalendaristica) MinusOneYear() DataCalendaristica {
	c.Year--
	return c
}

func NumberOfDaysBetweenTwoDates(date1, date2 DataCalendaristica) int {
	numberOfDaysPerMonth := [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	var biggerDate DataCalendaristica
	var smallerDate DataCalendaristica
	if date1.Year > date2.Year {
		biggerDate = date1
		smallerDate = date2
	} else if date2.Year > date1.Year {
		biggerDate = date2
		smallerDate = date1
	} else if date1.Month > date2.Month {
		biggerDate = date1
		smallerDate = date2
	} else if date2.Month > date1.Month {
		biggerDate = date2
		smallerDate = date1
	} else if date1.Day > date2.Day {
		biggerDate = date1
		smallerDate = date2
	} else if date2.Day > date1.Day {
		biggerDate = date2
		smallerDate = date1
	} else {
		return 0
	}


	yearDifference := biggerDate.Year - smallerDate.Year
	var numberOfDaysForYears = 0
	if biggerDate.IsBisect()||smallerDate.IsBisect() {
		numberOfDaysForYears = yearDifference*366
	}else {
		numberOfDaysForYears = yearDifference*365
	}
	monthDifference := biggerDate.Month - smallerDate.Month

	zileBD :=0
	zileSD := numberOfDaysPerMonth[smallerDate.Month] - smallerDate.Day
	if biggerDate.Month!=smallerDate.Month{
		zileBD =  biggerDate.Day
	}

	var daysBetween = 0
	if monthDifference>1{
		for i :=1 ;i<monthDifference; i++{
			daysBetween +=numberOfDaysPerMonth[smallerDate.Month+i]
		}
	}


	dayDifference := daysBetween+zileSD+zileBD


	numberOfDays := dayDifference +  numberOfDaysForYears

	return numberOfDays
}

func ShowDate(date DataCalendaristica) {
	fmt.Println(date.Day, "/", date.Month, "/", date.Year)
}

func ValidateDate(date DataCalendaristica) bool {
	ok := true
	if date.Month == 2 && date.Day > 29 {
		ok = false
	}

	return ok
}
