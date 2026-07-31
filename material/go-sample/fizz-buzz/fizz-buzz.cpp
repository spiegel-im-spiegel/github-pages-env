#include <iostream>
#include <string>

auto newCounter() {
	return [ct=0]() mutable {
		return ++ct;
    };
}

auto fizzBuzz(auto n) {
	if (n%3==0 && n%5==0) {
		return std::string("Fizz Buzz");
	} else if (n%3==0) {
		return std::string("Fizz");
	} else if (n%5==0) {
		return std::string("Buzz");
	}
	return std::to_string(n);
}

int main() {
	auto max = 30;
	auto c = newCounter();
	for (auto n = c(); ; n = c()) {
		if (n < max) {
			std::cout << fizzBuzz(n) << ", ";
		} else {
			std::cout << fizzBuzz(n) << std::endl;
			break;
		}
	}
	return 0;
}
