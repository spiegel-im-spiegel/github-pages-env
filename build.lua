require("hugolib")

nyagos.write("build...\n")
local errorlevel, errormessage = hugolib.build()
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
