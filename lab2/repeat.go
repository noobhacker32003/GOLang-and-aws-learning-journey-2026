package tournament
import "strings"

func Repeat(text string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result = result + text
	}
	return result
}

func RepeatBuilder(text string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(text)
	}
	return builder.String()
}