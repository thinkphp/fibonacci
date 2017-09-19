#
# This is a tail-call recursion of the algorithm
#
def fibonacci_tail(n, a = 0, b = 1)

    if n == 0 
       a
    else
       fibonacci_tail(n - 1, b, a + b)
    end

end

#
# The only limit you have is the maximum INT value.
#
def fibonacci_iterative(n)

    a, b = 0, 1       
    p a
    p b
    n -= 2
    while n != 0                      
          a, b = b, a + b
          p b
          n -= 1
    end  
    b
end

#
# Recursive version of the algorithm.
# fib(0) = 0
# fib(1) = 1
# fib(n) = fib(n-1) + fib(n-2) 
#
def fibonacci_rec(n)

    if n < 2 
       return n 
    else
       return fibonacci_rec( n - 1) + fibonacci_rec(n - 2)
    end

end

def fibonacci2(count)

    a, b = 0, 1 
    p a
    p b
    count -= 2 
    while count != 0
          p (b > a) ? (a += b) : (b += a)
          count -= 1 
    end 
    b
end

def fibonacci_loop( count )

    a, b = 0, 1
    p a
    p b
    (count-1).times do |i|        
          a, b = b, a + b
          p b
    end     
    p a
end

#p fibonacci_iterative(10)
#p fibonacci_tail(5)
#p fibonacci_rec(5)
#p fibonacci2(10)
fibonacci_loop(4)