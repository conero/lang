/*==============================================================*/
/* DBMS name:      MySQL 5.7 初始化数据                          */
/* Created on:     2019/8/15 16:05:34                           */
/*==============================================================*/


-- /-------------------- 【系统常量】
-- 基本状态
INSERT INTO config
   (`key`, descrip, value, text_value, sub_mk)
   VALUES ('status', '通用状态描述', '', '', '1')
;

set @pid := @@identity;
INSERT INTO config_subset
   (pid, subval, subdesc)
   VALUES
   (@pid, '90', '有效')
   ,(@pid, '40', '已禁数据无效')
   ,(@pid, '10', '待处理，初始化状态')
;




-- /-------------------- 【系统管理员账户】
INSERT INTO `user` (id, account, password, status) values
   (0, 'dev', '12345', '90')
;
set @user_dev := @@identity;

-- /-------------------- 【角色】
INSERT INTO `role` (name, code, status)
   VALUES
      ('开发', 'dev', '90')
;

set @role_dev := @@identity;
INSERT INTO userole (uid, roleid)
   VALUES
   (@user_dev, @role_dev)
;

