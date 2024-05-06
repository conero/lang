/**
 *  @date 2018年11月2日 星期五
 *  @description 编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary） 。
 *  @link https://leetcode-cn.com/problems/second-highest-salary/
**/

/*
    编写一个 SQL 查询，获取 Employee 表中第二高的薪水（Salary） 。

    +----+--------+
    | Id | Salary |
    +----+--------+
    | 1  | 100    |
    | 2  | 200    |
    | 3  | 300    |
    +----+--------+
    例如上述 Employee 表，SQL查询应该返回 200 作为第二高的薪水。如果不存在第二高的薪水，那么查询应返回 null。

    +---------------------+
    | SecondHighestSalary |
    +---------------------+
    | 200                 |
    +---------------------+
*/


-- SQL 架构
    Create table If Not Exists Employee (Id int, Salary int);
    Truncate table Employee;
    insert into Employee (Id, Salary) values ('1', '100');
    insert into Employee (Id, Salary) values ('2', '200');
    insert into Employee (Id, Salary) values ('3', '300');

-- MySQL    
    -- 解决
        -- raw
        select 
            ifnull((
                select distinct(Salary) as SecondHighestSalary 
                from Employee 
                order by Salary desc
                limit 1,1
            ), null) as SecondHighestSalary
            from dual
            ;
        -- 优化
        select 
            (
                select distinct(Salary)
                from Employee 
                order by Salary desc
                limit 1,1
            ) as SecondHighestSalary
            from dual
            ;


-- oracle
        select (
            with t1 as (
                select distinct Salary from Employee order by Salary desc
            ) 
            select Salary from (select Salary, rownum as rn from t1 where rownum<=2) where rn>=2
        ) 
        as SecondHighestSalary from dual
        ;