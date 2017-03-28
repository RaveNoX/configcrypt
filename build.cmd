@ECHO OFF
setlocal

set GOARCH=amd64

cd %~dp0
md artifacts

echo Linux
set GOOS=linux
call go build -o artifacts\jenigma
if not %ERRORLEVEL% == 0 (exit %ERRORLEVEL%)

echo Windows
set GOOS=windows
call go build -o artifacts\jenigma.exe
if not %ERRORLEVEL% == 0 (exit %ERRORLEVEL%)

echo Darwin
set GOOS=darwin
call go build -o artifacts\jenigma_darwin
if not %ERRORLEVEL% == 0 (exit %ERRORLEVEL%)

echo Build done
