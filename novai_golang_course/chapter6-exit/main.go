package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Before os.Exit")

    // Keluar dari program dengan kode keluar 1
    os.Exit(1)

    // Perintah ini tidak akan pernah dijalankan karena os.Exit sudah dipanggil sebelumnya
    fmt.Println("After os.Exit")
}

