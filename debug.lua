require("hugolib")

nyagos.write("Start: hugo.exe server ...\n")
local errorlevel, errormessage = hugolib.server("text", true)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
