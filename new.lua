require("hugolib")

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
		if year == "2015" then
			path = hugolib.pathjoin(arg[1], year, month..day.."-stories.md")
		else
			path = hugolib.pathjoin(arg[1], year, month, day.."-stories.md")
		end
	end
else
	section = arg[1]
	file = arg[2]
	if section == "remark" then
		if year == "2015" then
			path = hugolib.pathjoin(section, year, file)
		else
			path = hugolib.pathjoin(section, year, month, file)
		end
	else
		path = hugolib.pathjoin(section, file)
	end
end
nyagos.write("Create: "..path.."\n")
local errorlevel, errormessage = hugolib.new(path)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
