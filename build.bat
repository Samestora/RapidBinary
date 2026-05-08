@echo off
setlocal enabledelayedexpansion

:: Configuration
set BUILD_DIR=build
:: List of apps in cmd folder to iterate through
set APPS=rbserver rbhash

if "%1"=="" goto help
if "%1"=="build" goto build
if "%1"=="release" goto release
if "%1"=="clean" goto clean

:help
echo Usage: build.bat [command]
echo.
echo Commands:
echo   build    Compile all apps for Windows (local)
echo   release  Cross-compile all apps for Linux, Windows, Mac
echo   clean    Remove the build directory
goto :eof

:build
echo Building for Windows...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

for %%A in (%APPS%) do (
    echo   - Building %%A...
    :: Builds from the respective cmd subdirectory
    go build -ldflags="-s -w" -o %BUILD_DIR%\%%A.exe ./cmd/%%A
    if !ERRORLEVEL! EQU 0 (
        echo      %%A ready
    ) else (
        echo      %%A failed
    )
)
goto :eof

:release
echo Cross-compiling RapidBinary Toolsuite...
if not exist %BUILD_DIR% mkdir %BUILD_DIR%

:: Windows Build
set GOOS=windows
set GOARCH=amd64
for %%A in (%APPS%) do (
    echo   - Windows: %%A
    go build -ldflags="-s -w" -o %BUILD_DIR%/%%A_windows_amd64.exe ./cmd/%%A
)

:: Linux Build
set GOOS=linux
set GOARCH=amd64
for %%A in (%APPS%) do (
    echo   - Linux: %%A
    go build -ldflags="-s -w" -o %BUILD_DIR%/%%A_linux_amd64 ./cmd/%%A
)

:: MacOS Build (Silicon/ARM)
set GOOS=darwin
set GOARCH=arm64
for %%A in (%APPS%) do (
    echo   - MacOS: %%A
    go build -ldflags="-s -w" -o %BUILD_DIR%/%%A_macos_arm64 ./cmd/%%A
)

:: Reset environment variables
set GOOS=
set GOARCH=
echo.
echo Release binaries ready in /%BUILD_DIR%
goto :eof

:clean
if exist %BUILD_DIR% (
    rd /s /q %BUILD_DIR%
    echo Cleaned.
) else (
    echo Nothing to clean.
)
goto :eof
