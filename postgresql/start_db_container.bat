@echo off
SET container_name=postgres_db
docker ps -q -f name=%container_name% | findstr /q .
IF ERRORLEVEL 1 (
    echo Starting Docker Compose services...
    docker-compose up -d
    echo Services started.
) ELSE (
    echo %container_name% is already running.
)

@pause