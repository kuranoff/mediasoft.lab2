package preproc

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r)
}

// Normalize ...
func Normalize(line string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, line)
	m := regexp.MustCompile("[:;()!,.0-9№-]")
	rgx := m.ReplaceAllString(result, "")
	res := strings.ToLower(rgx)
	return res
}

// TextPreprocessing ...
func TextPreprocessing(filename string) {

	// Открыть файл с данными

	file, err := os.Open("./upload/" + filename)
	if err != nil {
		fmt.Println("failed to open")
	}

	// Считывание текстовых данных из файла

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	// Запись нормализованных текстовых данных в файл

	fw, err := os.OpenFile("./sys/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	// Процедура нормализации

	for _, line := range text {
		normline := Normalize(line)
		_, err = fmt.Fprintln(fw, normline)
	}

	if err != nil {
		fmt.Println(err)
		fw.Close()
	}
	err = fw.Close()
	if err != nil {
		fmt.Println(err)
	}

}
