
## 服务使用说明

这个服务是做成租户级别的，后期可做成ui配置界面


### 配置应用基本信息

```conf
# 集群配置
# group id 同一group id的分配到一个集群
dsl.group.id=spring.group-1
# 集群注册地址
dsl.register.url = http://192.168.0.1/add


# server 配置
# 同一group id下的一台服务，如果不写则随机分配
dsl.server.name = spring.group-1.server1
dsl.server.port = 8000
dsl.server.log.level = debug
# 加载哪些插件
dsl.server.plugins = [log,trace]
# 服务健康检查地址
dsl.server.health.check.url = http://xxxxx
# 服务健康检查 5s
dsl.server.health.check.interval=5

```

### 配置服务发现

Eureka或者其他config server的的配置接入方式

接入的rpc服务需要支持泛化调用

```conf
# eureka
eureka.name=test.application.name
eureka.port=8000
eureka.vipAddress=127.0.0.1
eureka.client.register-with-eureka=true
eureka.client.serviceUrl.defaultZone=http://localhost:8761/eureka/
# service url可以先不写，后面会自动注入
eureka.serviceUrls=[]
```

### 配置服务空间和服务

这里的未来是需要有控制台可以在页面上进行操作

```conf
## 配置服务空间

dsl.service.app = test.app
dls.service.app.

```