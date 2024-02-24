package main

import (
	"fmt"
	"interview/algorithmn/questions/back_track"
)

func main() {

	fmt.Println(back_track.FindItinerary([][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}))

}
