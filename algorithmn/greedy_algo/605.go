package greedy_algo

func canPlaceFlowers(flowerbed []int, n int) bool {
	allowed := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] != 1 {
			//check back
			if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
				continue
			}

			if i != 0 && flowerbed[i-1] == 1 {
				continue
			}

			flowerbed[i] = 1
			allowed++
		}

	}

	return n <= allowed
}
