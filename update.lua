require("hugolib")

if #arg < 1 then
	nyagos.writerr("Usage: update [<section>] <file name>\n")
	os.exit(2)
end

local path = arg[1]
if #arg > 1 then
	local section = arg[1]
	if section == "remark" then
		if #arg > 2 then
			path = hugolib.pathjoin(section, arg[2], arg[3])
		else
			path = hugolib.pathjoin(section, os.date("%Y"), arg[2])
		end
	else
		path = hugolib.pathjoin(section, arg[2])
	end
end
nyagos.write("Update: "..path.."\n")
local errorlevel, errormessage = hugolib.update(path)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
