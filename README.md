## 功能设计与实现

![1700312747480](image/reademe/1700312747480.png)

### 1. 策略管理

数据库设计：

```
表名：share_files
字段：
id : bigint unsigned 自增主键
userId : bigint unsigned 用户ID
fileId : bigint unsigned 文件ID
desc : varchar(200) # 策略描述 
target_user_id : bigint unsigned # 目标用户ID
expire : DateTime # 过期时间 默认值为当前时间增加一天
use_limit : int  unsigned # 使用次数 默认值为10
is_delete : tinyint unsigned 是否删除 0 未删除 1 已删除 默认值为0
gmt_create: datetime # 创建时间 默认值为当前时间
gmt_modified : datetime # 更新时间 默认值为当前时间
```

创建表语句：  

```SQL
CREATE TABLE `share_files` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `userId` bigint(20) unsigned NOT NULL COMMENT '用户ID',
    `fileId` bigint(20) unsigned NOT NULL COMMENT '文件ID',
    `desc` varchar(200) NOT NULL COMMENT '策略描述',
    `target_user_id` bigint(20) unsigned NOT NULL COMMENT '目标用户ID',
    `expire` bigint NOT NULL default 1737302400000  COMMENT '过期时间 默认值为当前时间增加一天',
    `use_limit` int(10) unsigned NOT NULL DEFAULT 10 COMMENT '使用次数 默认值为10',
    `is_delete` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT '是否删除 0 未删除 1 已删除 默认值为0',
    `gmt_create` datetime NOT NULL COMMENT '创建时间 默认值为当前时间',
    `gmt_modified` datetime NOT NULL COMMENT '更新时间 默认值为当前时间',
    PRIMARY KEY (`id`),
    KEY `idx_userId` (`userId`),
    KEY `idx_fileId` (`fileId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分享文件表';
