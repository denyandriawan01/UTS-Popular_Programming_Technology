package main

import "fmt"

func compareArrays(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2){
		return false
	}
	for i, x := range arr1 {
		if x != arr2[i] {
			return false
		}
	}
	return true
}

func main() {
	var size int

	fmt.Println("-----------------------------------------------------------------------------------         ")
  fmt.Println("   ___  _____  __  __  ____   __    ____  ____      __    ____  ____    __   _  _           ")
  fmt.Println("  / __)(  _  )(  \\/  )(  _ \\ /__\\  (  _ \\( ___)    /__\\  (  _ \\(  _ \\  /__\\ ( \\/ ) ")
  fmt.Println(" ( (__  )(_)(  )    (  )___//(__)\\  )   / )__)    /(__)\\  )   / )   / /(__)\\ \\  /       ")
  fmt.Println("  \\___)(_____)(_/\\/\\_)(__) (__)(__)(_)\\_)(____)  (__)(__)(_)\\_)(_)\\_)(__)(__)(__)     ")
	fmt.Println("                                                                                            ")
  fmt.Println("-----------------------------------------------------------------------------------         ")
	fmt.Print("[>] Masukkan ukuran array: ")
	fmt.Scanln(&size)

	arr1 := make([]string, size)
	arr2 := make([]string, size)

	fmt.Println(" | ")
	fmt.Print(" | [>] Array 1: ")
	for i := range arr1 {
		fmt.Scan(&arr1[i])
	}
	fmt.Print(" | [>] Array 2: ")
	for i := range arr2 {
		fmt.Scan(&arr2[i])
	}

	fmt.Println(" | ")
	fmt.Println("-----------------------------------------------------------------------------------         ")
	if compareArrays(arr1, arr2){
		fmt.Println("[!] Hasil: Kedua array memiliki value yang sama!")
	} else {
		fmt.Println("[!] Hasil:")
		for i, x := range arr1 {
			if x != arr2[i] {
				fmt.Printf(" | [>] Index ke-%d berbeda\n", i)
			}
		}
		fmt.Println("-----------------------------------------------------------------------------------         ")
	}
	
}