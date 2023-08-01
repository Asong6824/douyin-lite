#!/usr/bin/bash
clear
SELF_FOLDER="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )" && cd "${SELF_FOLDER}" || exit
CARGOTAINER="${SELF_FOLDER}"/__cargotainer__
SYSTEMFILES="${SELF_FOLDER}"/__systemfiles__
# --判断当前用户是否有执行权限---------------------------------------------------------------------
uid=$(id -u)
if [ "$uid" != 0 ]; then
    echo "当前用户无执行权限，系统退出..."
    exit 1
fi
# --选择国内源或者国外源---------------------------------------------------------------------------
PS3='please select apt source: '
menu=("chinese" "foreign")
select option in "${menu[@]}"; do
    case $option in
        "foreign")
            SOURCE_NAME="foreign"
            ;;
        "chinese")
            SOURCE_NAME="chinese"
            ;;
        *)  SOURCE_NAME="chinese"
    esac
    echo ---------------- selected apt source is "${SOURCE_NAME}" ----------------
    break
done
# 读取 /etc/os-release 文件中的 ID 字段,判断发行版并赋值给变量-------------------------------------
oid=$(grep "^ID=" /etc/os-release | cut -d= -f2)
if   [ "$oid" = "debian" ]; then
    cp "${SYSTEMFILES}"/__debian__/bulleye_"${SOURCE_NAME}".list /etc/apt/sources.list
elif [ "$oid" = "ubuntu" ]; then
    cp "${SYSTEMFILES}"/__ubuntu__/jammy-"${SOURCE_NAME}".list   /etc/apt/sources.list
    cp "${SYSTEMFILES}"/__ubuntu__/supervisor                    /etc/sudoers.d/supervisor
    cp "${SYSTEMFILES}"/__ubuntu__/10periodic                    /etc/apt/apt.conf.d/10periodic
    cp "${SYSTEMFILES}"/__ubuntu__/20auto-upgrades               /etc/apt/apt.conf.d/20auto-upgrades
else
    echo "Unknown distribution."
fi
# -------------------------------------------------------------------------------------------------
find /var/lib/apt/lists/lock     | awk '{ print $1 }' | xargs -r rm -rf
find /var/lib/dpkg/lock-frontend | awk '{ print $1 }' | xargs -r rm -rf
find /var/lib/dpkg/updates/*     | awk '{ print $1 }' | xargs -r rm -rf
find /var/lib/dpkg/lock          | awk '{ print $1 }' | xargs -r rm -rf
apt update -y --fix-missing && apt upgrade -y && apt autoremove -y
# -------------------------------------------------------------------------------------------------
echo -e "\e[92m----------------------------------S1.系统工具---------------------------------------\e[39m"
apt install -y git curl wget sudo nano lsof htop makeself dmidecode shellcheck bsdmainutils
# -------------------------------------------------------------------------------------------------
echo -e "\e[92m----------------------------------S2.网络工具---------------------------------------\e[39m"
apt install -y netcat ntpdate net-tools traceroute netscript-2.4 inetutils-ping dnsutils
# -------------------------------------------------------------------------------------------------
echo -e "\e[92m----------------------------------S3.数据工具.mysql--------------------------------\e[39m"
apt install -y mysql-server
# -------------------------------------------------------------------------------------------------
echo -e "\e[92m----------------------------------S4.开发工具.golang--------------------------------\e[39m"
rm -rf /usr/local/go && tar -C /usr/local -xzf "${CARGOTAINER}"/go1.20.6.linux-amd64.tar.gz
if ! grep -qxF "export PATH=\$PATH:/usr/local/go/bin" ~/.bashrc; then
    sed -i '$ a\export PATH=\$PATH:/usr/local/go/bin' ~/.bashrc
    source ~/.bashrc
fi
# --容器中：---------------------------------------------------------------------------------------
if grep -q '/docker/' /proc/self/cgroup; then
    cp "${SYSTEMFILES}"/Shanghai /etc/localtime && date && hwclock
# --非容器：---------------------------------------------------------------------------------------
else
    # --移除工具：错误报告-------------------------------------------------------------------------
    apt autoremove -y apport
    # --安装工具-----------------------------------------------------------------------------------
    sed -i 's/#WaylandEnable=false/WaylandEnable=false/g' /etc/gdm3/custom.conf
    apt install -y xorg openssh-server openssh-sftp-server
    # apt install -y "${CARGOTAINER}"/rustdesk*.deb
    apt install -y "${CARGOTAINER}"/todesk*.deb
    # --交换空间-----------------------------------------------------------------------------------
    swapoff /swapfile           # 禁用交换空间
    rm -rf /swapfile            # 删除交换空间文件
    fallocate -l 16G /swapfile  # 创建一个用于交换的文件
    chmod 600 /swapfile         # 将文件的权限设置为仅限 root 用户读写 
    mkswap /swapfile            # 将文件设置为交换空间
    swapon /swapfile            # 启用交换空间
    grep -qF '/swapfile' /etc/fstab || echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab   # 系统启动时自动启用
fi
# --清空历史记录并重启-----------------------------------------------------------------------------
apt -f install && echo "" > ~/.bash_history && reboot





















# --暂时不用---------------------------------------------------------------------------------------
# pip install git+https://github.com/EntySec/Shreder                  # SSH brute
# pip install git+https://github.com/EntySec/CamOver                  # camera
# pip install git+https://github.com/EntySec/CamRaptor                # camera
# pip install git+https://github.com/EntySec/RomBuster                # router
# apt install -y ruby-full 
# --系统工具---------------------------------------------------------------------------------------
# apt install -y systemctl
# -------------------------------------------------------------------------------------------------
# echo -e "\e[92m-------------------------------------开发工具.golang--------------------------------\e[39m"
# rm -rf /usr/local/go && tar -C /usr/local -xzf "${CARGOTAINER}"/go1.20.6.linux-amd64.tar.gz
# if ! grep -qxF "export PATH=\$PATH:/usr/local/go/bin" ~/.bashrc; then
#     sed -i '$ a\export PATH=\$PATH:/usr/local/go/bin' ~/.bashrc
#     source ~/.bashrc
# fi
# --开发工具.ruby-full-----------------------------------------------------------------------------
# apt install -y ruby-full squashfs-tools
# gem sources -r https://rubygems.org/ -a https://gems.ruby-china.com/        # 移除gem默认源，改成ruby-china源
# gem sources -u                                                              # 更新源的缓存
# gem sources -l                                                              # 检查源的状态
# gem install bundler
# bundler config mirror.https://rubygems.org https://gems.ruby-china.com      # 使用Gemfile和Bundle的项目，可以做下面修改，就不用修改Gemfile的source
# bundler config --delete 'mirror.https://rubygems.org'                       # 删除Bundle的一个镜像源
# cp "${CARGOTAINER}"/rubyc* /usr/local/ && chmod a+x /usr/local/rubyc*       # 复制rubyc编译器
# --资产工具.assets--------------------------------------------------------------------------------
# apt install -y snmp snmpd libsnmp-dev snmp-mibs-downloader onesixtyone && gem install snmp
# --漏扫工具.vulner--------------------------------------------------------------------------------
# go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest


    # # --安装容器-----------------------------------------------------------------------------------
    # apt install -y docker docker-compose
    # # --克隆镜像-----------------------------------------------------------------------------------
    # docker search portainer
    # docker pull portainer/portainer
    # docker run -dit --restart=always --publish=9000:9000 --volume=/var/run/docker.sock:/var/run/docker.sock --name=matrix-portainer portainer/portainer
