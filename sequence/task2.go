package sequence

import(
	"fmt"
)

type DeretBilang struct {
    Limit int
}

func (d DeretBilang) PrimeSeq(){
	limit := d.Limit*4
	if d.Limit < 2 {
		fmt.Println("Enter number from 2 to n")
	}
	arr := []int{}
    for i:=2; i<=limit; i++{
        isPrime:=true
        for j:=2; j<i; j++{
            if i % j == 0 {
                isPrime = false
            }
        }
        if isPrime == true {
            arr = append(arr, i)
        } 
    }
	var prime []int = arr[0:d.Limit]
	fmt.Println(prime)
}

func (d DeretBilang) FibonacciSeq(){
	limit := d.Limit-3
	fib := []int{0,1}
	for i := 0; i<=limit; i++{
		new_fib := fib[i]+fib[i+1]
		fib = append(fib, new_fib)
	}
	fmt.Println(fib)
}

func (d DeretBilang) EvenNum(){
	limit := d.Limit*2
	arr := []int{}
	for i:=0; i<=limit; i++{
		if i%2==0{
			arr = append(arr, i)
		}
	}
	var even []int = arr[0:d.Limit]
	fmt.Println(even)
}

func (d DeretBilang) OddNum(){
	limit := d.Limit*2
	arr := []int{}
	for i:=0; i<=limit; i++{
		if i%2!=0{
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)
}