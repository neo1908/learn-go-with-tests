package src

// Repeat Can be replaced with strings.Repeat()
func Repeat(character string, count int) string {
	var repeated string

	// There is only "for()" no "while" etc
	for i := 0; i < count; i++ {
		repeated += character
	}

	return repeated
}
