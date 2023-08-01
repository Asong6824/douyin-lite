# 设置要使用的用户名和密码
USERNAME="asong"
PASSWORD="mysql201209"

# 执行 MySQL 命令来创建用户和设置密码
mysql -u root <<EOF
CREATE USER '$USERNAME'@'localhost' IDENTIFIED BY '$PASSWORD';
GRANT ALL PRIVILEGES ON *.* TO '$USERNAME'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF

# 输出结果
echo "MySQL 用户 '$USERNAME' 已创建，密码为 '$PASSWORD'"