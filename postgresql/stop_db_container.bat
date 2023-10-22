@echo off
SET container_name=postgres_db
docker ps -q -f name=%container_name% | findstr /q .
IF ERRORLEVEL 1 (
    echo %container_name% is not running.
) ELSE (
    echo Stopping Docker Compose services...
    docker-compose down
    echo Services stopped.
)

@pause