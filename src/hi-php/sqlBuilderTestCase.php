<?php
/**
 * 2017年10月12日 星期四
 * sqlBuilder 测试用例
 */

require_once "./sqlBuilder.php";
class TestCase{
    public function __construct(){
        $this->baseTest();
    }
    // 基本测试
    public function baseTest(){
        // 最基本的对象
        $dataOrc = (new sqlBuilder('ocltb'))
            ->data([
                'c1' => 'Joshua',
                '"mtime" = sysdate'
            ])
            ->where(["id"=> 5])
            ->update()
            ;
        print_r($dataOrc);

        // 单例程序脚本测试
        $mysql = sqlBuilder::table('mysqltb', 'mysql');
        $dataMySql = $mysql
            ->data([
                'c1' => 'Joshua',
                'test_date' => '20171012',
                '"mtime" = sysdate'
            ])
            ->where(["id"=> 5])
            ->update()
        ;
        print_r([
            'mysql' => $dataMySql,
            'sql'   => $mysql->getRawSql($dataMySql)
        ]);
    }
}
new TestCase;
 