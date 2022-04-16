package string_matching

import "fmt"

func KMP() bool {
	// TODO
	return false
}

func BoyerMoore(str1 string, str2 string) bool {
	if len(str1) < len(str2) {
		return false
	} else {
		var i, j, k, l, count int
		var flag bool = false
		i = len(str2) - 1
		count = 0

		for !flag && i < len(str1) {
			j = len(str2) - 1
			k = i
			for j >= 0 && str1[k] == str2[j] {
				k--
				j--
				count++
			}
			count++

			if j == -1 {
				count--
				flag = true
			} else {
				l = len(str2) - 1
				for l >= 0 && str1[k] != str2[l] {
					l--
				}

				if l != -1 {
					if l < j {
						// CASE 1
						i += (j - l)
					} else {
						// CASE 2
						i += 1
					}
				} else {
					// CASE 3
					i += len(str2)
				}
			}
		}
		fmt.Println(count)
		return flag
	}
}

func Regex() bool {
	// TODO
	return false
}