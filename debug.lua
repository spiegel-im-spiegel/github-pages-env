local Hugolib = require("hugolib")

local theme = "text"

nyagos.write("Start: hugo.exe server ...\n")
local errorlevel, errormessage = Hugolib.server(theme, true)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
