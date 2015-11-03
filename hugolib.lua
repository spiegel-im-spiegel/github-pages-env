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

function new(path)
	return nyagos.rawexec("hugo.exe", "new", path)
end

function release(path)
	return nyagos.rawexec("hugo.exe", "undraft", "content/"..path)
end

function update(path)
	local hugopath = "content/"..path

	local lines = {}
	local reader, errmsg, errno = io.open(hugopath, "r")
	if reader == nil then return errno, errmsg end
	for line in reader:lines() do
		table.insert(lines, line)
	end

	local writer, errmsg, errno = io.open(hugopath, "w+b")
	if writer == nil then return errno, errmsg end
	local updatestr = "update = \""..os.date("%Y-%m-%dT%H:%M:%S+09:00").."\"\n" -- JST only
	local flag = false
	for i, line in ipairs(lines) do
		if (not flag) and line == "+++" then
			writer:write(line.."\n")
			flag = true
		elseif  flag and line == "+++" then
			writer:write(line.."\n")
			flag = false
		elseif  flag and string.sub(line, 0, 4) == "date" then
			writer:write(line.."\n")
			nyagos.write("Add: "..updatestr)
			writer:write(updatestr) -- add line ("update" element)
		elseif  flag and string.sub(line, 0, 6) == "update" then
			-- delete line
		else
			writer:write(line.."\n")
		end
	end
	return 0, nil
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
