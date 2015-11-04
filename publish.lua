require("hugolib")

nyagos.write("build...\n")
local errorlevel, errormessage = hugolib.build("text", "../published")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\nmove publish folder\n")
errorlevel, errormessage = nyagos.exec("pushd ..\\published")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

nyagos.write("\npublish...\n")
errorlevel, errormessage = hugolib.git_add_all()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = hugolib.git_commit("Publish (auto commit in "..os.date("%Y-%m-%dT%H:%M:%S+09:00")..")")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = hugolib.git_push("origin", "master")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = nyagos.exec("popd")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
