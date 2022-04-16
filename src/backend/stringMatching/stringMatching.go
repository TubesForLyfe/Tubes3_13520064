package string_matching

import "fmt"

func Border(pattern string) []int {
	// return size
	var borderArr []int
	borderArr = append(borderArr, 0)
	i := 1
	j := 0
	size := len(pattern)
	for loop := true; loop; loop = (i < size) {
		// fmt.Println("i = ", i, "c[i] = ", string(pattern[i]))
		// fmt.Println("j = ", j, "c[j] = ", string(pattern[j]))

		if pattern[j] == pattern[i] {
			borderArr = append(borderArr, j+1)
			i++
			j++
			// fmt.Println("1")
		} else if j > 0 {
			j = borderArr[j-1]
			// fmt.Println("2")
		} else {
			borderArr = append(borderArr, 0)
			i++
			// fmt.Println("3")
		}
		// fmt.Println(borderArr)
		// fmt.Println("")
	}

	return borderArr

}

func KMP(pattern string, text string) bool {
	// TODO
	n := len(text)
	m := len(pattern)

	borderArr := Border(pattern)

	i := 0
	j := 0

	for loop := true; loop; loop = (i < n) {
		if pattern[j] == text[i] {
			if j == m-1 { // end of pattern
				return true
			} else {
				i += 1
				j += 1
			}
		} else if j > 0 {
			// mismatch not at first letter
			j = borderArr[j-1]
		} else {
			// mismatch at first letter
			i += 1
		}
	}

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
