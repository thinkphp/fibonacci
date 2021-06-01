#fibonacci with lists
def fib_list(n):

    list = [0] * (n + 1)

    list[0] = 0

    list[1] = 1

    for i in range(2, n + 1):

        list[i] = list[i-1] + list[i-2]

    return list[n]
 
     
#fibonacci with tail recursion
def fib_tail(n, a = 0, b = 1):

    if n == 0:

       return a

    else:

       return fib_tail(n-1, b, a + b) 

#fibonacci with memoization
def memoize(f):

    memo = {}

    def helper(x):

        if x not in memo:

           memo[x] = f(x)

        return memo[x]

    return helper

@memoize
def fib(n):

    if n == 0:
       return 0
    elif n == 1:
       return 1
    else:
       return fib( n - 1 ) + fib( n - 2 ) 

#fibonacci naive approach with for range
def fibo(n):
    a = 0
    b = 1
    for i in range(2, n + 1):
        a, b = b, a + b
    return b

#fibo with while
def fibo2(n):
    a = 0
    b = 1
    i = 2
    while i <= n:
       c = a + b       
       a = b
       b = c 
       i += 1
    return b 

# 0,1,1,2,3,5,8,...
def main():
    
   print(fib_tail(290))
   print(fib(290))
   print(fib_list(290))
   print(fibo(290))
   print(fibo2(290))
main()
