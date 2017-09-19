##
## Reference -> https://stackoverflow.com/questions/3953749/python-fibonacci-generator
## 
def fibonacci_rec(n):
    
    if(n < 2):

       return n

    else:
         return fibonacci_rec(n - 1) + fibonacci_rec(n - 2)


def fibonacci_ite(n):
    
      a, b = 0, 1
      n -= 2 

      print a, b      

      while(n != 0):

          a, b = b, a + b
          n -= 1
          print b
   
      return b 
    
#
#  Python is dinamically typed language, so the type of the variable is determined at the run time and it
#  can vary as the execution is in progress.
#

def fib():

    a, b = 0, 1
    while 1:
          yield a
          a, b = b, a + b


x = fib()
n = 5
for i in range(n):
    print x.next()         
