package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
	// Write your code here
	encoded := ""

	k = k % ((122 - 97) + 1)
	letters := strings.Split(s, "")
	for _, v := range letters {
		if v[0] >= 65 && v[0] <= 90 {
			letter := (int32(v[0]) + k) % (90 + 1)
			if letter < 65 {
				letter += 65
			}
			encoded = fmt.Sprintf("%s%c", encoded, letter)
		} else if v[0] >= 97 && v[0] <= 122 {

			letter := (int32(v[0]) + k) % (122 + 1)

			if letter < 97 {
				letter += 97
			}

			fmt.Printf("%c %d letter: %c %d \n", v[0], v[0], letter, letter)

			encoded = fmt.Sprintf("%s%c", encoded, letter)
		} else {
			encoded = fmt.Sprintf("%s%c", encoded, v[0])
		}
	}

	fmt.Printf("%s\n", encoded)

	return encoded
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
