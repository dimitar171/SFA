package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	dates := []string{"Sep-14-2008", "Dec-03-2021", "Mar-18-2022"}
	format := "Jan-02-2006"
	fmt.Print(sortDates(format, dates...))
}

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

func sortDates(format string, dates ...string) ([]string, error) {
	var sl timeSlice = []time.Time{}
	//converting string to time.Time
	for _, date := range dates {
		time, err := time.Parse(format, date)
		sl = append(sl, time)
		if err != nil {
			fmt.Println(err)
		}
	}
	//sorting
	sort.Sort(sl)
	//converting time.Time to string
	var result []string
	for _, s := range sl {
		result = append(result, s.Format(format))
	}

	return result, nil
}
