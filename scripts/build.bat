@echo off
echo Building pm...

if not exist bin mkdir bin
go build -o bin\pm.exe .\cmd\pm

echo Build complete: bin\pm.exe
bin\pm.exe --version
pause
