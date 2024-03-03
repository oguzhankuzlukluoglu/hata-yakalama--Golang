package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("lutfen aldiginiz notu giriniz  ")
	grade, _ := getGrade()
	var result string
	if grade >= 45 {
		result = "gectiniz"
	} else {
		result = "kaldiniz"
	}
	fmt.Println(result)

}
func getGrade() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
	}
	return num, nil

}
