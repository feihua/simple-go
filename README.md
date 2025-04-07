# 基于Gin + Ant Design Pro的前后端分离管理系统

## 快速启动

### 1.下载代码
```shell
git clone https://github.com/feihua/simple-go.git
```

### 2.修改配置
vim config/app.yaml
```yaml
# MySQL数据库的连接配置
mysql:
    host: 127.0.0.1
    port: 3306
    database: simple-go
    username: root
    password: 123456

# Redis数据库的连接配置
redis:
    host: 127.0.0.1:6379
    password: 123456
```

### 3.启动
```shell
make && make restart
```

### 4.验证
```shell
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{
             "mobile": "18613030352",
             "password": "123456"
         }' \
     "{{host}}:6677/api/user/login"

```
请将{{host}}替换为实际的服务器主机地址或域名。这段cURL命令执行