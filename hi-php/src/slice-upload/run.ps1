# 启动服务器
# starts with a specific browser
# iexplore
#Start-Process -FilePath iexplore -ArgumentList 'http://127.0.0.1:8021/index.html'


# 打开浏览器
Start-Process -FilePath 'http://127.0.0.1:8021/index.html'

# 启动服务器
php -S 0.0.0.0:8021