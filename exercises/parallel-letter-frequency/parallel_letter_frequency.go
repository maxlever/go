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
	doneProducing := make(chan bool)
	doneConsuming := make(chan bool)
	letters := make(chan rune)
	endMap := make(FreqMap)

	go func() { // consumer of letters
		for {
			letter, more := <-letters
			if more {
				endMap[letter]++
			} else {
				doneConsuming <- true
			}
		}
	}()

	go func() { // producer of letters
		for _, str := range strings {
			for _, letter := range str {
				letters <- letter
			}
		}
		doneProducing <- true
		close(letters)
	}()

	<-doneProducing
	<-doneConsuming

	return endMap
}
