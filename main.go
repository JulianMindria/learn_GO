package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {
	// Task 1
	// segitiga(9)

	// Task 2
	// genPass("pass", "easy", 32)

	// Task 3
	// elems := []int{1, 7, 3, 4, 8, 9}
	// chooseMovie(elems, 8)
}

// TASK 1 - INVERTED TRIANGLE

func segitiga(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}




// TASK 2 - GENERATE PASSWORD

var (
	combination = ""
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialBytes = "!@#$%^&*()_+-=[]{}\\|;':\",.<>/?`~"
	numBytes = "0123456789"
)


func genPass(pass string, level string, password_length int) {
	rand.Seed(time.Now().UnixNano())
	pass_length := password_length - len(pass)
	min := 0
	max := 1
	password := ""

	for _, splitPass := range strings.Split(pass, ""){
		if rand.Intn(max - min + 1) + min == 1 {
			password += strings.ToUpper(splitPass)
		}else{
			password += splitPass
		}
	}
	
	if level == "strong" {
		combination += letterBytes+specialBytes+numBytes
	} else if level == "medium"{
		combination += letterBytes+numBytes
	} else if level == "easy"{
		combination += letterBytes
	}
	
	for i := 0; i <=pass_length; i++{
		c := combination[rand.Intn(len(combination))]
		password += string(c)
	}
	fmt.Println(password)
}


// TASK 3 - MOVIE FLIGHT

func sum(arr []int) int {
    sum := 0
    for _, valueInt := range arr {
        sum += valueInt
    }
    return sum
}

func chooseMovie(elems []int, flight_time int) {
	n := len(elems)
	for num:=0;num < (1 << uint(n));num++ {
		combination := []int{}
		for i:=0;i<n;i++ {
			if num & (1 << uint(i)) != 0 {
				combination = append(combination, elems[i])
			}
		}
		if sum(combination) == flight_time{
			if len(combination) != 0 {
				fmt.Println(combination)
			}else if len(combination) == 0 {
				fmt.Println("there are no film you could possibly watch")
			}
		}
	}
}



