local Hugolib = require("hugolib")

local theme = "text"

local diskFlag = true
if #arg > 0 and arg[1] == "-m" then diskFlag = false end

nyagos.write("Start: hugo.exe server ...\n")
local errorlevel, errormessage = Hugolib.server(theme, true, diskFlag)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
