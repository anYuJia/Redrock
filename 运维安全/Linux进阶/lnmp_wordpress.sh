#!/bin/bash
#把wordpress 赋值给INSTALLER_NAME变量
INSTALLER_NAME=wordpress
#wordpress下载地址
ARCHIVE_URL="https://cn.wordpress.org/latest-zh_CN.tar.gz"
#安装包名称
ARCHIVE_DIRNAME="wordpress"
#这是一个地址的变量：通过dd读取目录/dev/urandom，通过“|”管道，把读取的数据进行base64编码，报错后输入到/dev/null中，然后拷贝为16块，每块字节的文件，同样报错输入到/dev/null下
DB_PASSWD="$(dd if=/dev/urandom | (base64 -w0 2> /dev/null || base64 2> /dev/null) | dd bs=1 count=16 2>/dev/null)"

#判断是否为root用户
root_need(){
   if [[ $EUID -ne 0 ]]; then
   #如果不是root用户就不能执行，然后退出
   	echo "Error:This script must be run as root!" 1>&2
            exit 1
   fi
}


welcome(){
    echo "Hello! This is a simple script to deploy a LNMP environment."
    echo "And it also automatically install wordpress for you."
    echo "You should first set your domain name with an A DNS record to this server."
    echo
    echo "Here is the database credentials we generated for you."
    echo "PLEASE CAREFULLY SAVE THEM! VERY IMPORTANT!"
    echo "Database name: ${INSTALLER_NAME}"
    echo "Database user: ${INSTALLER_NAME}"
    echo "Database pass: ${DB_PASSWD}"
    echo
    echo "Press enter key to continue..."
    read -n 1         #从键盘读取一个数字
}

# 用sed指令用 sed 的修改结果直接修改读取数据的文件，换源讲ubuntu的源换成清华镜像源 s将···换成#之后的（字符串替换）  /后面是文件
change_source(){
    sed -i 's#archive.ubuntu.com#mirrors.tuna.tsinghua.edu.cn#g' /etc/apt/sources.list && \
    sed -i 's#ports.ubuntu.com#mirrors.tuna.tsinghua.edu.cn#g' /etc/apt/sources.list && \
    sed -i 's#security.ubuntu.com#mirrors.tuna.tsinghua.edu.cn#g' /etc/apt/sources.list && \
    #更新源
    apt update && return 0
    return 1
}

install_software(){
    #安装 wegt gnupg git curl 这几个软件
    apt install -y wget gnupg git curl && \
    # MariaDB Server Package
    #将MariaDB Server Package放入mariadb.list
    echo "deb [arch=amd64,arm64,ppc64el] http://mirrors.tuna.tsinghua.edu.cn/mariadb/repo/10.5/ubuntu focal main" > /etc/apt/sources.list.d/mariadb.list && \
    #通过wget获得文件并输入到apt-key中更新软件源
    wget -O - 'http://keyserver.ubuntu.com/pks/lookup?op=get&search=0xF1656F24C74CD1D8' | apt-key add && \
    apt update && \
    #安装相关合适的依赖
    apt install -y nginx php7.4-fpm php7.4-mysql php7.4-mbstring php7.4-gd php7.4-curl php7.4-xml php7.4-xmlrpc php7.4-iconv mariadb-server && return 0
    return 1
}
#启动数据库
start_database(){
    #启动服务
    systemctl start mariadb
    #当服务器还活着，就休眠
    while ! mysqladmin ping --silent; do
        sleep 1
    done
    #服务器停止了打印 wordpress
    echo "CREATE DATABASE \`${INSTALLER_NAME}\` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci; GRANT ALL PRIVILEGES ON \`${INSTALLER_NAME}\`.* TO '${INSTALLER_NAME}'@'localhost' IDENTIFIED BY '$DB_PASSWD';" | mysql
}

#设置端口
query_option(){
    printf "The port you want to use? [default=80]"
    #控制台读取数字赋值给port
    read -r port < /dev/tty
    #如果port为空，默认设置端口为80端口
    if [ -z "$port" ]; then
        port=80
    fi
    printf "The https port you want to use? (type 0 to disable https) [default=443]"
    #控制台读取数字赋值给port
    read -r sslport < /dev/tty
    if [ -z "$port" ]; then
    #如果port为空，默认设置端口为443端口
        sslport=443
    fi
    #如果port为0，默认设置端口为0端口
    if [ "$sslport" = "0" ]; then
        sslport=0
    fi
    printf "The domain name you want to use? [default=_]"
    #控制台读取赋值给domain
    read -r domain < /dev/tty
    if [ -z "$domain" ]; then
        domain=_
    fi
    #设置ssl_enabled为1
    ssl_enabled=1
    if [ "$sslport" = "0" ] || [ "$domain" = "_" ]; then
        ssl_enabled=0
    fi
}

