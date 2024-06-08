package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var path = "/2_Study/Enigma_Code_Camp/GoLang/File/data/"
	var fileName = "example.txt"
	var filePath = path + fileName

	//Tahap buat File (Create)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Println("File", fileName, "Has been created")
	}

	//Tahap buka file (Open) untuk mulis/baca sesuatu di file harus buka file dulu
	file, err := os.OpenFile(filePath, os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//Tahap tulis ke file (Write)
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Write your name: ")
	scanner.Scan()
	input = scanner.Text()

	writer := bufio.NewWriter(file) //"file" is the os.OpenFile()
	writer.WriteString(input + "\n") //collect the "input" where we write to the "file"
	writer.Flush() //add all the string arguement to the file

	//Tahap untuk baca file (Read)
	scan_file := bufio.NewScanner(file)
	for  scan_file.Scan() {
		fmt.Println(scan_file.Text())
	}
}
