package main

func Repeat(str string, repeat int) string {
	var repeatSize = repeat

	toReturn := ""
	// could swap out the 5 to a constant
	// ie. const repeatCount int = 5
	for i := 0; i < repeatSize; i++ {
		toReturn += str
	}
	return toReturn
}
