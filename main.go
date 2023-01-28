package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		// qateni zhaz
		return
	}
	filename := arguments[0]
	if !strings.HasSuffix(filename, ".txt") {
		fmt.Println("Qate format !!!")
		return
	}
	text, err := os.Open(filename)
	var text1 []string
	if err != nil {
		os.Exit(0)
	}
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		Split := Split(scanner.Text())
		text1 = append(text1, Split...) // Mende bul zherde Bukil stroka berolgen !!!
		text1 = append(text1, "\n")
	}
	if len(text1) != 0 {
		text1 = text1[:len(text1)-1]
	}
	if err != nil {
		log.Fatal(err)
	}
	err = text.Close()
	if err != nil {
		os.Exit(0)
	}
	if scanner.Err() != nil {
		scanner.Err()
	}
	for i := 0; i < len(text1); i++ {
		switch text1[i] {
		// (cap) , (low),(hex) ,(bin)
		case "(cap)":
			san := 1
			if i != 0 {
				if text1[i-san] == "" {
					san++
				}
				text1[i-san] = Cap(text1[i-san])
				text1[i] = ""
			} else {
				text1[i] = ""
			}

		case "(up)":
			san := 1
			if i != 0 {
				if text1[i-san] == "" {
					san++
				}
				text1[i-san] = up(text1[i-san])
				text1[i] = ""
			} else {
				text1[i] = ""
			}
		case "(hex)":
			san := 1
			if i != 0 {
				if text1[i-san] == "" {
					san++
				}
				text1[i-san] = hex(text1[i-san])
				text1[i] = ""
			} else {
				text1[i] = ""
			}
		case "(low)":
			san := 1
			if i != 0 {
				if text1[i-san] == "" {
					san++
				}
				text1[i-san] = low(text1[i-san])
				text1[i] = ""
			} else {
				text1[i] = ""
			}
		case "(bin)":
			san := 1
			if i != 0 {
				if text1[i-san] == "" {
					san++
				}
				text1[i-san] = bin(text1[i-san])
				text1[i] = ""
			} else {
				text1[i] = ""
			}
			// Bul zherde (cap, 2(nemese kez kelgen san qoldanduq))
		case "(cap,":
			a := text1[i+1]
			san := ""
			var flag bool = true
			for j, k := range a {
				if j == len(a)-1 && k == ')' {
					break
				}
				if unicode.IsDigit(k) {
					san += string(k)
				} else {
					flag = false
					break
				}
			}
			if flag == false {
				fmt.Println("kate")
				return
			}
			number, _ := strconv.Atoi(san)
			for j := 1; j <= number; j++ {
				if i-j < 0 {
					break
				}
				if text1[i-j] == "" || text1[i-j] == "\n" {
					number++
					if i-number < 0 {
						break
					}
					continue
				}
				text1[i-j] = Cap(text1[i-j])
			}
			// }
			text1[i] = ""
			text1[i+1] = ""
		case "(low,":
			a := text1[i+1]
			san := ""
			var flag bool = true
			for j, k := range a {
				if j == len(a)-1 && k == ')' {
					break
				}
				if unicode.IsDigit(k) {
					san += string(k)
				} else {
					flag = false
					break
				}
			}
			if flag == false {
				fmt.Println("kate")
				return
			}
			number, _ := strconv.Atoi(san)

			for j := 1; j <= number; j++ {
				if i-j < 0 {
					break
				}
				if text1[i-j] == "" || text1[i-j] == "\n" {
					number++
					if i-number < 0 {
						break
					}
					continue
				}
				text1[i-j] = low(text1[i-j])
			}
			text1[i] = ""
			text1[i+1] = ""
		case "(up,":
			a := text1[i+1]
			san := ""
			var flag bool = true
			for j, k := range a {
				if j == len(a)-1 && k == ')' {
					break
				}
				if unicode.IsDigit(k) {
					san += string(k)
				} else {
					flag = false
					break
				}
			}
			if flag == false {
				fmt.Println("kate")
				return
			}
			number, _ := strconv.Atoi(san)

			for j := 1; j <= number; j++ {
				if i-j < 0 {
					break
				}
				if text1[i-j] == "" || text1[i-j] == "\n" {
					number++
					if i-number < 0 {
						break
					}
					continue
				}
				text1[i-j] = up(text1[i-j])
			}
			text1[i] = ""
			text1[i+1] = ""
		case ",", ".", "!", "?", ":", ";":
			pre := 1
			str := ""
			if i == 0 {
				continue
			}
			for str == "" {
				if i-pre >= 0 {
					str = text1[i-pre]
				}
				pre++
			}
			text1[i-pre+1] = Join(str, text1[i])
			text1[i] = ""
		case "a", "an":
			if Istrue(string(text1[i+1][0])) {
				text1[i] = "an"
			} else {
				text1[i] = "a"
			}
		case "A", "An":
			if Istrue(string(text1[i+1][0])) {
				text1[i] = "An"
			} else {
				text1[i] = "A"
			}
		}
	}
	// Result.txt berilgen mandi saqtay ushin !!!
	text1 = DeleterSpace(text1)
	text1 = Apostroph(text1)
	str := Fileds(text1)
	raz := arguments[1]
	if !strings.HasSuffix(raz, ".txt") {
		fmt.Println("Qate format")
		return
	}
	err = os.WriteFile(raz, []byte(str), 0644)
	if err != nil {
		fmt.Println("Error", err.Error)
		return
	}
}

