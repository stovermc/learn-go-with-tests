package iteration

func Repeat(character string, iterationTimes int) string {
	var repeated string

	for i := 0; i < iterationTimes; i++ {
		repeated = repeated + character
	}
	return repeated
}
