@echo off
setlocal

set "ROOT_DIR=%~dp0.."
set "ENV_FILE=%ROOT_DIR%\.env"

if not exist "%ENV_FILE%" (
  echo Error: .env file not found at "%ENV_FILE%"
  echo Create it first ^(for example from .env.example^).
  exit /b 1
)

echo Loading environment from "%ENV_FILE%"
for /f "usebackq tokens=1,* delims==" %%A in ("%ENV_FILE%") do (
  if not "%%A"=="" (
    if not "%%A:~0,1"=="#" (
      set "%%A=%%B"
    )
  )
)

echo Starting bot (manual run)...
go run "%ROOT_DIR%\cmd\bot"
