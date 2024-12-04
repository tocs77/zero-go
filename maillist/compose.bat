setlocal
FOR /F "tokens=*" %%i in ('type .env') do SET %%i
docker compose up --build
endlocal