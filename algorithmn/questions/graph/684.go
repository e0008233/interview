package graph

// https://www.youtube.com/watch?v=FXWRE67PLL0
func FindRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	rank := make([]int, len(edges)+1)

	for i := 1; i <= len(edges); i++ {
		parent[i] = i
		rank[i] = 1
	}

	for _, v := range edges {
		if !union(v[0], v[1], &parent, &rank) {
			return v
		}
	}

	return nil

}

func findParent(n int, parent *[]int) int {
	p := (*parent)[n]

	for p != (*parent)[p] {
		(*parent)[p] = (*parent)[(*parent)[p]]
		p = (*parent)[p]
	}

	return p
}

func union(n1, n2 int, parent, rank *[]int) bool {
	p1 := findParent(n1, parent)
	p2 := findParent(n2, parent)

	if p1 == p2 {
		return false
	}

	r1 := (*rank)[n1]
	r2 := (*rank)[n2]

	if r1 > r2 {
		(*parent)[p2] = p1
		(*rank)[n1] = r1 + r2
	} else {
		(*parent)[p1] = p2
		(*rank)[n2] = r1 + r2
	}
	return true
}
