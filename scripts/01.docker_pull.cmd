cls && cd %~dp0
@echo off
@REM 参数初始化----------------------------------------------------------------
CALL 00.docker-set.cmd
set IMG_ORIG=ubuntu
set SYS_NAME=ubuntu
@REM 容器-停止所有容器的运行---------------------------------------------------
for /f "skip=1" %%i in ('docker ps -aq') do (
    docker stop %%i
)
@REM 镜像-下载-----------------------------------------------------------------
docker pull %IMG_ORIG%
@REM 容器-删除-----------------------------------------------------------------
docker rm -f %CON_NAME% >nul 2>&1 || (echo No such container & exit /b 0)
@REM 容器-创建和运行-----------------------------------------------------------
docker run -dit             ^
    --net=bridge            ^
    --restart=always        ^
    --privileged=true       ^
    --volume %VOL_NAME%     ^
    --name   %CON_NAME% %IMG_ORIG%
@REM 容器-列表-----------------------------------------------------------------
docker ps -a
@REM 暂停----------------------------------------------------------------------
pause




