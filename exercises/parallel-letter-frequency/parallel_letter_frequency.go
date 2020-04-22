package letter

import "sync"

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

// SyncFrequency counts the frequency of each rune in a given text
// and mutates the given sync.Map
// sideeffect: decrements given wait group
func SyncFrequency(str string, waitGroup *sync.WaitGroup, m *sync.Map) {
	for _, letter := range str {
		oldValue, _ := m.LoadOrStore(letter, 0)
		m.Store(letter, oldValue.(int)+1)
	}
	waitGroup.Done()
}

// ConcurrentFrequency gets the total frequency of runes in the given strings
func ConcurrentFrequency(strings []string) FreqMap {
	endMap := make(FreqMap)
	concurrentMap := sync.Map{}
	waitGroup := sync.WaitGroup{}
	for _, str := range strings {
		waitGroup.Add(1)
		go SyncFrequency(str, &waitGroup, &concurrentMap)
	}
	waitGroup.Wait()
	concurrentMap.Range(func(k interface{}, v interface{}) bool {
		endMap[k.(rune)] = v.(int)
		return true
	})
	return endMap
}
