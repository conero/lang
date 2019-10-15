-- 副本数据移动到新表中
-- Joshua Conero
-- 2019年10月15日 星期二
-- 备份表: cp_category_191015, cp_article_191015, cp_article_data_191015

-- 创建临时存储过程并执行(如： _jc_tmp_untitle_sqlblock)
-- 删除过程过程，词语句可以去除点
drop procedure if exists _jc_tmp_untitle_sqlblock;
drop procedure if exists _jc_tmp_untitle_move_by_catid;
-- 备份表创建
create table if not exists cp_category_191015 as select * from fg_category where 1=0;
create table if not exists cp_article_191015 as select * from fg_article where 1=0;
create table if not exists cp_article_data_191015 as select * from fg_article_data where 1=0;

-- 多级递归目录移动
delimiter //
create procedure _jc_tmp_untitle_move_by_catid(v_catid int)
begin
    declare not_more_record int default false;
    declare v_bm_catid int;

    -- cursor 定义语法
    declare bumen_cs cursor for select `catid` from fg_category where parentid=v_catid and `catid`<>v_catid;
	declare continue handler for not found set not_more_record= true;

	-- 修复 eror: 1456 Recursive limit 0 (as set by the max_sp_recursion_depth variable) was exceeded for routine sp_rebuild_booktype.
	-- link: https://blog.csdn.net/qq_40741855/article/details/89179029
	SET @@max_sp_recursion_depth = 20;
	-- 导文章
    insert into cp_article_191015 select * from fg_article where catid=v_catid;
    insert into cp_article_data_191015 select dd.* from fg_article_data dd
        inner join fg_article fat on dd.id=fat.id and fat.catid=v_catid
    ;
    -- 数据参数
    delete dd from fg_article_data dd
        inner join fg_article fat on dd.id=fat.id and fat.catid=v_catid
    ;
    delete from fg_article where catid=v_catid;

    -- 结果集循环
    open bumen_cs;
        -- 循环语句： loop, repeat, while
        while not_more_record<>true do
            fetch bumen_cs into v_bm_catid;
       		if v_bm_catid > 0 and  v_bm_catid <> v_catid then
            	call _jc_tmp_untitle_move_by_catid(v_bm_catid);
            end if;
        end while;
    -- 关闭游标
    close bumen_cs;

    -- 菜单导入
    insert into cp_category_191015 select * from fg_category where parentid=v_catid;
    delete from fg_category where parentid=v_catid;
    -- 删除当前的菜单
    insert into cp_category_191015 select * from fg_category where `catid`=v_catid;
    delete from fg_category where `catid`=v_catid;
end;
//
delimiter ;


-- 创建
-- 修改分割符号
-- 创建
-- 修改分割符号
delimiter //
create procedure _jc_tmp_untitle_sqlblock()
begin
	declare not_more_record int default false;
    declare v_bm_catid int;
	-- cursor 定义语法
    declare bumen_cs cursor for select `catid` from fg_category where catname='办事指南-旧栏目备份';
	declare continue handler for not found set not_more_record= true;

    open bumen_cs;
        -- fetch 获取游标中的参数
        fetch bumen_cs into v_bm_catid;
        call _jc_tmp_untitle_move_by_catid(v_bm_catid);

    -- 关闭游标
    close bumen_cs;
end;
//

-- 恢复分割符号
delimiter ;

-- 执行匿名块
call _jc_tmp_untitle_sqlblock();
-- 删除匿名块
drop procedure if exists _jc_tmp_untitle_sqlblock;
drop procedure if exists _jc_tmp_untitle_move_by_catid;