func hex(word string) string {
	num, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		fmt.Println("Qate!!!")
	}
	res := strconv.Itoa(int(num))
	return res
}

func up(word string) string {
	res := strings.ToUpper(word)
	return res
}

func bin(word string) string {
	res, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		fmt.Println("Qate")
	}
	num := strconv.Itoa(int(res))
	return num
}

func low(word string) string {
	res := strings.ToLower(word)
	return res
}

func Join(s string, sep string) string {
	str := s + sep
	return str
}

func Split(word string) []string {
	var res []string
	str := ""
	for _, v := range word {
		if v == ' ' || v == '	' {
			if len(str) > 0 {
				res = append(res, str)
				str = ""
			}
		} else if v == '\n' {
			res = append(res, str+string(v))
			str = ""
		} else if v == ',' || v == '.' || v == ':' || v == ';' || v == '!' || v == '?' || v == '"' {
			str += string(v)
			if len(str) > 0 {
				res = append(res, str)
				str = ""
				continue
			}
		} else {
			str += string(v)
		}
	}
	if str != "" {
		res = append(res, str)
	}

	return res
}

func Istrue(word string) bool {
	v := word
	if v == "a" || v == "e" || v == "u" || v == "i" || v == "o" || v == "h" || v == "A" || v == "E" || v == "U" || v == "I" || v == "O" || v == "H" {
		return true
	}

	return false
}

func Fileds(word []string) string {
	str := ""
	for _, v := range word {
		if v != "" || v == "\n" {
			if v == "\n" {
				if len(str) != 0 {
					str = str[:len(str)-1]
				}
				str += v
			} else {
				str += v + " "
			}
		} else if v != "" {
			str += v
		}
	}
	if len(str) != 0 {
		return str[:len(str)-1]
	}
	return str
}

func Cap(s string) string {
	r := []rune(s)
	h := true
	for i := 0; i < len(r); i++ {
		if r[i] >= 97 && r[i] <= 122 || r[i] >= 65 && r[i] <= 90 || r[i] >= 48 && r[i] <= 57 || r[i] == 39 {
			if h && r[i] >= 97 && r[i] <= 122 {
				r[i] = (r[i] - 32)
			} else if r[i] >= 65 && r[i] <= 90 && !h {
				r[i] = (r[i] + 32)
			}
			h = false
		} else {
			h = true
		}
	}
	return string(r)
}

func Apostroph(ar []string) []string {
	counter := 0
	for i := 0; i < len(ar); i++ {
		val := ar[i]
		if i == len(ar)-1 && counter%2 == 0 {
			break
		}
		if val == "'" {
			if counter%2 == 0 {
				fmt.Println(counter)
				ar[i+1] = val + ar[i+1]
				ar[i] = ""
				i++
			} else {
				ar[i-1] = ar[i-1] + val
				ar[i] = ""
			}
			counter++
			continue
		}
		if len(val) >= 2 {
			if val[0] == '\'' {
				if i == 0 {
					counter++
					continue
				}
				if counter%2 == 0 {
					counter++
					continue
				}
				ar[i-1] = ar[i-1] + "'"
				ar[i] = val[1:]
				counter++
				continue
			} else if val[len(val)-1] == '\'' {
				if counter%2 == 1 {
					counter++
					continue
				}
				ar[i+1] = "'" + ar[i+1]
				ar[i] = val[:len(val)-1]
				counter++
				i++
				continue
			}
		}
	}
	return ar
}

func DeleterSpace(ar []string) []string {
	var newarr []string
	for _, val := range ar {
		if val == "" {
			continue
		}
		newarr = append(newarr, val)
	}
	return newarr
}
