local Hugolib = require("hugolib")

local pubdir = "../published"
local theme = "text"

nyagos.write("build for "..pubdir.." folder.\n")
local errorlevel, errormessage = Hugolib.build(theme, pubdir)
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\nmove publish folder\n")
errorlevel, errormessage = nyagos.exec("pushd "..pubdir)
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\npublish...\n")
errorlevel, errormessage = Hugolib.git_add_all()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = Hugolib.git_commit("Publish (auto commit in "..os.date("%Y-%m-%dT%H:%M:%S+09:00")..")")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = Hugolib.git_push("origin", "master")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\nreturn folder\n")
errorlevel, errormessage = nyagos.exec("popd")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
