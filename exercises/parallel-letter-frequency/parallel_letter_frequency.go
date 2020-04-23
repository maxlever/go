package letter

import "fmt"

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

	go func() {
		for i, str := range strings {
			fmt.Println("started processing string", i)
			for _, letter := range str {
				fmt.Println("sending letter", letter)
				letters <- letter
			}
			if i == len(strings)-1 {
				fmt.Println("finished producing")
				doneProducing <- true
			} else {
				fmt.Println("finished processing string", i)
			}
		}
	}()

	<-doneProducing
	fmt.Println("finished producing letters")

	go func() {
		for {
			letter, more := <-letters
			if more {
				fmt.Println("consuming letter", letter)
				endMap[letter]++
			} else {
				fmt.Println("consumed all letters")
				doneConsuming <- true
			}
		}
	}()
	<-doneConsuming
	return endMap

}