#设置nginx配置
do_nginx_config(){
    cat > /etc/nginx/conf.d/${INSTALLER_NAME}.conf <<EOF
#写一个service服务
server {
    listen ${port};
    #DISABLE_SSL_PREFIX listen ${sslport} ssl http2;
    #DISABLE_SSL_PREFIX ssl_certificate /var/www/ssl/${domain}/public.cer;
    #DISABLE_SSL_PREFIX ssl_certificate_key /var/www/ssl/${domain}/private.key;
    server_name ${domain};
    index index.php;
    root /var/www/${INSTALLER_NAME};
    location = /favicon.ico {
        log_not_found off;
        access_log off;
    }
    location = /robots.txt {
        allow all;
        log_not_found off;
        access_log off;
    }
    location / {
        # This is cool because no php is touched for static content.
        # include the "?\$args" part so non-default permalinks doesn't break when using query string
        try_files \$uri \$uri/ /index.php?\$args;
    }
    location ~ \.php\$ {
        #NOTE: You should have "cgi.fix_pathinfo = 0;" in php.ini
        include snippets/fastcgi-php.conf;
        fastcgi_pass unix:/var/run/php/php7.4-fpm.sock;
        fastcgi_intercept_errors on;
        fastcgi_buffers 16 16k;
        fastcgi_buffer_size 32k;
    }
    location ~* \.(js|css|png|jpg|jpeg|gif|ico)\$ {
        expires max;
        log_not_found off;
    }
    gzip on;
}
EOF
    #重启nginx服务
    systemctl restart nginx
    #如果$ssl_enabled的值是1，即前面两个设置未通过。则进行如下代码
    if [ "$ssl_enabled" = "1" ]; then
        echo "Dowlnoading certbot..."
        git clone https://codechina.csdn.net/mirrors/acmesh-official/acme.sh.git ~/.acme.sh
        echo "Issueing certificate..."
        # 创建一个文件夹，递归创建文件夹
        mkdir -p /var/www/ssl/${domain}
        (~/.acme.sh/acme.sh --issue --server letsencrypt -d "${domain}" -w /var/www/${INSTALLER_NAME} \
            --key-file "/var/www/ssl/${domain}/private.key" --fullchain-file "/var/www/ssl/${domain}/public.cer" --reloadcmd "systemctl force-reload nginx" \
        && sed -i 's/#DISABLE_SSL_PREFIX\ //g' /etc/nginx/conf.d/${INSTALLER_NAME}.conf \
	&& systemctl force-reload nginx) || echo "apply the certification fail!"#启动服务
    fi
}
#判断是否为root用户
root_need
welcome
#换源
change_source
if [ $? -eq 0 ]; then
    echo "Change source success to tsinghua.edu.cn"
else
    echo "Change source failed, use the default source."
fi
#安装软件
install_software
if [ $? -eq 0 ]; then
    echo "Install software success"
else
    echo "Install software failed"
    exit 1
fi

#关闭apache2的开机自启
systemctl disable apache2
#让nginx和php环境 mariadb开机自启
systemctl enable nginx php7.4-fpm mariadb
#让nginx服务停止
systemctl stop nginx
#启动数据库，若启动失败则退出
start_database || (echo "Start database failed" && exit 1)

cd /var/www/ && \
wget -O - "${ARCHIVE_URL}" | tar -xzf - && \
mv "${ARCHIVE_DIRNAME}" "${INSTALLER_NAME}" && \
chown -R www-data:www-data "${INSTALLER_NAME}" && \

query_option

do_nginx_config

echo "Here is the database credentials we generated for you, shown again."
echo "PLEASE CAREFULLY SAVE THEM! VERY IMPORTANT!"
echo "Database name: ${INSTALLER_NAME}"
echo "Database user: ${INSTALLER_NAME}"
echo "Database pass: ${DB_PASSWD}"
