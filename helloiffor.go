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
 


    fmt.Println("【For sum】")
    sum := 0;
    star:="*";
    result:="";
    for index:=0; index < 10 ; index++ {
        sum += index
    result=result+star
        fmt.Println(result)
    }
    fmt.Println("sum is equal to ", sum)


    sum = 1
    for sum < 1000 {
        sum += sum
    } 
    fmt.Println("sum is equal to ", sum)   


    fmt.Println("【SumAndProduct】")
    x1 := 3
    y1 := 4

    xPLUSy, xTIMESy := SumAndProduct(x1, y1)

    fmt.Printf("%d + %d = %d\n", x1, y1, xPLUSy)
    fmt.Printf("%d * %d = %d\n", x1, y1, xTIMESy)    




}





//返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
    return A+B, A*B
}