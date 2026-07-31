#include "friends.h"

#pragma once

namespace kf {
	class Raccoon : public Friends {
	private:
		std::string s = u8"よーしアライさんにお任せなのだー！";
	public:
		Raccoon() {}
		virtual ~Raccoon() {}
		virtual std::string sound() { return s; }
	};
}
