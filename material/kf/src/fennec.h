#include "friends.h"

#pragma once

namespace kf {
	class Fennec : public Friends {
	private:
		std::string s = u8"はーいよー頑張っていきますよー";
	public:
		Fennec() {}
		virtual ~Fennec() {}
		virtual std::string sound() { return s; }
	};
}
