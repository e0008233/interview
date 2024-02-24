package back_track

import (
	"fmt"
	"sort"
	"strings"
)

// https://www.youtube.com/watch?v=ZyB_gQ8vqGA

func FindItinerary(tickets [][]string) []string {
	adjList := make(map[string][]string)

	for _, src := range tickets {
		startCity := src[0]
		v, ok := adjList[startCity]
		if ok {
			v = append(v, src[1])
		} else {
			v = []string{src[1]}
		}
		adjList[startCity] = v

	}
	for _, v := range adjList {
		sort.Strings(v)
	}
	cache := make(map[string]bool)
	fmt.Println(adjList)
	result := make([]string, 0)
	result = append(result, "JFK")
	isSuccess := dfs(&result, &adjList, tickets, "JFK", &cache)
	fmt.Println(isSuccess)
	return result
}

func dfs(result *[]string, adjList *map[string][]string, tickets [][]string, src string, cache *map[string]bool) bool {
	if len(*result) == len(tickets)+1 { // all tickets are used
		return true
	}
	key := strings.Join(*result, "")
	if v, ok := (*cache)[key]; ok {
		return v
	}
	dest, ok := (*adjList)[src]
	if !ok {
		(*cache)[key] = false
		return false
	}

	if len(dest) == 0 {
		(*cache)[key] = false
		return false
	}

	temp := make([]string, len(dest))
	copy(temp, dest)
	for i, v := range temp {
		*result = append(*result, v)
		temp2 := make([]string, len(dest))
		temp2 = append(temp[:i], temp[i+1:]...)

		(*adjList)[src] = temp2
		isSuccess := dfs(result, adjList, tickets, v, cache)
		if isSuccess {
			(*cache)[key] = true
			return true
		}
		*result = (*result)[0 : len(*result)-1]
		(*adjList)[src] = dest
		copy(temp, dest)

	}
	(*cache)[key] = false
	return false
}
