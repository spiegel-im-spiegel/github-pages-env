require("hugo")

local errorlevel, errormessage = hugo.build()
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
