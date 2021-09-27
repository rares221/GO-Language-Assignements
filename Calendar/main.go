package main

import (
	"./Calendar"
)

func main(){

	data := calendar.DataCalendaristica{22,04,2021}
	data1 := calendar.DataCalendaristica{21,04,2021}
	calendar.ShowDate(data)
	calendar.ShowDate(data1)

	calendar.NumberOfDaysBetweenTwoDates(data, data1)



}
