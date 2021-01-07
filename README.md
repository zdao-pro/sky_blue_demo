# sky_blue_demo


## I. 服务器相关环境变量配置:

#### 1 服务器环境变量配置,详情见[env模块](https://github.com/zdao-pro/sky_blue/blob/main/pkg/env/README.md)

| 名称 | 说明 |
|:------|:------|
| DEPLOYENV | dev:开发环境 pre:预发布 online:测试 |
| HOSTNAME | 服务名 |
| APPID | 项目ID |

> shell脚本(Linux)
```bash
	export DEPLOYENV=dev
	export HOSTNAME=user
	export APPID=app
```

> powershell(win10)
```powershell
    $env:DEPLOYENV="dev"
	$env:HOSTNAME="user"
	$env:APPID="app"
```
---
#### 2 日志模块环境变量配置,详情见[log模块](https://github.com/zdao-pro/sky_blue/blob/main/pkg/log/README.md)

| 名称 | 说明 |
|:------|:------|
| UDP_LOG_ADDR | UDP远程日志中心IP地址 |
| UDP_LOG_PORT | UDP远程日志中心port |
| SYSLOG_ADDR | syslog远程日志中心IP地址 |
| SYSLOG_PORT | syslog远程日志中心port |

> shell脚本(Linux)
```bash
    export UDP_LOG_ADDR=118.178.140.41
    export UDP_LOG_PORT=1223
    export SYSLOG_ADDR=118.178.140.41
    export SYSLOG_PORT=1224
```

> powershell(win10)
```powershell
    $env:UDP_LOG_ADDR="118.178.140.41"
    $env:UDP_LOG_PORT="1223"
    $env:SYSLOG_ADDR="118.178.140.41"
    $env:SYSLOG_PORT="1224"
```

---
#### 3 配置中心环境变量配置,详情见[peach模块](https://github.com/zdao-pro/sky_blue/blob/main/pkg/peach/README.md)
| 名称 | 说明 |
|:------|:------|
| APOLLO_META_ADDR | apollo配置中心地址 |
| APOLLO_APP_ID | apollo的appid |

> shell脚本(Linux)
```bash
    export APOLLO_META_ADDR=118.178.140.41:58079
    export APOLLO_APP_ID=backend_server
```

> powershell(win10)
```powershell
    $env:APOLLO_META_ADDR="118.178.140.41:58079"
    $env:APOLLO_APP_ID="backend_server"
```

## II. 项目结构介绍
```
--- api(路由配置)
    +--- route.go
--- cmd(项目启动项)
    +--- main.go
--- config(配置目录)
--- internal
    +--- dao
         +--- mysqlDao.go
         +--- redisDao.go
    +--- model
         +--- model.go
    +--- server
         +--- http(http服务器)
              +--- server.go
--- pkg(项目自建包)
    +--- libs(项目自建库)
    +--- util(项目工具类函数库)
--- script(项目脚本目录)
--- service
    +--- service.go
```

---
## III. 开始使用(暂时需要人工操作)
```
    1. pull demo代码
       git clone https://github.com/zdao-pro/sky_blue_demo.git
    2. 替换demo中sky_blue_demo为新的工程名
       (包括go.mod & demo中其他字段)
    3. go run ./cmd/main.go
       (此时您会看到golang nb)
    4. curl -i 'http://127.0.0.1:8080/ping'
```

## IV. 其他模块介绍