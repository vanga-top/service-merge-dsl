# service-merge-dsl
面向服务化架构的服务拼装DSL层


## 目标

1、 解决低代码服务或者业务服务化架构下数据扩展性的问题

2、 解决多服务数据拼装问题（特别给端上使用，减少前后端扯皮）

3、 降数据渲染、模版渲染都放在服务端，减少端上压力

4、 支持静态化

5、 支持市面上主流的服务框架和服务发现，和多语言环境

6、 支持自定义服务协议和调用框架


## 相关文档

由于不是集中式的开发，经常会被打断，所以会有一些思路上的断层，设计文档会按照思考的迭代也做一个版本的迭代。

系统设计v0.1 [系统设计](docs/arch.md)

概要设计v0.2 [概要设计，详细一点的struct设计](docs/arch_01.md)

更新设计v0.3 [关于当前设计的一些思考，在引入gateway之后](docs/arch_02.md)