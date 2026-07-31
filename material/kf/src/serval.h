#include "friends.h"

#pragma once

namespace kf {
	class Serval : public Friends {
	private:
		std::string s = u8"私はサーバルキャットのサーバルだよ！";
	public:
		Serval() {}
		virtual ~Serval() {}
		virtual std::string sound() { return s; }
	};
}
