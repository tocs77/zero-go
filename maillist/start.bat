setlocal
FOR /F "tokens=*" %%i in ('type .env') do SET %%i
docker run --rm %IMAGE_NAME% 
endlocal