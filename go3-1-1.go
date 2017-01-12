package hangul

import "fmt"

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func main() {
	for i, r := range "가나다" {
		fmt.Println(i, r)
	}
	fmt.Println(len("가나다"))

	fmt.Println()

	for _, r := range "가갛힣" {
		fmt.Println(string(r), r)
	}
}
