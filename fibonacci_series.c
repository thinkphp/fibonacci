#include <stdio.h>

void fibo(int count) {

     int a = 0, b = 1;

     count -= 2;

     printf("%d\n%d\n", a, b);

     while(count--) {

           printf("%d\n", (b > a)? (a += b) : (b += a));
     } 
};

int fibo_rec(int c) {

     if(c < 2) return c;

     else return fibo_rec(c - 1) + fibo_rec(c - 2);
}; 

int fibo_iter(int count) {

     int a = 0, b = 1, tmp;

     printf("%d\n%d\n", 0, 1);

     while(count!=0) {
           tmp = a + b;
           a = b;
           b = tmp;
           printf("%d\n", b);
           count--;  
     }
     return b;
};

int fibo_iter2(int count) {

     int i, a = 0, b = 1, tmp;
     printf("%d\n%d\n", 0, 1);
     
     for(i = 0; i < count; ++i) {

           tmp = a + b;
           a = b;
           b = tmp;

           printf("%d\n", b);
     }
     return b;
};


int main(int argc, int **argv) {
    
    fibo(argv[1] ? atoi(argv[1]) : 7);
    printf("%d", fibo_iter2(10));
    return(0);
};