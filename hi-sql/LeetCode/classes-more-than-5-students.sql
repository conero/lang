/**
 *  @date 2018年11月26日 星期一
 *  @description #596 超过5名学生的课
 *  @link https://leetcode-cn.com/problems/classes-more-than-5-students/
**/



/*
    有一个courses 表 ，有: student (学生) 和 class (课程)。

    请列出所有超过或等于5名学生的课。

    例如,表:

    +---------+------------+
    | student | class      |
    +---------+------------+
    | A       | Math       |
    | B       | English    |
    | C       | Math       |
    | D       | Biology    |
    | E       | Math       |
    | F       | Computer   |
    | G       | Math       |
    | H       | Math       |
    | I       | Math       |
    +---------+------------+
    应该输出:

    +---------+
    | class   |
    +---------+
    | Math    |
    +---------+
*/


-- SQL 架构
    Create table If Not Exists courses (student varchar(255), class varchar(255));
    Truncate table courses;
    insert into courses (student, class) values ('A', 'Math');
    insert into courses (student, class) values ('B', 'English');
    insert into courses (student, class) values ('C', 'Math');
    insert into courses (student, class) values ('D', 'Biology');
    insert into courses (student, class) values ('E', 'Math');
    insert into courses (student, class) values ('F', 'Computer');
    insert into courses (student, class) values ('G', 'Math');
    insert into courses (student, class) values ('H', 'Math');
    insert into courses (student, class) values ('I', 'Math');


-- MySQL
    -- 两层子查询
    select tt.class from (
        select t.class, count(1) ct from(
            select student, class from courses group by student, class
        ) t group by t.class
    ) tt where tt.ct>=5
    ;

