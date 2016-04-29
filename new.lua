local Hugolib = require("hugolib")

if #arg < 1 then
	nyagos.writerr("Usage: new [<section>] <file name>\n")
	os.exit(2)
end

local path = arg[1]
local year = os.date("%Y")
local month = os.date("%m")
local day = os.date("%d")
if #arg == 1 then
	if arg[1] == "remark" then
		local section = arg[1]
		if year == "2015" then
			path = Hugolib.pathjoin(section, year, month..day.."-stories.md")
		else
			path = Hugolib.pathjoin(section, year, month, day.."-stories.md")
		end
	elseif arg[1] == "bookmarks" then
		path = Hugolib.pathjoin(arg[1], year, month, day.."-bookmarks.md")
	end
else
	local section = arg[1]
	file = arg[2]
	if section == "remark" or section == "bookmarks" then
		if year == "2015" then
			path = Hugolib.pathjoin(section, year, file)
		else
			path = Hugolib.pathjoin(section, year, month, file)
		end
	else
		path = Hugolib.pathjoin(section, file)
	end
end
nyagos.write("Create: "..path.."\n")
local errorlevel, errormessage = Hugolib.new(path)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
