/*
Setiap bambu memiliki sekat. 
Buatlah aplikasi yang dapat menentukan jumlah bambu yang kamu miliki melalui inputan, 
setelah itu kamu dapat menentukan jumlah sekat dari setiap bambu yang kamu miliki. 

Jika sudah, potonglah bambu beberapa kali berdasarkan inputan yang kamu inginkan. 
Setiap potongan yang kamu lakukan akan menghilangkan 1 sekat disetiap bambu yang kamu miliki. 
Lalu tampilkanlah pada console sisa sekat bambu setiap pemotongan.

Kriteria:
  > Gunakan variabel untuk menerima inputan dari console
  > Gunakan function 

Sample Input:
Masukan banyak bamboo : 2
sekat bamboo ke-1 : 5
sekat bamboo ke-2 : 3
Berapa kali potong? 2

Sample Output :
Potongan ke- 1
4
2
Potongan ke- 2
3
1
*/

package main

import "fmt"

func main() {
	var jumlahBamboo int
	//fmt.Print("Masukan banyak bamboo : ")
	fmt.Scanln(&jumlahBamboo)
	var sekatBamboo = make([]int, jumlahBamboo)

	var sekat int
	for i := range sekatBamboo {
		//fmt.Printf("sekat bamboo ke-%d : ", i + 1)
		fmt.Scanln(&sekat)
		sekatBamboo[i] = sekat
	}

	var potongan int
	//fmt.Print("Berapa kali potong? ")
	fmt.Scanln(&potongan)
	potongBamboo(potongan, sekatBamboo)
}

func potongBamboo(potongan int, bamboo []int) {
	for i := 0; i < potongan; i++ {
		fmt.Println("Potongan ke-", i + 1)
		for j := 0; j < len(bamboo); j++ {
			if(bamboo[j] <= 0) {
				fmt.Println(bamboo[j])
				continue
			}
			bamboo[j]--
			fmt.Println(bamboo[j])
		}
	}
}