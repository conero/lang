/**
 *  @date 2018年11月2日 星期五
 *  @description 查找重复的电子邮箱
 *  @link https://leetcode-cn.com/problems/duplicate-emails/
**/

/*
 编写一个 SQL 查询，查找 Person 表中所有重复的电子邮箱。

    示例：

    +----+---------+
    | Id | Email   |
    +----+---------+
    | 1  | a@b.com |
    | 2  | c@d.com |
    | 3  | a@b.com |
    +----+---------+
    根据以上输入，你的查询应返回以下结果：

    +---------+
    | Email   |
    +---------+
    | a@b.com |
    +---------+
    说明：所有电子邮箱都是小写字母。
*/


-- SQL 架构
    Create table If Not Exists Person (Id int, Email varchar(255));
    Truncate table Person;
    insert into Person (Id, Email) values ('1', 'a@b.com');
    insert into Person (Id, Email) values ('2', 'c@d.com');
    insert into Person (Id, Email) values ('3', 'a@b.com');

-- mysql/oracle
    select tt.Email from (select Email, count(1) as ctt from Person group by Email) tt where ctt>1
