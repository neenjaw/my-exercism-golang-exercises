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
	ch := make(chan FreqMap)
	for _, s := range strs {
		go func(str string) { ch <- Frequency(str) }(s)
	}

	var totalFreq FreqMap
	for i := range strs {
		freq := <-ch
		if i == 0 {
			totalFreq = freq
			continue
		}

		for k, v := range freq {
			totalFreq[k] = totalFreq[k] + v
		}
	}
	return totalFreq
}
