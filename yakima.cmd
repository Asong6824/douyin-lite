cls && cd %~dp0
@echo off
@REM 参数初始化----------------------------------------------------------------
CALL ./scripts/00.docker-set.cmd && cd %~dp0
@REM --------------------------------------------------------------------------
@REM 容器-执行
docker exec -w %DEV_PATH% -it %CON_NAME% bash -c "bash"
@REM --生成python-requirements.txt---------------------------------------------
@REM pipreqs . --encoding=utf8 --force
pause