require("hugolib")

local pubdir = "../published"

nyagos.write("build for "..pubdir.." folder.\n")
local errorlevel, errormessage = hugolib.build("text", pubdir)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
