require("hugolib")

if #arg < 1 then
	nyagos.writerr("Usage: new [<section>] <file name>\n")
	os.exit(2)
end

local path = arg[1]
local year = os.date("%Y")
local monthday = os.date("%m%d")
if #arg == 1 then
	if arg[1] == "remark" then
		path = hugolib.pathjoin(arg[1], year, monthday.."-diary.md")
	end
else
	section = arg[1]
	file = arg[2]
	if section == "remark" then
		path = hugolib.pathjoin(section, year, file)
	else
		path = hugolib.pathjoin(section, file)
	end
end
nyagos.write("Create: "..path.."\n")
local errorlevel, errormessage = nyagos.rawexec("hugo.exe", "new", path)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
