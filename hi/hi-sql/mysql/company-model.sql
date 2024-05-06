-- 2021年3月23日 星期二
-- 公司模型数表设计：公司基本信息、公司排名

-- 删除名称
drop table if exists comapny_rank_relation;
drop table if exists comapny_rank;
drop table if exists comapny;


-- 公司名称
create table if not exists comapny(
	`id` int NOT NULL AUTO_INCREMENT COMMENT '数据主键',
	`name` varchar(200) not null comment '公司名称',
	`en_name` varchar(200) not null comment '公司英文名称',
	`main_business` varchar(200) comment '主要业务与品牌',
	`contry` varchar(100) comment '所在国家',
	`province` varchar(100) comment '所在省份',
	`mtime` datetime NOT NULL default current_timestamp COMMENT '更新时间',
	`ctime` datetime NOT NULL default current_timestamp COMMENT '创建时间',
	PRIMARY KEY (`id`)
)
	DEFAULT CHARSET=utf8 COMMENT='公司库存表'
;


-- 公示排名级别
create table if not exists comapny_rank(
	`id` int NOT NULL AUTO_INCREMENT COMMENT '数据主键',
	`name` varchar(200) not null comment '排行名称',
	`income` numeric(10, 2) not null default 0 comment '营收入',
	`profit` numeric(10, 2) not null default 0 comment '利润',
	`unit` varchar(20) comment '金额单位',
	`time_year` year not null default 0 comment '年度',
	`rank` int not null default 0 comment '排行',
	`mtime` datetime NOT NULL default current_timestamp COMMENT '更新时间',
	`ctime` datetime NOT NULL default current_timestamp COMMENT '创建时间',
	PRIMARY KEY (`id`)
)
	DEFAULT CHARSET=utf8 COMMENT='公司排名记录'
;

-- 公示排名级别关联
create table if not exists comapny_rank_relation(
	`id` int NOT NULL AUTO_INCREMENT COMMENT '数据主键',
	`rank` int not null default 0 comment '排行',
	`mtime` datetime NOT NULL default current_timestamp COMMENT '更新时间',
	`ctime` datetime NOT NULL default current_timestamp COMMENT '创建时间',
    `com_id` int NOT NULL COMMENT '公司主键',
    `rank_id` int NOT NULL COMMENT '排名关系',
	PRIMARY KEY (`id`)
)
	DEFAULT CHARSET=utf8 COMMENT='公司排名记录'
;
alter table comapny_rank_relation add constraint FK_comapny_rank__com_id foreign key (com_id)
      references comapny (id) on delete restrict on update restrict;
alter table comapny_rank_relation add constraint FK_comapny_rank_relation__rank_id foreign key (rank_id)
      references comapny_rank (id) on delete restrict on update restrict;