package main

import (
	"fmt"
	"strings"

	"github.com/redpandaathome/learngo/mydict"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

// 💜 naked return
func lenAndUpper2(name string) (length int, uppercase string) {
	// 💜 defer
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

// 💜 for...
func superAdd(numbers ...int) int {
	// for i:=0; i<len(numbers); i++ {
	// 	fmt.Println((numbers[i]))
	// }
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// 💜 if && variable expression
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true

	// 💜 switch
	// switch {
	// case age >= 21:
	// 	fmt.Println("🍺🍺🍺")
	// 	return true
	// case age < 21:
	// 	fmt.Println("getSome milk instead")
	// 	return false
	// default:
	// 	fmt.Println("??")
	// 	return true
	// }

	// switch koreanAge := age+2; koreanAge {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
}

// struct...
// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

func main() {
	// fmt.Println("Hello world!")
	// fmt.Println(multiply(2, 3))
	// // totalLength, _ := lenAndUpper("yumi")
	// totalLength, upperName := lenAndUpper("yumi")
	// fmt.Println(totalLength, upperName)

	// totalLength2, upperName2 := lenAndUpper2("yodi")
	// fmt.Println(totalLength2, upperName2)

	// repeatMe("yumi", "bread", "icecream", "pasta", "drink")

	// total := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println(total)

	// fmt.Println(canIDrink(18))
	// fmt.Println(canIDrink(21))

	// 💜
	// a := 2
	// b := a
	// a = 10
	// fmt.Println(a, b)
	// fmt.Println(&a, &b) // memory address

	// c := 2
	// d := &c
	// c = 10
	// fmt.Println(c, d)
	// fmt.Println(&c, d, *d) //c의 주소, c의주소, c의 내용물(value)

	// e := 100
	// f := &e
	// *f = 555
	// fmt.Println(e, *f)

	// 💜 Arrays and Slices
	// names := [5]string{"banana", "kiwi", "apple"}
	// names[3] = "pear"
	// names[4] = "mango"
	// fmt.Println(names)

	// //slice - unlimited
	// names2 := []string{"banana"}
	// names3 := append(names2, "hi") //returns a new slice! 👀
	// fmt.Println((names3))

	// 💜 map -key&value
	// yumi := map[string]string{"name": "yumi", "age": "12"} //👀
	// fmt.Println(yumi)
	// for key, value := range yumi {
	// 	fmt.Println(key, value)
	// }

	// 💜 struct
	// favFood := []string{"ramen", "pasta", "grilled veges"}
	// // yumi := person{"yumi", 12, favFood}
	// yumi := person{name: "yumi", age: 12, favFood: favFood}

	// fmt.Println(yumi)

	// 💰 Bank !
	// account := banking.Account{Owner: "yumi", Balance: 100}
	// fmt.Println(account)

	// account := accounts.NewAccount("yumi")
	// fmt.Println(account)
	// account.Deposit(10)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(11)
	// if err != nil {
	// 	// log.Fatalln(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// account.ChangeOwner("yodi")
	// fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account.String())

	//Dict
	dictionary := mydict.Dictionary{"first": "First word"}
	// dictionary["hello"] = "hello"
	definition1, err1 := dictionary.Search("first")
	definition2, err2 := dictionary.Search("second")
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(definition1, definition2)

	word := "food"
	definition := "love"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	love, _ := dictionary.Search(word)
	fmt.Println("found:", word, "def:", love)
	err3 := dictionary.Add(word, definition)
	if err3 != nil {
		fmt.Println(err3)
	}
}
