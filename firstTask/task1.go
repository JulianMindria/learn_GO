package firstTask

import(
	"fmt"
	"math"
)

func RoundNum(num float64){
	dividedBy := math.Pow(10, float64(1))
	fmt.Printf("%.2f\n", math.Round(num*dividedBy)/dividedBy)
}

