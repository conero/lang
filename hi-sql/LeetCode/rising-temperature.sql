/**
 *  @date 2018年11月3日 星期六
 *  @description #197 上升的温度
 *  @link https://leetcode-cn.com/problems/rising-temperature/
**/


/*
    给定一个 Weather 表，编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 Id。

    +---------+------------------+------------------+
    | Id(INT) | RecordDate(DATE) | Temperature(INT) |
    +---------+------------------+------------------+
    |       1 |       2015-01-01 |               10 |
    |       2 |       2015-01-02 |               25 |
    |       3 |       2015-01-03 |               20 |
    |       4 |       2015-01-04 |               30 |
    +---------+------------------+------------------+
    例如，根据上述给定的 Weather 表格，返回如下 Id:

    +----+
    | Id |
    +----+
    |  2 |
    |  4 |
    +----+

*/


-- SQL 架构
    Create table If Not Exists Weather (Id int, RecordDate date, Temperature int);
    Truncate table Weather;
    insert into Weather (Id, RecordDate, Temperature) values ('1', '2015-01-01', '10');
    insert into Weather (Id, RecordDate, Temperature) values ('2', '2015-01-02', '25');
    insert into Weather (Id, RecordDate, Temperature) values ('3', '2015-01-03', '20');
    insert into Weather (Id, RecordDate, Temperature) values ('4', '2015-01-04', '30');

-- MySQL
    //@TODO    
