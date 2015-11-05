require("hugolib")

errorlevel, errormessage = hugolib.git_pull("origin")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = hugolib.git_submodule_update()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
