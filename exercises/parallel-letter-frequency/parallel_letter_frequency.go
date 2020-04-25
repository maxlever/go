package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency gets the total frequency of runes in the given strings
func ConcurrentFrequency(strings []string) FreqMap {
	freqmaps := make(chan FreqMap)
	endMap := make(FreqMap)

	for _, str := range strings {
		go func(str string) {
			freqmaps <- Frequency(str)
		}(str)
	}

	for range strings {
		for r, n := range <-freqmaps {
			endMap[r] += n
		}
	}

	return endMap
}
