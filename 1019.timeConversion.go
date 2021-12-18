package main

import "fmt"

func main() {
	var secs, hours, mins int
	fmt.Scanf("%d", &secs)
	for secs > 60 {
		if secs >= 3600 {
			hours = secs / 3600
			secs -= (hours * 3600)
		} else if secs >= 60 {
			mins = secs / 60
			secs -= (mins * 60)
		}
	}
	fmt.Printf("%d:%d:%d\n", hours, mins, secs)
}
