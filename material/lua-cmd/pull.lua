local Hugolib = require("hugolib")

errorlevel, errormessage = Hugolib.git_pull("origin")
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
	os.exit(errorlevel)
end

errorlevel, errormessage = Hugolib.git_submodule_init()
if errorlevel ~= 0 then
	nyagos.writerr("Error Message: "..errormessage.."\n")
end
os.exit(errorlevel)
