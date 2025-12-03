package main

import "fmt"

func main() {
    var menu int 
    var loopMenu string
    loopMenu = "y"
    
    for loopMenu != "n"  {
        fmt.Println("Pilihan Menu Kalkulator")
        fmt.Println("1. Pertambahan")
        fmt.Println("2. Pengurangan")
        fmt.Println("3. Perkalian")
        fmt.Println("4. Pembagian")
        fmt.Println("")
        fmt.Print("Masukkan Pilihan Menu: ")
        
        fmt.Scan(&menu)
        switch menu {
            case 1: inputAngka(add)
            case 2: inputAngka(sub)
            case 3: inputAngka(multi)
            case 4: inputAngka(div)
            default: fmt.Println("Menu Pilihan Tidak ada")
        }
        
        fmt.Print("Apakah Ingin Lanjut atau Tidak (Ketik (y) jika lanjut (n) jika tidak): ")
        fmt.Scan(&loopMenu)
        
    }
    
    
}

func inputAngka(f func (int, int) ) {
    var a, b int
    
    fmt.Print("Masukkan Angka 1: ")
    fmt.Scan(&a)
    
    fmt.Print("Masukkan Angka 2: ")
    fmt.Scan(&b)
    
    f(a,b)
    
}

func add(a, b int){
    fmt.Println("Hasil Pertambahan ",a, " + ", b,"= ", a+b)
}

func sub(a, b int){
    fmt.Println("Hasil Pengurangan ",a, " - ", b,"= ", a-b)
}

func multi(a, b int){
    fmt.Println("Hasil Perkalian ",a, " x ", b,"= ", a*b)
}

func div(a, b int){
    fmt.Println("Hasil Pembagian ",a, " : ", b,"= ", a/b)
}