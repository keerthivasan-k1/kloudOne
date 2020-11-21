package main

import "fmt"

func main() {

    a:= [5]int{3,4,5,6,7}
    fmt.Println("emp:", a)

    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    

// no. of values in matrices is x
// no.of matrices are represented by y
   const x,y =4,5
   var  twoD  [x][y]int
    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}