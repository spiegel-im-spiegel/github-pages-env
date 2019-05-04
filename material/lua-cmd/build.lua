local Hugolib = require("hugolib")

local pubdir = "../published"
local theme = "baldanders-info"

nyagos.write("build for "..pubdir.." folder.\n")
local errorlevel, errormessage = Hugolib.build(theme, pubdir)
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
os.exit(errorlevel)
