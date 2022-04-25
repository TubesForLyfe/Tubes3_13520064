package string_matching

import "fmt"
import "regexp"

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

func max(a int, b int) int{
	if a > b {
		return a;
	} else {
		return b;
	}
}


func Lcs(str1 string, str2 string) int {
	m := len(str1)
	n := len(str2)
	L := make([][]int, m+1)      // Make the outer slice and give it size 10
	for i := 0; i < m+1; i ++ {
		L[i] = make([]int, n+1)  // Make one inner slice per iteration and give it size 10
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				L[i][j] = 0;
			} else if str1[i-1] == str2[j-1] {
				L[i][j] = L[i-1][j-1] + 1
			} else {
				L[i][j] = max(L[i-1][j], L[i][j-1])
			}
		}
	}

	var percentage float64;
	percentage = float64(L[m][n]) / float64(n)

	return int(percentage * 100);
}

func Regex(str1 string) bool {

	if (len(str1) == 0) {
		return false
	}

	var regex, err = regexp.Compile(`^[AGTC]*$`)

	if err != nil {
		fmt.Println(err.Error())
	}

	return regex.MatchString(str1) 
}