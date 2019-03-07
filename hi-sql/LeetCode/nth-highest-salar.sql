/**
 *  @date 2018年12月28日 星期五
 *  @description 177. 第N高的薪水
 *  @link https://leetcode-cn.com/problems/nth-highest-salary/
**/


/*
    编写一个 SQL 查询，获取 Employee 表中第 n 高的薪水（Salary）。

    +----+--------+
    | Id | Salary |
    +----+--------+
    | 1  | 100    |
    | 2  | 200    |
    | 3  | 300    |
    +----+--------+
    例如上述 Employee 表，n = 2 时，应返回第二高的薪水 200。如果不存在第 n 高的薪水，那么查询应返回 null。

    +------------------------+
    | getNthHighestSalary(2) |
    +------------------------+
    | 200                    |
    +------------------------+
*/



-- MySQL
    CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
        BEGIN
        RETURN (
            # Write your MySQL query statement below.
            -- limit 不可做运算
            -- select distinct Salary from Employee order by Salary desc limit N, N+1

            -- 子查询使用 limit
            -- select t.Salary from (select distinct Salary from Employee order by Salary desc limit N) t order by t.Salary limit 1

            -- 子查询使用 min
            select min(t.Salary) from (select distinct Salary, found_rows() as ctt from Employee order by Salary desc limit N) t where t.ctt = N

        );
    END

-- Oracle
    CREATE FUNCTION getNthHighestSalary(N IN NUMBER) RETURN NUMBER IS
    result NUMBER;
    BEGIN
        result := 0;
        /* Write your PL/SQL query statement below */
        SELECT COUNT(1) INTO result FROM (
            SELECT distinct Salary FROM Employee
        );        
        IF N > result THEN
            RETURN NULL;
        END IF;

        -- @todo
        RETURN result;
    END;    



