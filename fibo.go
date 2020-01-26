package main

import "fmt"

func fibo(n int) int {

	 var a, b int = 1, 1
     var i int = 3
     var s int
     fmt.Println(a, b)       
	 for i <= n {

         s = a + b
         a, b = b, s
         i = i + 1
         fmt.Println(s)
	 }

	 return s
}

func fibo_rec(n int) int {

	 if n == 1 || n == 0 { 

	 	return 1 

	 } else { 

	      return fibo_rec(n-1) + fibo_rec(n-2)
	 }    
}

func main() {

	 fmt.Println(fibo(10))
}
