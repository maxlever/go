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
	done := make(chan bool)
	endMap := make(FreqMap)

	go func(strings []string, endMap FreqMap, freqmaps chan FreqMap) {
		for range strings {
			for r, n := range <-freqmaps {
				endMap[r] += n
			}
		}
		done <- true
	}(strings, endMap, freqmaps)

	for i, str := range strings {
		go func(i int, str string) { // producer of freqmaps
			freqmaps <- Frequency(str)
		}(i, str)
	}

	<-done
	return endMap
}
