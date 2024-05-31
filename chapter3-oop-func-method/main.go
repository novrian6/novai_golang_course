package main

import "fmt"

// Definisikan sebuah struktur 'Person'
type Person struct {
	Name string
	Age  int
}

// Metode untuk struktur 'Person' yang akan mencetak informasi tentang orang tersebut
func (p Person) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// Fungsi untuk membuat dan mengembalikan instans 'Person' baru
func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func main() {
	// Membuat instans 'Person' menggunakan fungsi konstruktor 'NewPerson'
	person1 := NewPerson("John", 30)
	// Memanggil metode 'PrintInfo' pada instans 'Person'
	person1.PrintInfo()

	// Menggunakan penugasan langsung untuk membuat instans 'Person'
	person2 := Person{Name: "Jane", Age: 25}
	// Memanggil metode 'PrintInfo' pada instans 'Person'
	person2.PrintInfo()
}

