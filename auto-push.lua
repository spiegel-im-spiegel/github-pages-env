require("hugolib")

nyagos.write("build...\n")
local errorlevel, errormessage = hugolib.build()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = nyagos.exec("pushd ..\\published")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\ngit add...\n")
errorlevel, errormessage = nyagos.rawexec("git", "add", "*")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

local timestr = os.date("%Y-%m-%dT%H:%M:%S+09:00")
nyagos.write("\ngit commit...\n")
errorlevel, errormessage = nyagos.rawexec("git", "commit", "-v", "-m",  "Publish (auto commit in "..timestr..")")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\ngit push...\n")
errorlevel, errormessage = nyagos.rawexec("git", "push", "-u", "origin", "master")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = nyagos.exec("popd")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
