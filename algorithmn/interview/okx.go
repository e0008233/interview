package interview

/*
Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
Note that the same word in the dictionary may be reused multiple times in the segmentation.
*/

// wordDict := []string{"ok", "exg"} // 2
//
//	s := "okexg"
// wordBreak checks if a given string can be segmented into one or more dictionary words.
func wordBreak(s string, wordDict []string) bool {
	wordLength := make(map[int]bool)
	wordDictMap := make(map[string]bool)

	dp := make([]bool, len(s))
	maxWordLength := 0
	for _, v := range wordDict {
		wordDictMap[v] = true
		length := len(v)
		if length > maxWordLength {
			maxWordLength = length
		}
		_, ok := wordLength[length]
		if !ok {
			wordLength[length] = true
		}
	}

	for i := 0; i < len(s); i++ {
		for k, _ := range wordLength {
			// k is the length of the word in the dict
			newWordingStarting := i + 1 - k

			if (newWordingStarting-1 >= 0 && dp[newWordingStarting-1]) || newWordingStarting == 0 {
				newWord := s[newWordingStarting : i+1]

				_, ok := wordDictMap[newWord]
				if ok {
					// new word in the wordDict
					dp[i] = true
					break
				}

			}

		}

	}

	// Your implementation here
	return dp[len(s)-1]
}
