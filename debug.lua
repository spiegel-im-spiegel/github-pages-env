-- hugo.exe server --theme=text --buildDrafts --watch
nyagos.write("Start: hugo.exe server ...\n")
errorlevel, errormessage = nyagos.rawexec("hugo.exe", "server", "--theme=text", "--buildDrafts", "--watch")
if errorlevel ~= 0 then nyagos.write("Error Message: "..errormessage.."\n") end
