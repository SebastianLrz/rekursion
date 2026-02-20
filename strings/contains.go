package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	if s == "" {
		return false
	}

	//s ist vond er Form a + rest
	//Falls seq am Anfang steht
	if s[:1] == seq {
		return true
	}

	//Falls seq niucht am Anfang steht
	return Contains(s[1:], seq)
}
