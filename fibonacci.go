package main

import "fmt"

//
// Simple recurence relation which expresses the Fibonacci sequence.
//
func fibonacci_recursive( n int ) int {

     if n < 2 {

        return n

     }  else  {

        return fibonacci_recursive(n - 1) + fibonacci_recursive(n - 2)
     }
}

//
// This is the iterative implementation of the Fibonacci Sequence.
//

func fibonacci_iterative( n int ) int {

     a, b := 0, 1

     fmt.Printf("%d\n%d\n", a, b)

     for ; n>=0; n-- {

           a, b = b, a + b

           fmt.Printf("%d\n", b)
     }      

     return a
}

//
//  This Fibonacci implementation uses a closure to generate succesive Fibonacci terms.
//
func FibClosure() func() int {

	a, b := 0, 1

	return func() int {

		a, b = b, a + b

		return a
	}
}

func fibonacci(n int) {

     a, b := 0, 1
     count := n - 2 

     fmt.Printf("%d\n%d\n", a, b)

     for ;count >= 0; count-- {

          if b > a {

             a += b

             fmt.Printf("%d\n", a)

          } else {

             b += a 

             fmt.Printf("%d\n", b)
          }
     }
}


func main() {

     fmt.Printf("%d", fibonacci_iterative(5))
     fmt.Printf("%d", fibonacci_recursive(5))

     nextFib := FibClosure()
     fmt.Println(nextFib())
     fmt.Println(nextFib())
     fmt.Println(nextFib())
     fmt.Println(nextFib())
     fmt.Println(nextFib())

     fibonacci(5)

}

