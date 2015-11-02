require("hugolib")

if #arg < 1 then
	nyagos.writerr("Usage: update [<section>] <file name>\n")
	os.exit(2)
end

local path = hugolib.pathjoin("content", arg[1])
if #arg > 1 then
	local section = arg[1]
	if section == "remark" then
		if #arg > 2 then
			path = hugolib.pathjoin("content", arg[1], arg[2], arg[3])
		else
			path = hugolib.pathjoin("content", arg[1], os.date("%Y"), arg[2])
		end
	else
		path = hugolib.pathjoin("content", arg[1], arg[2])
	end
end
nyagos.write("Update: "..path.."\n")

local lines = {}
local reader = io.input(path)
for line in reader:lines() do
	table.insert(lines, line)
end

local writer = io.open(path, "wb")
local timestr = os.date("%Y-%m-%dT%H:%M:%S+09:00")
local flag = false
nyagos.write("Add: update = \""..timestr.."\"\n")
for i, line in ipairs(lines) do
	if (not flag) and line == "+++" then
		writer:write(line.."\n") -- そのまま出力
		flag = true
	elseif  flag and line == "+++" then
		writer:write(line.."\n") -- そのまま出力
		flag = false
	elseif  flag and string.sub(line, 0, 4) == "date" then
		writer:write(line.."\n") -- そのまま出力
		writer:write("update = \""..timestr.."\"\n") -- update を出力
	elseif  flag and string.sub(line, 0, 6) == "update" then
		-- 何もしない（行削除）
	else
		writer:write(line.."\n") -- そのまま出力
	end
end
