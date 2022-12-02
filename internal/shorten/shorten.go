package shorten

import "strings"

const a = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ0987654321"

var alen = len(a)

func Shorten(id int) string {
	var (
		nums    []int
		num     = id
		builder strings.Builder
	)
	//panic("not implemented")

	for num > 0 {
		nums = append(nums, num%alen)
		num /= alen
	}
	Reverse(nums)

	for _, num := range nums {
		builder.WriteString(string(a[num]))
	}
	return builder.String()
}

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
