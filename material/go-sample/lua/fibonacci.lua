function fibonacciNumber(n)
	if n < 2 then
		return n
	end
	return fibonacciNumber(n - 2) + fibonacciNumber(n - 1)
end

print(fibonacciNumber(38))
