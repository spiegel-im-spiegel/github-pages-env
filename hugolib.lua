--[[
 Lua Library for Hugo

 These codes are licensed under CC0.
 http://creativecommons.org/publicdomain/zero/1.0/deed
]]
local hugolib = {}

--
-- Hugo File Management
--

-- Create New Post (hugo new ...)
hugolib.new = function(path)
	return nyagos.rawexec("hugo.exe", "new", path)
end

-- Release Post (hugo undraft ...)
hugolib.release = function(path)
	local hugopath = "content/"..path

	local lines = {}
	local reader, errmsg, errno = io.open(hugopath, "r")
	if reader == nil then return errno, errmsg end
	for line in reader:lines() do
		table.insert(lines, line)
	end

	local writer, errmsg, errno = io.open(hugopath, "w+b")
	if writer == nil then return errno, errmsg end
	local datestr = "date = \""..os.date("%Y-%m-%dT%H:%M:%S+09:00").."\"\n" -- JST only
	local tomlFlag = false
	for i, line in ipairs(lines) do
		if (not tomlFlag) and line == "+++" then
			writer:write(line.."\n")
			tomlFlag = true
		elseif  tomlFlag and line == "+++" then
			writer:write(line.."\n")
			tomlFlag = false
		elseif  tomlFlag and string.sub(line, 0, 4) == "date" then
			nyagos.write(datestr)
			writer:write(datestr) -- replace line ("date" element)
		elseif  tomlFlag and string.sub(line, 0, 5) == "draft" then
			-- delete line
		else
			writer:write(line.."\n")
		end
	end
	return 0, nil
end

-- Update Post (add "update" element in front matter)
--   caution: support TOML format only
hugolib.update = function(path)
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
	local tomlFlag = false
	for i, line in ipairs(lines) do
		if (not tomlFlag) and line == "+++" then
			writer:write(line.."\n")
			tomlFlag = true
		elseif  tomlFlag and line == "+++" then
			writer:write(line.."\n")
			tomlFlag = false
		elseif  tomlFlag and string.sub(line, 0, 4) == "date" then
			writer:write(line.."\n")
			nyagos.write(updatestr)
			writer:write(updatestr) -- add line ("update" element)
		elseif  tomlFlag and string.sub(line, 0, 6) == "update" then
			-- delete line
		else
			writer:write(line.."\n")
		end
	end
	return 0, nil
end

-- Build Site
hugolib.build = function(theme, dist)
	if hugolib.isBlank(theme) then
		if hugolib.isBlank(dist) then
			return nyagos.rawexec("hugo.exe")
		else
			return nyagos.rawexec("hugo.exe", "--destination="..dist)
		end
	elseif hugolib.isBlank(dist) then
		return nyagos.rawexec("hugo.exe", "--theme="..theme)
	else
		return nyagos.rawexec("hugo.exe", "--theme="..theme, "--destination="..dist)
	end
end

-- Server Mode (hugo server ...)
hugolib.server = function(theme, enableDraft, port)
	if hugolib.isBlank(port) then port = "1313" end
	if hugolib.isBlank(theme) then
		if enableDraft then
			return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--buildDrafts")
		else
			return nyagos.rawexec("hugo.exe", "server", "--watch", "--port=" .. port)
		end
	elseif enableDraft then
		return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--buildDrafts", "--theme="..theme)
	else
		return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--theme="..theme)
	end
end

--
-- Git Simple Operation
--

-- Git Add on stage (git add --all)
hugolib.git_add_all = function()
    nyagos.write("git add --all\n")
    return nyagos.rawexec("git", "add", "--all")
end

-- Git Commit (git commit -v -m  "comment")
hugolib.git_commit = function(comment)
    nyagos.write("git commit -v -m \""..comment.."\"\n")
    return nyagos.rawexec("git", "commit", "-v", "-m",  comment)
end

-- Git Push to remote repository (git push -u remote branch)
hugolib.git_push = function(remote, branch)
    if hugolib.isBlank(remote) then remote = "origin" end
    if hugolib.isBlank(branch) then branch = "master" end
    nyagos.write("git push -u "..remote.." "..branch.."\n")
    return nyagos.rawexec("git", "push", "-u", remote, branch)
end

-- Git Pull from remote repository (git pull --progress remote)
hugolib.git_pull = function(remote)
    if hugolib.isBlank(remote) then remote = "origin" end
    nyagos.write("git pull --progres "..remote.."\n")
    return nyagos.rawexec("git", "pull", "--progress", remote)
end

-- Initialize Git Submodule (git submodule update --init --recursive)
hugolib.git_submodule_init = function()
    nyagos.write("git submodule update --init --recursive\n")
    return nyagos.rawexec("git", "submodule", "update", "--init", "--recursive")
end

-- Update Git Submodule (git submodule update --init --recursive)
hugolib.git_submodule_update = function()
    nyagos.write("git submodule update --remote --merge\n")
    return nyagos.rawexec("git", "submodule", "update", "--remote", "--merge")
end

--
-- Utility Functions
--

-- Make Hugo Path (join strings. Delimiter is "/" character.)
hugolib.pathjoin = function(...)
	local elements = { ... }
	local ret = ""
	local sep = ""
	for i, elem in ipairs(elements) do
		ret = ret..sep..elem
		if sep == "" then sep = "/" end
	end
	return ret
end

hugolib.isBlank = function(str)
	return (str == nil or str == "")
end

--
-- return module
--

return hugolib
