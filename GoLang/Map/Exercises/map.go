package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    month := [12]string{"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}

    penjualan := map[string]int{
        "januari":   2836,
        "februari":  3282,
        "maret":     787,
        "april":     6238,
        "mei":       1992,
        "juni":      824,
        "juli":      2903,
        "agustus":   602,
        "september": 930,
        "oktober":   1002,
        "november":  922,
        "desember":  3219,
    }
	
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    bulan1, _ := strconv.Atoi(scanner.Text())
    scanner.Scan()
    bulan2, _ := strconv.Atoi(scanner.Text())

    //Tulis kode disini
	for i := bulan1; i <= bulan2; i++ {
		fmt.Printf("Bulan %s : %d tusuk\n", month[i - 1], penjualan[month[i - 1]])
	}
}