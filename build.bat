@echo off
setlocal enabledelayedexpansion

:: Configuration
set BINARY_NAME=rbserver.exe
set BUILD_DIR=build
set MAIN_PATH=./cmd/rbserver

if "%1"=="" goto help
if "%1"=="build" goto build
if "%1"=="release" goto release
if "%1"=="clean" goto clean

:help
echo Usage: build.bat [command]
echo.
echo Commands:
echo   build    Compile for Windows
echo   release  Cross-compile for Linux and Mac
echo   clean    Remove the bin directory
goto :eof

:build
echo Building %BINARY_NAME%...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%
go build -o %BUILD_DIR%\%BINARY_NAME% %MAIN_PATH%
if %ERRORLEVEL% EQU 0 (
    echo Done! Binary located at: %BUILD_DIR%\%BINARY_NAME%
) else (
    echo ❌ Build failed.
)
goto :eof

:release
echo Cross-compiling for all platforms...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

echo Building Windows...
set GOOS=windows
set GOARCH=amd64
go build -o %BUILD_DIR%\rbserver_windows_amd64.exe %MAIN_PATH%

echo Building Linux...
set GOOS=linux
set GOARCH=amd64
go build -o %BUILD_DIR%\rbserver_linux_amd64 %MAIN_PATH%

echo Building MacOS...
set GOOS=darwin
set GOARCH=arm64
go build -o %BUILD_DIR%\rbserver_macos_arm64 %MAIN_PATH%

:: Reset env vars
set GOOS=
set GOARCH=
echo Release binaries ready in %BUILD_DIR%
goto :eof

:clean
if exist %BUILD_DIR% (
    rd /s /q %BUILD_DIR%
    echo Cleaned.
) else (
    echo Nothing to clean.
)
goto :eof
