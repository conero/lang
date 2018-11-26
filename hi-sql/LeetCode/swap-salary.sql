/**
 *  @date 2018年11月26日 星期一
 *  @description #627 交换工资
 *  @link https://leetcode-cn.com/problems/swap-salary/
**/



/*
    给定一个 salary表，如下所示，有m=男性 和 f=女性的值 。交换所有的 f 和 m 值(例如，将所有 f 值更改为 m，反之亦然)。要求使用一个更新查询，并且没有中间临时表。

    例如:

    | id | name | sex | salary |
    |----|------|-----|--------|
    | 1  | A    | m   | 2500   |
    | 2  | B    | f   | 1500   |
    | 3  | C    | m   | 5500   |
    | 4  | D    | f   | 500    |
    运行你所编写的查询语句之后，将会得到以下表:

    | id | name | sex | salary |
    |----|------|-----|--------|
    | 1  | A    | f   | 2500   |
    | 2  | B    | m   | 1500   |
    | 3  | C    | f   | 5500   |
    | 4  | D    | m   | 500    |
*/


-- SQL 架构
    create table if not exists salary(id int, name varchar(100), sex char(1), salary int);
    Truncate table salary;
    insert into salary (id, name, sex, salary) values ('1', 'A', 'm', '2500');
    insert into salary (id, name, sex, salary) values ('2', 'B', 'f', '1500');
    insert into salary (id, name, sex, salary) values ('3', 'C', 'm', '5500');
    insert into salary (id, name, sex, salary) values ('4', 'D', 'f', '500');


-- MySQL
    update salary set sex=if(sex='m', 'f', 'm')
    ;

