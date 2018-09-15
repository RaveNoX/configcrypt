@ECHO OFF
setlocal

set GOARCH=amd64

cd %~dp0
md artifacts

echo Windows
set GOOS=windows
call go build -o artifacts\jenigma.exe || goto :error

echo Linux
set GOOS=linux
call go build -o artifacts\jenigma || goto :error

echo Darwin
set GOOS=darwin
call go build -o artifacts\jenigma_darwin || goto :error

echo Build done
exit

:error
exit /b %errorlevel%

