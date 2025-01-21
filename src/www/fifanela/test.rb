def fibonacci(n)
  if n == 0 or n == 1
    return 1
  else
    return fibonacci(n-1) + fibonacci(n-2)
  end
end

def print_fib_serie(n)
  output = fibonacci(0).to_s
  for x in 1..n
    output += ", #{fibonacci(x)}"
  end
  return output
end

puts print_fib_serie(10)