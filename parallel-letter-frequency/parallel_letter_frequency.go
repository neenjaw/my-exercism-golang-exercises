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

// ConcurrentFrequency makes Frequency computed concurrently returns the frequencies
// totalled
func ConcurrentFrequency(strs []string) FreqMap {
	ch := make(chan FreqMap, 10)
	for _, s := range strs {
		go func(str string) {
			ch <- Frequency(str)
		}(s)
	}

	totalFreq := FreqMap{}
	for range strs {
		for k, v := range <-ch {
			totalFreq[k] += v
		}
	}
	return totalFreq
}
