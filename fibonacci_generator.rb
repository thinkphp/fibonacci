fib = Enumerator.new do |y|
    a, b = 0, 1
    loop do
         y.yield a
         a, b = b, a + b
    end 
end

p fib.take(200)

    
