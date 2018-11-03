/**
 *  @date 2018年11月3日 星期六
 *  @description #196 删除重复的电子邮件
 *  @link https://leetcode-cn.com/problems/delete-duplicate-emails/
**/


/*
    编写一个 SQL 查询，来删除 Person 表中所有重复的电子邮箱，重复的邮箱里只保留 Id 最小 的那个。

    +----+------------------+
    | Id | Email            |
    +----+------------------+
    | 1  | john@example.com |
    | 2  | bob@example.com  |
    | 3  | john@example.com |
    +----+------------------+
    Id 是这个表的主键。
    例如，在运行你的查询语句之后，上面的 Person 表应返回以下几行:

    +----+------------------+
    | Id | Email            |
    +----+------------------+
    | 1  | john@example.com |
    | 2  | bob@example.com  |
    +----+------------------+

*/

//@TODO  失败
-- MySQL  
    -- / 原型
    select Email, MaxId from (select Email, count(1) ctt, max(Id) MaxId from Person group by Email) tt
        where tt.ctt > 1
    ;

    -- 执行代码
    delete from Person a where 
    exists (select 1 from (select count(1) ctt from Person where Email = a.Email group by Email) tt where tt.ctt > 1)
    and 
    a.Id <> (
        select MaxId from (select count(1) ctt, max(Id) MaxId from Person where Email = a.Email group by Email) tt
        where tt.ctt > 1
    ) 
    ;

    delete from Person a where 
    exists (select 1 from (select count(Email) ctt from Person where Email = a.Email) tt where tt.ctt > 1)
    and 
    a.Id <> (
        select MaxId from (select count(1) ctt, max(Id) MaxId from Person where Email = a.Email group by Email) tt
        where tt.ctt > 1
    ) 
    ;

    delete from Person a where 
    exists (
        select 1 from (select count(Email) as ctt, max(Id) as maxId from Person where Email=a.Email) tt where tt.ctt > 1 and tt.maxId <> a.Id
    ) 
    ;


