package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	tokens := strings.Split(s, ":")

	tmpHour, err := strconv.Atoi(tokens[0])
	if err != nil {
		log.Fatal("Cannot convert")
		return s
	}

	if strings.Contains(tokens[2], "PM") {

		tokens[0] = fmt.Sprintf("%02d", (tmpHour%12)+12)
		tokens[2] = strings.ReplaceAll(tokens[2], "PM", "")
	} else {

		tokens[0] = fmt.Sprintf("%02d", tmpHour%12)
		tokens[2] = strings.ReplaceAll(tokens[2], "AM", "")
	}

	return strings.Join(tokens, ":")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
