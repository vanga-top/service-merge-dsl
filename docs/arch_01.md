## how to use

### 启动service-merge-dsl服务

1、配置服务的基础信息

服务基础信息包含：应用分组名、应用名、保活信息等

```yaml
merge-dsl:
  application:
    name:application-name
    group-id:service.group.1
    port:8080
  slb:
    host:http://106.14.167.79:8761/eureka/ #这个slb和微服务的服务发现地址一致
    interval:30  #心跳间隔时间 单位s
  log:
    level:debug
```

相关设计：

#### application


```go

```


#### slb-client

需要有一个eureka的client，用于注册到eureka-server。这边可以抽象出一个通用的slb-client配置，因为未来除了有eureka之外，还有可能存在其他形式的client。
另外，未来应该可以支持多slb的方式。

```go
// SLBClient Service Load Balance Client
type SLBClient interface {
    Connect(url string, port int, namespace string, opts *SLBOptions) *SLBResult
    DisConnect() *SLBResult
}

type SLBResult struct {
    Code      int
    Success   bool
    Namespace string
}

type SLBOptions struct {
    Username string
    Password string
    Token    string
}
```

注意：这边有个设计需要讨论一下

1、如果dsl作为gateway的一个插件来进行使用，可以直接拿gateway的slb上下文，只需要拿到service provider的地址即可发起调用，由于gateway会做lb，所以并不需要dsl层关心lb的问题

2、而如果dsl作为一个独立的服务部署在外层，直接面对客户端。则需要考虑load balance谁来做。

    2.1 如果是挂个nginx集群，有nginx来分配到不同的slb，则需要slb提供出一个http server供nginx来进行链接
    2.2 如果连nginx也不挂，那dsl需要考虑集群的设计，如选主逻辑等（这种模式暂不考虑）

最次的情况下，lb由上层来做，dsl还是作为一个无状态的服务。

2、启动服务

go app -c application.yaml

### 配置服务

服务的配置分几种方式：

* ui界面配置，采用web控制台直接配置（推荐方式）
* 采用config文件配置
* 直接调用api进行配置


2、采用config文件配置的方式

```config
    
```

### 服务调用

两种模式：

* 同步调用
* 异步调用