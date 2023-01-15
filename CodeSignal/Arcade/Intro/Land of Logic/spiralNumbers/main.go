package main

import "fmt"
import "strconv"

func spiralNumbers(n int) [][]int {
   c:=1
       mat := make([][]int,n)
       for i:=range mat{
           mat[i]=make([]int,n)
       }
       for i:=0; i< n/2; i++{
           for j:=i; j < n -i -1;j++{
               mat[i][j]=c
               c++
           }
           for j:=i; j < n -i -1;j++{
               mat[j][n-i-1]=c
               c++
           }
           for j:=i; j < n -i -1;j++{
               mat[n-i-1][n-j-1]=c
               c++
           }
           for j:=i; j < n -i -1;j++{
               mat[n-j-1][i]=c
               c++
           }
       }
       if n%2==1{
           mat[n/2][n/2]=n*n
       }
       return mat
}


func main() {
	fmt.Printf("%v", spiralNumbers(3))
}
