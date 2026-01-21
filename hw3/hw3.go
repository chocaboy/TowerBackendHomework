package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func rez_words(line string, n int) string {
	if n == 0 {
		return line
	}
	flag1 := true
	index := len(line)
	for i := 0; i < len(line); i++ {
		if n == -1 {
			break
		}
		if flag1 && line[i] != 9 && line[i] != 32 {
			flag1 = false
			n--
			index = i
		}
		if line[i] == 9 || line[i] == 32 {
			flag1 = true
		}
	}
	return line[index:]
}

func key_c(file []string) []string {
	result := []string{}
	if len(file) == 0 {
		return result
	}
	var temp int
	temp = 1
	last := file[0]
	for i := 1; i < len(file); i += 1 {
		if file[i] != last {
			result = append(result, strconv.Itoa(temp)+" "+last)
			temp = 1
		} else {
			temp = temp + 1
		}
		last = file[i]
	}
	result = append(result, strconv.Itoa(temp)+" "+last)
	return result
}

func key_d(file []string) []string {
	result := []string{}
	if len(file) == 0 {
		return result
	}
	temp := 1
	last := file[0]
	for i := 1; i < len(file); i++ {
		if file[i] == last {
			temp++
		} else {
			if temp != 1 {
				result = append(result, last)
			}
			temp = 1
		}
		last = file[i]
	}
	if temp != 1 {
		result = append(result, last)
	}
	return result
}

func key_u(file []string) []string {
	result := []string{}
	if len(file) == 0 {
		return result
	}
	temp := 1
	last := file[0]
	for i := 1; i < len(file); i++ {
		if file[i] == last {
			temp++
		} else {
			if temp == 1 {
				result = append(result, last)
			}
			temp = 1
		}
		last = file[i]
	}
	if temp == 1 {
		result = append(result, last)
	}
	return result
}

func key_f(file []string, n int) []string {
	result := []string{}
	if n == 0 {
		return file
	}
	if len(file) == 0 {
		return result
	}
	result = append(result, file[0])
	last := rez_words(file[0], n)
	for i := 1; i < len(file); i++ {
		temp := rez_words(file[i], n)
		if temp != last {
			result = append(result, file[i])
		}
		last = temp
	}
	return result
}
func key_s(file []string, n int) []string {
	result := []string{}
	if n == 0 {
		return file
	}
	if len(file) == 0 {
		return result
	}
	j := 1
	var last string
	for {
		if n < len(file[j]) {
			last = file[j][n:]
			break
		}
		j++
	}
	if last == "" {
		return result
	}
	result = append(result, file[0])
	for i := j; i < len(file); i++ {
		temp := file[i]
		if n < len(file[i]) {
			temp = temp[n:]
			if temp != last {
				result = append(result, file[i])
			}
		}
		last = temp
	}
	return result
}

func key_i(file []string) []string {
	result := []string{}
	if len(file) == 0 {
		return result
	}
	result = append(result, file[0])
	last := strings.ToLower(file[0])
	start := file[0]
	for i := 1; i < len(file); i++ {
		if last != strings.ToLower(file[i]) {
			start = file[i]
			result = append(result, start)
			last = strings.ToLower(file[i])
		}
	}
	return result
}

func key(file []string) []string {
	result := []string{}
	if len(file) == 0 {
		return result
	}
	result = append(result, file[0])
	last := file[0]
	for i := 1; i < len(file); i++ {
		if last != file[i] {
			result = append(result, file[i])
		}
		last = file[i]
	}
	return result
}

func main() {
	flagc := flag.Bool("c", false, "Подсчёт числа строк")
	flagd := flag.Bool("d", false, "Повторяющиеся строки")
	flagu := flag.Bool("u", false, "Не повторяющиеся строки")
	flagi := flag.Bool("i", false, "Не учитывать индекс")
	flagf := flag.Int("f", 0, "Не учитывать первые n полей")
	var flags = flag.Int("s", 0, "Не учитывать первые n строк")
	flag.Parse()

	var input io.Reader
	if fname := flag.Arg(0); fname != "" {
		file, err := os.Open(fname)
		if err != nil {
			println("Error:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}
	file := []string{}
	buf := bufio.NewScanner(input)
	for {
		if !buf.Scan() {
			break
		}
		file = append(file, buf.Text())
	}
	if (*flagc && *flagu) || (*flagd && *flagu) || (*flagc && *flagd) {
		fmt.Println("Ключи не совместимы")
	} else if *flagf < 0 || *flags < 0 {
		fmt.Println("Введено неверное значение для n")
	} else {
		base := key_s(key_f(file, *flagf), *flags)
		if *flagi {
			if *flagc {
				fmt.Println(key_i(base))
				file = key_c(key_i(base))
			} else if *flagd {
				file = key_d(key_i(base))
			} else if *flagu {
				file = key_u(key_i(base))
			} else {
				file = key(key_i(base))
			}
		} else {
			if *flagc {
				file = key_c(base)
			} else if *flagd {
				file = key_d(base)
			} else if *flagu {
				file = key_u(base)
			} else {
				file = key(base)
			}
		}
	}

	if flag.NArg() > 1 {
		output, err := os.Create(flag.Arg(1))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer output.Close()
		for _, line := range file {
			fmt.Fprintln(output, line)
		}
	} else {
		fmt.Println("---------------")
		for i := 0; i < len(file); i++ {
			fmt.Println(file[i])
		}
	}

}
