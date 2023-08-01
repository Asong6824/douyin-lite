# 定义MySQL登录信息
MYSQL_HOST="127.0.0.1"  # MySQL主机
MYSQL_PORT="3306"
MYSQL_USER="asong"  # MySQL用户名
MYSQL_PASSWORD="mysql201209"  # MySQL密码
MYSQL_DATABASE="douyin"  # 数据库名

mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD -e "CREATE DATABASE $MYSQL_DATABASE;"

SQL_FILES=(
    "./db/create_users.sql"
    "./db/create_follow_relations.sql"
    "./db/create_videos.sql"
    "./db/create_video_likes.sql"
    "./db/create_video_comments.sql"
    # 添加更多的 SQL 文件路径，每个路径占一行
)

# 循环遍历 SQL 文件列表并执行每个文件
for file in "${SQL_FILES[@]}"
do
    echo "Executing SQL file: $file"
    
    # 使用命令行工具（如 mysql）执行 SQL 文件
    # -h 为数据库主机
    # -P 为数据库端口
    # -u 为数据库用户名
    # -p 为数据库密码
    # < 为输入重定向符号，将 SQL 文件作为输入
    mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE < $file

    if [ $? -eq 0 ]; then
        echo "SQL file executed successfully"
    else
        echo "Failed to execute SQL file"
        exit 1
    fi
done