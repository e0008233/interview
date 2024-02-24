package dp

// https://www.youtube.com/watch?v=HAA8mgxlov8

type key struct {
	i int
	j int
}

func isMatch(s string, p string) bool {
	cache := make(map[key]bool)
	return recur(0, 0, s, p, &cache)
}

func recur(i int, j int, s string, p string, cache *map[key]bool) bool {
	if i >= len(s) && j >= len(p) { // both checking index reach the end, all matched
		return true
	}
	if j >= len(p) { // if input string reaches the end
		return false
	}

	k := key{i: i, j: j} // create a key instance
	//if v, ok := (*cache)[k]; ok {
	//	return v // return cached result if available
	//}

	match := i < len(s) && (s[i] == p[j] || p[j] == '.')
	result := false
	// if the following in the patter is a *, it can be re-written to "",  e.g. "a*" -> ""
	if j+1 < len(p) && p[j+1] == '*' {
		// choice 1: do not use the "*"
		// choice 2: use the "*" if the character is matched
		result = recur(i, j+2, s, p, cache) || (match && recur(i+1, j, s, p, cache))
		(*cache)[k] = result
		return result
	}
	if match {
		result = recur(i, j+2, s, p, cache) || (match && recur(i+1, j, s, p, cache))
		(*cache)[k] = result
		return result
	}
	(*cache)[k] = result
	return result
}
