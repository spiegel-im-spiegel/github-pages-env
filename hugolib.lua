module("hugolib", package.seeall)

function pathjoin(...)
	local elements = { ... }
	local ret = ""
	local sep = ""
	for i, elem in ipairs(elements) do
		ret = ret..sep..elem
		if sep == "" then sep = "/" end
	end
	return ret
end

function build()
	return nyagos.rawexec("hugo.exe", "--theme=text", "--destination=../published")
end
