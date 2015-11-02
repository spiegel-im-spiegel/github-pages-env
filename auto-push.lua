require("hugo")

local errorlevel, errormessage = hugo.build()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = nyagos.rawexec("git", "add", "*")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

local timestr = os.date("%Y-%m-%dT%H:%M:%S+09:00")
errorlevel, errormessage = nyagos.rawexec("git", "commit", "-v", "-m  \"Publish (auto commit in "..timestr..")\"")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = nyagos.rawexec("git", "push", "-u origin master")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
