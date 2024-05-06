/**
 *  @date 2018年11月26日 星期一
 *  @description #595 大的国家
 *  @link https://leetcode-cn.com/problems/big-countries/
**/



/*
    这里有张 World 表

    +-----------------+------------+------------+--------------+---------------+
    | name            | continent  | area       | population   | gdp           |
    +-----------------+------------+------------+--------------+---------------+
    | Afghanistan     | Asia       | 652230     | 25500100     | 20343000      |
    | Albania         | Europe     | 28748      | 2831741      | 12960000      |
    | Algeria         | Africa     | 2381741    | 37100000     | 188681000     |
    | Andorra         | Europe     | 468        | 78115        | 3712000       |
    | Angola          | Africa     | 1246700    | 20609294     | 100990000     |
    +-----------------+------------+------------+--------------+---------------+
    如果一个国家的面积超过300万平方公里，或者人口超过2500万，那么这个国家就是大国家。

    编写一个SQL查询，输出表中所有大国家的名称、人口和面积。

    例如，根据上表，我们应该输出:

    +--------------+-------------+--------------+
    | name         | population  | area         |
    +--------------+-------------+--------------+
    | Afghanistan  | 25500100    | 652230       |
    | Algeria      | 37100000    | 2381741      |
    +--------------+-------------+--------------+
*/


-- SQL 架构
    Create table If Not Exists World (name varchar(255), continent varchar(255), area int, population int, gdp int);
    Truncate table World;
    insert into World (name, continent, area, population, gdp) values ('Afghanistan', 'Asia', '652230', '25500100', '20343000000');
    insert into World (name, continent, area, population, gdp) values ('Albania', 'Europe', '28748', '2831741', '12960000000');
    insert into World (name, continent, area, population, gdp) values ('Algeria', 'Africa', '2381741', '37100000', '188681000000');
    insert into World (name, continent, area, population, gdp) values ('Andorra', 'Europe', '468', '78115', '3712000000');
    insert into World (name, continent, area, population, gdp) values ('Angola', 'Africa', '1246700', '20609294', '100990000000');


-- MySQL
    -- 300 0000
    select name, population, area from World where area > 3000000 or population > 25000000;
    select name, population, area from World where area > (300 * 10000) or population > (2500 * 10000);

