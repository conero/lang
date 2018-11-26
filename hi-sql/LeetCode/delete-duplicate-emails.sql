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

-- MySQL  
    -- / 原型
    -- MySQL 语法报错
    delete from Person ps where 
        exists (
            select 1 from Person p1 where ps.Email=p1.Email and p1.Id < ps.Id
        )
    ;
    -- 看了官网答案后
    delete ps from Person ps, Person p1
        where ps.Email=p1.Email and ps.Id > p1.Id
    ;
