require("hugolib")

if #arg < 1 then
	nyagos.writerr("Usage: rel [<section>] <file name>\n")
	os.exit(2)
end

local path = hugolib.pathjoin("content", arg[1])
if #arg > 1 then
	section = arg[1]
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
nyagos.write("Release: "..path.."\n")
local errorlevel, errormessage = nyagos.rawexec("hugo.exe", "undraft", path)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
