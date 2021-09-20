package iteration

func Repeat(s string) (result string) {
	for i := 0; i < 5; i++ {
		result += s
	}
	return
}
