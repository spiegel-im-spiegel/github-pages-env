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

function build(theme, dist)
	if theme == "" then
		if dist == "" then
			return nyagos.rawexec("hugo.exe")
		else
			return nyagos.rawexec("hugo.exe", "--destination="..dist)
		end
	elseif dist == "" then
		return nyagos.rawexec("hugo.exe", "--theme="..theme)
	else
		return nyagos.rawexec("hugo.exe", "--theme="..theme, "--destination="..dist)
	end
end

function server(theme, enableDraft, port)
	if port == nil or port == "" then port = "1313" end
	if theme == "" then
		if enableDraft == true then
			return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--buildDrafts")
		else
			return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port)
		end
	elseif enableDraft == true then
		return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--buildDrafts", "--theme="..theme)
	else
		return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--theme="..theme)
	end
end
