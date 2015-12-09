require("hugolib")

if #arg < 1 then
	nyagos.writerr("Usage: rel [<section>] <file name>\n")
	os.exit(2)
end

local path = arg[1]
if #arg > 1 then
	section = arg[1]
	if section == "remark" then
		local year = os.date("%Y")
		if year == "2015" then
			if #arg > 2 then
				path = hugolib.pathjoin(section, arg[2], arg[3])
			else
				path = hugolib.pathjoin(section, year, arg[2])
			end
		else
			if #arg > 3 then
				path = hugolib.pathjoin(section, arg[2], arg[3], arg[4])
			elseif #arg > 2 then
				path = hugolib.pathjoin(section, year, arg[2], arg[3])
			else
				path = hugolib.pathjoin(section, year, os.date("%m"), arg[2])
			end
		end
	else
		path = hugolib.pathjoin(section, arg[2])
	end
end
nyagos.write("Release: "..path.."\n")
local errorlevel, errormessage = hugolib.release(path)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
