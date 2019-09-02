# 数据字典
> 为了方便查看数据库信息，以前生成静态数据字典，由于数据库结构变动频繁导致更新不及时，在开发过程中造成很多困扰。
> 为了解决这个问题开发了这个项目可以动态实时查看最新的数据字典
## 安装部署
### 安装mysql

以下操作均在Mac下执行

可以参考这片博客： https://blog.csdn.net/w605283073/article/details/80417866
```shell
# 安装mysql
brew install mysql
# 启动mysql
mysql.server start
# 设置密码
mysql_secure_installation
# 这里密码设置为mysql123
-- mysql123
# 进入mysql控制台
mysql -uroot -p
```
```sql
# 创建数据库
create DATABASE dbdoc;
# 切换数据库
use dbdoc
# 创建测试表
create table user
(
	id bigint auto_increment,
	nickname varchar(20) default '' not null comment '昵称',
	birthday timestamp default current_timestamp not null comment '出生日期',
	sex int default 1 not null comment '1:男 2：女',
	address varchar(200) default '' not null comment '地址',
	constraint user_pk
		primary key (id)
)
comment '用户表';

create table book
(
	id bigint auto_increment,
	bookname varchar(50) default '' not null,
	word_count int default 0 not null comment '总字数',
	auth varchar(20) default '' not null comment '作者',
	constraint book_pk
		primary key (id)
)
comment '书籍信息';
```
### 编译启动

修改配置文件，conf目录下的三个配置文件，主要用于不同环境使用

> conf-dev.yaml     ——— 本地开发环境
> conf-prod.yaml   ———  线上环境
> conf-test.yaml     ——— 测试环境

```yaml
env: test  // 环境信息

sys:
  port: 8888 // 服务启动端口

database:  // mysql 数据连接信息
  mysqlConn: root:mysql123@tcp(127.0.0.1:3306)/test_dbdoc?timeout=3s&parseTime=true&loc=Local&charset=utf8mb4
  mysqlMaxActive: 500
  mysqlMaxIdle: 20
```

配置文件修改好之后运行下面命令启动服务

```shell
go run app/main.go dev
```

上面命令会自动启动浏览器访问 http://127.0.0.1:8888/web/

如图所示

![页面截图](https://github.com/akkagao/dbdoc/blob/master/web/img/dbdoc.png?raw=true)

## 后续工作计划

- [ ] 支持修改字段注释
- [ ] 增加登陆页面
- [ ] 支持多数据库配置
- [ ] 支持PostgreSQL  