```
### 2. 用户管理

用户表设计

```
表名：user
user_id : bigint unsigned 用户ID 自增主键
user_name : char(100) # 用户名称
password : char(100) # 用户密码
role_id : int unsigned # 角色ID 管理员 1 普通用户 0 默认为0
public_key : char(100) # 公钥 默认为dog
secret_key : char(100) # 私钥 默认为dog_secret
is_delete : tinyint unsigned 是否注销 0 未删除 1 已删除 默认为0
gmt_create: datetime # 创建时间 默认为当前时间
gmt_modified : datetime # 更新时间 默认为当前时间
```

建表语句

```SQL
# 创建user表并在创建语句中添加注释并且添加默认值
CREATE TABLE `user` (
  `user_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID 自增主键',
  `user_name` char(100) NOT NULL COMMENT '用户名称',
  `password` char(100) NOT NULL COMMENT '用户密码', 
  `role_id` int unsigned NOT NULL DEFAULT 0 COMMENT '角色ID 管理员 1 普通用户 0 默认为0',
  `public_key` char(100) NOT NULL DEFAULT 'dog' COMMENT '公钥 默认为dog',
  `secret_key` char(100) NOT NULL DEFAULT 'dog_secret' COMMENT '私钥 默认为dog_secret',
  `is_delete` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否注销 0 未删除 1 已删除 默认0',
  `gmt_create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间 默认当前时间',
  `gmt_modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间 默认当前时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';
```
用户模块功能设计：
- 用户注册
- 用户登录
- 密钥（扩充用户分享链）
  - 重新生成公钥与私钥

### 3.病例管理

病例表设计：只记录管理员上传的文件
```
表名：share_files
file_id : bigint unsigned 病例ID 自增主键
file_name : char(100) # 病例名称
file_size : int unsigned # 病例大小
use_count : int unsigned # 文件被使用次数 默认为0次
owner_id : bigint unsigned # 拥有者用户ID 
is_delete : tinyint unsigned 是否注销 0 未删除 1 已删除 默认为0
gmt_create: datetime # 上传时间 默认为当前时间
gmt_modified : datetime # 更新时间 默认为当前时间
```

建表语句

```SQL
# 创建share_files表并在创建语句中添加注释并且添加默认值 
CREATE TABLE `share_files`(
  `file_id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '病例ID 自增主键',
  `file_name` char(100) NOT NULL  COMMENT '病例名称',
  `file_size` int unsigned NOT NULL COMMENT '病例大小',
  `use_count` int unsigned NOT NULL DEFAULT 0 COMMENT '文件被使用次数',
  `owner_id` bigint unsigned NOT NULL  COMMENT '拥有者用户ID',
  `is_delete` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否注销 0 未删除 1 已删除 默认为0',
  `gmt_create` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间 默认为当前时间',
  `gmt_modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间 默认为当前时间',
  PRIMARY KEY (`file_id`),
  UNIQUE KEY `file_name` (`file_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='共享病例表';
```

## 更新项目执行流程
- 分角色
  - 管理员
    - 登录
    - 上传病例
  - 普通用户
    - 注册
    - 登录
    - 病历管理
      - 查看病历
        - 查看已共享与被共享病历
      - 共享病历
        - 创建共享策略进行共享


## 开发日志记录

11/18 ：

- 数据库表设计，完善需求分析

  - 用户表 + 用户策略表重构
- 代码处理部分进行微重构

  - 重构代码，完善异常处理
- 编写用户注册模块 + 微重构用户登录模块

  - 用户注册模块
    - 用户名唯一性校验
    - 密码加密
    - 用户注册
  - 用户登录模块微重构
    - JWT中存放加密后的用户对象，方便在后续过程中获取公钥与用户ID
    - 用户登录时，将用户对象加密后存入JWT中
    - 用户登录

11/20：

- 编写用户策略模块
  - 用户创建新策略
  - 用户删除策略
  - 用户修改策略
  - 用户分享病例时获取自身创建策略
  - 管理员查看用户策略
- 数据库表设计 + 病例管理需求分析

12/14
- 首次前后端联调，实现用户注册与登录模块，初步完善最初设计
- 增加管理员端 上传文件选项（管理员上传文件之后，用户可以进行共享以及创建相关的策略）
- 在不改变原有基础上进行扩展，不进行大幅度修改！
- 后端可以增加接口，前端尽量不要动

12/16
- 完成文件上传与预览功能，可以暂时完成系统的运行，基本框架完成构建
- 规划后续任务：病历共享，围绕基本框架进行展开

12/18
- 继续完善文件共享模块，提供查看当前已共享的文件与被共享文件
  - 完善后端接口，完善前端页面，完成基本功能
  - 修改设计缺陷，删除文件目标用户ID,删除用户策略ID，删除用户策略表，使用共享文件表
- 提供当前文件拥有者进行文件共享
  - 创建共享策略
  - 获取当前所有用户列表
  - 初步完成用户共享文件需求
- 发现新的问题
  - jwt频繁验证签名，每次请求都要进行再次查看用户登录态，比较耗费性能，后续考虑进行优化
  - 前端请求频繁导致后端无法及时响应，后续考虑引入缓存进行优化，避免频繁请求
- 新的计划
  - 明天完成查看当前用户共享病历与共享给当前用户病历，以及相关优化内容
  - 打包成apk，发给老师进行检查
  - 着手新技术，撰写毕业论文

## TODO
**接口开发任务**

- [ ] 用户模块
  - [x] 用户表设计
  - [x] 用户注册
  - [x] 用户登录
    - [x] 使用jwt记录登录态
    - [x] 验证token作为登录条件
  - [ ] 用户密钥更新
- [ ] 策略模块
  - [x] 共享策略设计
    - [x] 使用Json字符串存储共享规则
    - [x] 关联用户与共享策略
  - [x] 共享策略创建
  - [x] 查看当前用户的共享策略
  - [x] 共享策略修改
- [ ] 病例模块
  - [x] 病例表设计
  - [ ] 病例初次共享
    - [ ] 使用用户公钥进行加密
  - [ ] 病例二次共享
    - [ ] 创建新策略，符合当前用户创建的共享策略与病例原有者的共享策略
    - [ ] 使用当前用户公钥进行二次加密