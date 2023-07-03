package iteration

import "testing"

const RepeatTimes = 5

func Repeat(singleStr string) string {
	var result string
	for i := 0; i < RepeatTimes; i++ {
		result += singleStr
	}
	return result
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
