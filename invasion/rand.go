package invasion

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// randStrings returns a list of n random strings of length l
func (m *Map) randStrings(n, l int) []string {
	cityNames := make([]string, 0, n)
	for i := 0; i < n; i++ {
		cityNames = append(cityNames, m.randStringOfLength(l))
	}
	return cityNames
}

// randBytes returns a list of n random strings of length l
func (m *Map) CityBytes(n, l int) [][]byte {
	cityNames := make([][]byte, n)
	for i := 0; i < n; i++ {
		cityNames[i] = m.randBytesOfLength(l)
	}
	return cityNames
}

// Generate a random string of a particular length
func (m *Map) randBytesOfLength(n int) []byte {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, m.rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = m.rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

// Generate a random string of a particular length
func (m *Map) randStringOfLength(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, m.rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = m.rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
