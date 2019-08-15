/*==============================================================*/
/* DBMS name:      MySQL 5.7                                    */
/* Created on:     2019/8/15 17:30:52                           */
/*==============================================================*/


drop table if exists auth;

drop table if exists config;

drop table if exists config_subset;

drop table if exists role;

drop table if exists router;

drop table if exists router_group;

drop table if exists `user`;

drop table if exists userole;

/*==============================================================*/
/* Table: auth                                                  */
/*==============================================================*/
create table auth
(
   id                   int not null auto_increment comment '数据主键',
   refid                int not null comment '用户/角色',
   is_role              char(1) not null default '1' comment '为角色授权',
   access_rid           int comment '可访问路由键',
   access_grid          int comment '可访问分组路由键',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '编辑时间',
   muid                 int not null default 0 comment '编辑者',
   primary key (id)
);

alter table auth comment 'auth - 授权信息表';

/*==============================================================*/
/* Table: config                                                */
/*==============================================================*/
create table config
(
   id                   int not null auto_increment comment '主键',
   `key`                varchar(80) not null comment '键名称',
   descrip              varchar(100) comment '文本描述',
   value                varchar(200) comment '值',
   text_value           text comment '可选长文本值',
   sub_mk               char(1) not null default '0' comment '存在子集表',
   muid                 int not null default 0 comment '操作者id',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '创建时间',
   primary key (id)
);

alter table config comment 'config - 系统配置';

/*==============================================================*/
/* Table: config_subset                                         */
/*==============================================================*/
create table config_subset
(
   id                   int not null auto_increment comment '主键',
   pid                  int not null comment '父主键关联',
   subval               varchar(200) comment '键值',
   subdesc              varchar(100) comment '文本描述',
   muid                 int not null default 0 comment '操作者',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '创建时间',
   primary key (id)
);

alter table config_subset comment 'config_subset 配置子表';

/*==============================================================*/
/* Table: role                                                  */
/*==============================================================*/
create table role
(
   id                   int not null auto_increment comment '数据id',
   name                 varchar(35) not null comment '角色名称',
   code                 varchar(20) not null comment '角色代码; 用于不同系统中数据同步，有系统id不一致',
   status               char(2) not null comment '状态; 用于标记角色的状态，可进行禁用等操作',
   muid                 int not null default 0 comment '操作者',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '创建时间戳',
   primary key (id)
);

alter table role comment 'role - 角色表';

/*==============================================================*/
/* Table: router                                                */
/*==============================================================*/
create table router
(
   id                   int not null auto_increment comment '主键',
   method               varchar(30) not null comment '方法',
   url                  varchar(100) not null comment '地址',
   name                 varchar(80) not null comment '路由名称',
   remark               varchar(200) comment '描述',
   rgid                 int comment '路由分组键',
   muid                 int not null default 0 comment '编辑者',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '维护时间戳',
   primary key (id)
);

alter table router comment '鉴于目前系统通常都是路由绑定，随以此来记录系统模块';

/*==============================================================*/
/* Table: router_group                                          */
/*==============================================================*/
create table router_group
(
   id                   int not null auto_increment comment '主键id',
   name                 varchar(80) not null comment '分组名称',
   descrip              varchar(200) comment '分组描述',
   remark               varchar(300) comment '备注信息',
   muid                 int not null default 0 comment '编辑者',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '维护时间戳',
   primary key (id)
);

alter table router_group comment 'router_group 路由分组';

/*==============================================================*/
/* Table: `user`                                                */
/*==============================================================*/
create table `user`
(
   id                   int not null auto_increment comment '数据id',
   account              varchar(45) not null comment '账户名称',
   password             varchar(50) not null comment '账户密码',
   status               char(2) not null comment '状态',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '注册时间',
   ip                   varchar(20) comment '注册ip',
   last_mtime           datetime comment '最近登录时间',
   login_count          int not null default 0 comment '累计登录次数',
   primary key (id),
   unique key UQ_account (account)
);

alter table `user` comment 'user - 用户信息表';

/*==============================================================*/
/* Table: userole                                               */
/*==============================================================*/
create table userole
(
   uid                  int not null comment '用户id',
   roleid               int not null comment '角色id',
   mtime                datetime not null default CURRENT_TIMESTAMP comment '维护时间戳',
   primary key (uid, roleid)
);

alter table userole comment 'userole - 用户角色关联';

alter table config_subset add constraint FK_config_subset__pid foreign key (pid)
      references config (id) on delete restrict on update restrict;

alter table router add constraint FK_router__rgid foreign key (rgid)
      references router_group (id) on delete restrict on update restrict;

alter table userole add constraint FK_userole__roleid foreign key (roleid)
      references role (id) on delete restrict on update restrict;

alter table userole add constraint FK_userole__uid foreign key (uid)
      references `user` (id) on delete restrict on update restrict;

