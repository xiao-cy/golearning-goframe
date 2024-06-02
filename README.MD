# learning GoFrame

Quick Start: 

- https://goframe.org/pages/viewpage.action?pageId=1114399

## config

> 配置日志&database

    golearning-goframe\hack\config.yaml

    golearning-goframe\manifest\config\config.yaml

## 项目相关表 ddl

    golearning-goframe\manifest\sql\table.sql

## CRUD

> 学习项目：
> 
>     代码生成： gf gen
> 
>     web服务开发：路由 | 请求输入 | 接口文档

- 创建用户
  
   [AdminId=3http://127.0.0.1:8000/createAdminInfo?AdminName=admin&Mobile=18512345678&Password=123456&Department=root](AdminId=3http://127.0.0.1:8000/createAdminInfo?AdminName=admin&Mobile=18512345678&Password=123456&Department=root)

- 查询用户
  
   [http://127.0.0.1:8000/getAdminInfo?AdminId=3](http://127.0.0.1:8000/getAdminInfo?AdminId=3)

- 接口文档
  
  [http://127.0.0.1:8000/swagger](http://127.0.0.1:8000/swagger)

## jwt

- git

    [GitHub - gogf/gf-jwt: GoFrame HTTP JWT middleware.](https://github.com/gogf/gf-jwt)
