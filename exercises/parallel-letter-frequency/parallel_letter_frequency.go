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
	freqMaps := make(chan FreqMap, 10)
	freqTotal := FreqMap{}

	for _, str := range strings {
		go func(str string) {
			freqMaps <- Frequency(str)
		}(str)
	}

	for range strings {
		for letter, num := range <-freqMaps {
			freqTotal[letter] += num
		}
	}

	return freqTotal
}
