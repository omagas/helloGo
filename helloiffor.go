package main

import "fmt"

func main() {
  	fmt.Println("************Result*****************");
  	fmt.Println("【IF】")
	var x int
	x=20
	if x > 10 {
	    fmt.Println("x is greater than 10")
	} else {
	    fmt.Println("x is less than 10")
	}    
 


  	fmt.Println("【For】")
    sum := 0;
    star:="*";
    result:="";
    for index:=0; index < 10 ; index++ {
        sum += index
		result=result+star
        fmt.Println(result)
    }
    fmt.Println("sum is equal to ", sum)
}