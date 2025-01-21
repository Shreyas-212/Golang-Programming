package main 

import (
	"fmt"
)

func ArmstrongNum(num int) {
	temp := num
    sum := 0
	rem := 0
    for temp > 0 {
		rem = temp % 10
		sum += rem * rem * rem
		temp /= 10
	}
    if num == sum {
        fmt.Printf("%d is an Armstrong number\n", num)
    } else {
        fmt.Printf("%d is not an Armstrong number\n", num)
    }
}

func main() {
	ArmstrongNum(12345)
    ArmstrongNum(1634)
    ArmstrongNum(370)
    ArmstrongNum(12345)
}