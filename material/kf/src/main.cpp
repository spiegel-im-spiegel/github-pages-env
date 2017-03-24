#include <iostream>

#include "serval.h"
#include "fennec.h"
#include "raccoon.h"

void test(kf::Friends *f) {
	std::cout << f->sound() << std::endl;
}

int main() {
	test(new kf::Serval());
	test(new kf::Fennec());
	test(new kf::Raccoon());
	return 0;
}
