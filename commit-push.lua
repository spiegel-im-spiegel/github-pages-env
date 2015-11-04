require("hugolib")

nyagos.write("\ngit add...\n")
local errorlevel, errormessage = hugolib.git_add_all()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\ngit commit...\n")
errorlevel, errormessage = hugolib.git_commit("auto commit test in "..os.date("%Y-%m-%dT%H:%M:%S+09:00"))
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\ngit push...\n")
errorlevel, errormessage = hugolib.git_push("origin", "master")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
