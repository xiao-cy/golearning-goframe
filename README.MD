# learning GoFrame

Quick Start: 

- https://goframe.org/pages/viewpage.action?pageId=1114399

## config

> 配置日志&database

    golearning-goframe\hack\config.yaml

    golearning-goframe\manifest\config\config.yaml

## 项目相关表 ddl

    golearning-goframe\manifest\sql\table.sql

## Postman JSON

> 包含项目所有接口案例

golearning-goframe\manifest\postman\golearning-goframe.postman_collection.json

## CRUD

> 学习项目：
> 
>     代码生成： gf gen
> 
>     web服务开发：路由 | 请求输入 | 参数校验 | 接口文档

- 创建用户
  
  接口：/createAdminInfo
  
  参数：AdminName=xxx&Mobile=xxx&Password=xxx&Department=xxx

- 查询用户
  
  接口：/getAdminInfo?AdminId=xx

- 接口文档
  
  [http://127.0.0.1:8000/swagger](http://127.0.0.1:8000/swagger)

## jwt

- git

    [GitHub - gogf/gf-jwt: GoFrame HTTP JWT middleware.](https://github.com/gogf/gf-jwt)

- 获取token
  
  /admin/login

- 刷新token
  
  /admin/refresh_token

- 删除token
  
  /admin/logout

## session

- 登录
  
   /session/login

- 登出
  
  /session/logout

- 登录可访问
  
   /damo

  
