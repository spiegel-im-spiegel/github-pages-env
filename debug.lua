nyagos.write("Start: hugo.exe server ...\n")
local errorlevel, errormessage = nyagos.rawexec("hugo.exe", "server", "--theme=text", "--buildDrafts", "--watch")
if errorlevel ~= 0 then nyagos.writerr("Error Message: "..errormessage.."\n") end
