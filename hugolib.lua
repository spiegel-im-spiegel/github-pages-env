--[[
 Lua Library for Hugo

 These codes are licensed under CC0.
 http://creativecommons.org/publicdomain/zero/1.0/deed
]]
module("hugolib", package.seeall)

--
-- Hugo File Management
--

-- Create New Post (hugo new ...)
function new(path)
	return nyagos.rawexec("hugo.exe", "new", path)
end

-- Release Post (hugo undraft ...)
function release(path)
	return nyagos.rawexec("hugo.exe", "undraft", "content/"..path)
end

-- Update Post (add "update" element in front matter)
--   caution: support TOML format only
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
function build(theme, dist)
	if isBlank(theme) then
		if isBlank(dist) then
			return nyagos.rawexec("hugo.exe")
		else
			return nyagos.rawexec("hugo.exe", "--destination="..dist)
		end
	elseif isBlank(dist) then
		return nyagos.rawexec("hugo.exe", "--theme="..theme)
	else
		return nyagos.rawexec("hugo.exe", "--theme="..theme, "--destination="..dist)
	end
end

-- Server Mode (hugo server ...)
function server(theme, enableDraft, port)
	if isBlank(port) then port = "1313" end
	if isBlank(theme) then
		if enableDraft then
			return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port, "--buildDrafts")
		else
			return nyagos.rawexec("hugo.exe", "server", "--watch", "--port="..port)
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
function git_add_all()
    nyagos.write("git add --all\n")
    return nyagos.rawexec("git", "add", "--all")
end

-- Git Commit (git commit -v -m  "comment")
function git_commit(comment)
    nyagos.write("git commit -v -m \""..comment.."\"\n")
    return nyagos.rawexec("git", "commit", "-v", "-m",  comment)
end

-- Git Push to remote repository (git add -u remote branch)
function git_push(remote, branch)
    if isBlank(remote) then remote = "origin" end
    if isBlank(branch) then branch = "master" end
    nyagos.write("git push -u "..remote.." "..branch.."\n")
    return nyagos.rawexec("git", "push", "-u", remote, branch)
end

-- Git Pull from remote repository (git pull --progress remote)
function git_pull(remote)
    if isBlank(remote) then remote = "origin" end
    nyagos.write("git pull --progres "..remote.."\n")
    return nyagos.rawexec("git", "pull", "--progress", remote)
end

-- Update Git Submodule (git submodule update --init --recursive)
function git_submodule_update()
    -- nyagos.write("git submodule update --init --recursive\n")
    -- return nyagos.rawexec("git", "submodule", "update", "--init", "--recursive")
    nyagos.write("git submodule update --remote --merge\n")
    return nyagos.rawexec("git", "submodule", "update", "--remote", "--merge")
end

--
-- Utility Functions
--

-- Make Hugo Path (join strings. Delimiter is "/" character.)
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

function isBlank(str)
	return (str == nil or str == "")
end
