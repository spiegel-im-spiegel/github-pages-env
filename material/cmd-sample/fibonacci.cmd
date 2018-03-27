@echo off
setlocal

setlocal enabledelayedexpansion

set /a n = %~1
set /a b1 = 0
set /a b2 = 0

if %n% == 1 echo 1: %b2% && goto :end
echo 1: %b2%
set /a b2 = 1
if %n% == 2 echo 2: %b2% && goto :end
echo 2: %b2%
for /l %%i in (3, 1, %n%) do (
	set /a fib = !b1! + !b2!
	set /a b1 = !b2!
	set /a b2 = !fib!
	echo %%i: !b2!
)

:end
endlocal && set /a fib = %b2%

echo %~1th Fibonacci number is %fib%
endlocal
