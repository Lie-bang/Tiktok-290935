# Douyin（第五届字节跳动后端青训营项目）👏 
![](https://img.shields.io/badge/lisence-MIT-yellow)
![](https://img.shields.io/badge/build-1.0-orange)
![](https://img.shields.io/badge/go-1.19.5-blue)

## 项目简介🌟

采用Hertz/Kitex微服务架构进行项目开发，合理拆解各项业务便于团队协作开发，并使用Etcd进行服务注册与服务发现，利用Minio分布式对象存储来存储数据，采用MySQL、Redis相结合的方式最大化存储效率，实现了一个简易版抖音。

仓库: https://github.com/Lie-bang/Tiktok-290935

文档：https://rmlso9nib5.feishu.cn/docx/QHWmdNvWQokvTsx3ZsMcx6kcnWd

## 项目技术栈

### RPC框架使用Kitex
Kitex 是字节跳动内部的 Golang 微服务 RPC 框架，具有高性能、强可扩展的特点，在字节内部已广泛使用。如果对微服务性能有要求，又希望定制扩展融入自己的治理体系，Kitex 会是一个不错的选择。

### HTTP框架使用Hertz
Hertz 是一个 Golang 微服务 HTTP 框架，在设计之初参考了其他开源框架 fasthttp、gin、echo 的优势， 并结合字节跳动内部的需求，使其具有高易用性、高性能、高扩展性等特点，目前在字节跳动内部已广泛使用。 如今越来越多的微服务选择使用 Golang，如果对微服务性能有要求，又希望框架能够充分满足内部的可定制化需求，Hertz 会是一个不错的选择。

### 关系型数据库使用MySQL
MySQL是一种关系型数据库管理系统，由瑞典MySQL AB 公司开发，目前属于 Oracle 旗下产品。MySQL 是最流行的关系型数据库管理系统之一，也是最好的 RDBMS (Relational Database Management System，关系数据库管理系统) 应用软件之一。

### 非关系型数据库使用Redis
Redis是一个开源的内存数据库，它可以用作数据库、缓存和消息中间件。它支持多种数据结构，如字符串、哈希、列表、集合、有序集合等，可以用于存储各种类型的数据。它的主要特点是高性能、高可用性和高可扩展性。

### 服务注册与服务发现使用Etcd
Etcd是一个分布式键值存储系统，用于在分布式系统中存储和共享配置和服务发现信息。它是一个高可用的、可扩展的、高性能的分布式键值存储系统，可以用于存储和管理大量的配置信息。

### 分布式对象存储使用Minio
Minio是一个开源的对象存储服务，可以用于存储大量非结构化数据，如图片、视频、日志文件、备份数据等。它支持多种标准协议，如Amazon S3、OpenStack Swift和Google Cloud Storage，可以轻松部署在本地或云环境中。

## 项目架构

![image](https://github.com/Lie-bang/Tiktok-290935/blob/main/architecture.jpg)

## 项目结构

<details>
<summary>点击展开</summary>
<pre>
<code>
       ├─ .DS_Store
       ├─ .git
       ├─ .idea
       ├─ cmd
       │    ├─ api    //apiService 核心代码
       │    │    ├─ .gitignore
       │    │    ├─ .hz
       │    │    ├─ api    //解析业务逻辑
       │    │    ├─ biz
       │    │    ├─ fast.sh
       │    │    ├─ main.go //apiService启动
       │    │    ├─ router.go
       │    │    └─ router_gen.go
       │    ├─ comment    //commentService 核心代码
       │    │    ├─ build.sh
       │    │    ├─ comment
       │    │    ├─ dal    //commentService相关数据操作
       │    │    ├─ handler.go
       │    │    ├─ kitex.yaml
       │    │    ├─ main.go //rpcService启动
       │    │    ├─ output
       │    │    ├─ pack    //对象打包
       │    │    ├─ rpc    //申请rpc调用其他服务
       │    │    ├─ script
       │    │    └─ service //核心业务逻辑代码
       │    ├─ favorite    //favoriteService 核心代码
       │    │    ├─ build.sh
       │    │    ├─ dal    //favoriteService相关数据操作
       │    │    ├─ favorite
       │    │    ├─ handler.go
       │    │    ├─ kitex.yaml
       │    │    ├─ main.go //rpcService启动
       │    │    ├─ output
       │    │    ├─ pack    //对象打包
       │    │    ├─ rpc    //申请rpc调用其他服务
       │    │    ├─ script
       │    │    └─ service //核心业务逻辑代码
       │    ├─ message    //messageService 核心代码
       │    │    ├─ build.sh
       │    │    ├─ dal    //Service相关数据操作
       │    │    ├─ handler.go
       │    │    ├─ kitex.yaml
       │    │    ├─ main.go //rpcService启动
       │    │    ├─ message
       │    │    ├─ output
       │    │    ├─ pack    //对象打包
       │    │    ├─ script
       │    │    └─ service //核心业务逻辑代码
       │    ├─ relation    //relationService 核心代码
       │    │    ├─ build.sh
       │    │    ├─ dal    //Service相关数据操作
       │    │    ├─ handler.go
       │    │    ├─ kitex.yaml
       │    │    ├─ main.go //rpcService启动
       │    │    ├─ output
       │    │    ├─ pack    //对象打包
       │    │    ├─ relation
       │    │    ├─ rpc    //申请rpc调用其他服务
       │    │    ├─ script
       │    │    └─ service //核心业务逻辑代码
       │    ├─ user    //UserService 核心代码
       │    │    ├─ build.sh
       │    │    ├─ dal    //Service相关数据操作
       │    │    ├─ handler.go
       │    │    ├─ kitex.yaml
       │    │    ├─ main.go //rpcService启动
       │    │    ├─ output
       │    │    ├─ pack    //对象打包
       │    │    ├─ rpc    //申请rpc调用其他服务
       │    │    ├─ script
       │    │    ├─ service //核心业务逻辑代码
       │    │    └─ user
       │    └─ video    //VideoService 核心代码
       │           ├─ Service //核心业务逻辑代码
       │           ├─ build.sh
       │           ├─ dal    //Service相关数据操作
       │           ├─ handler.go
       │           ├─ kitex.yaml
       │           ├─ main.go //rpcService启动
       │           ├─ output
       │           ├─ pack
       │           ├─ rpc    //申请rpc调用其他服务
       │           ├─ script
       │           ├─ tempVideoFile
       │           └─ video
       ├─ docker-compose.yaml  // minio;mysql;redis;etcd docker配置文件
       ├─ go.mod
       ├─ go.sum
       ├─ idl //idl文件
       │    ├─ api.thrift
       │    ├─ comment.thrift
       │    ├─ favorite.thrift
       │    ├─ message.thrift
       │    ├─ relation.thrift
       │    ├─ user.thrift
       │    └─ video.thrift
       ├─ kitex_gen //kitex生成胶水代码
       └─ pkg
              ├─ configs //配置文件
              │    ├─ nosql
              │    ├─ sql
              │    └─ otel
              ├─ consts //常量定义文件（serviceName/ServiceAddr/localIP/...）
              │    └─ consts.go
              ├─ errno    //错误码定义
              │    └─ errno.go
              ├─ minio    //minio配置文件
              │    ├─ init.go
              │    └─ minio_op
              └─ mw
              │    └─ client.go
              │    └─ common.go
              │    └─ server.go
              └─ static //静态资源
</code>
</pre>
</details>

## 项目部署⚡

docker-compose
