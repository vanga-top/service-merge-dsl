## 当前能力情况

### 目前系统架构设计

1、配置驱动，所有的能力均可以通过配置化进行驱动

* 服务基本配置
* 插件的配置

application.yaml -->(create) server -->(load) plugins 


2、内建配置完成主要能力的构建

* slb plugin：主要完成服务发现和服务暴露的能力
* gateway： 主要完成请求转发，从http请求映射到后端的rpc请求
* remote：主要负责远程调用处理
* dsl：接口的业务逻辑处理

执行流程设计如下（这部分还未实现）：

http request (URI) --> gateway （router URI 2 rpc service） --> remote call rpc or http --> dsl processor --> response 

来说说这种设计存在的问题：

1、当前插件化的设计，本想相互独立，目前这样的设计存在着前后依赖的关系，搞不好，代码层面还会有依赖

2、如果代码上想没有依赖，主要两种方式：

* 以抽象接口的方式，写到一个处理的pipeline中去，初始化的时候就定义好顺序
* 写个状态机，然后用消息来进行传递

两种实现都有各自的优缺点。这里主要需要权衡的就是扩展性和复杂度的问题。


上述请求对于1对1的请求没有问题，如果考虑复杂的合并请求，更优化的请求链路应该是这样的：

http request（URI） --> gateway (router 1URI to 1DSL service) --> dsl call remote --> dsl processor --> response

1、提供一个注册 http uri和 dsl服务绑定的机制

2、将dsl服务和后端的rpc做交集，不带上http请求，这样复杂逻辑可以避免在http中传参数

3、dsl服务可以做合并和分拆，可以并发的做remote call和join

4、将结果回传

这样的设计可以处理更多的复杂情况，但是扩展性部分还是没能解决掉。


3、instance和plugin的生命周期管理和联动

* 实现了统一的生命周期管理和状态管理
* 实现了admin管理端口和gateway服务端口


### 想要怎么样的扩展能力？

还是先从用户怎么用dsl的能力开始设计，之前想过两种方式，