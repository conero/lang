<?php
/**
 * 2017年10月12日 星期四
 * Sqler 测试用例
 */
use conero\learn\Sqler;
require_once "./ConeroTestApp.php";
class TestCase{
    public function __construct(){
        // $this->baseTest();
        $this->selectBaseTest();
    }
    // 基本测试
    public function baseTest(){
        // 最基本的对象
        $dataOrc = (new Sqler('ocltb'))
            ->data([
                'c1' => 'Joshua',
                '"mtime" = sysdate'
            ])
            ->where(["id"=> 5])
            ->update()
            ;
        print_r($dataOrc);

        // 单例程序脚本测试
        $mysql = Sqler::table('mysqltb', 'mysql');
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
    public function selectBaseTest(){
        // 单表查询
        $user = (new Sqler('sys_user'))
            ->field(['name', 'gender', 'birthday', 'contry'])
            ->where(['name' => 'test'])
            ->select()
        ;
        $userMutTbs = Sqler::table(['sys_user', 'a'])
            ->join(['login', 'b'], 'a.uid = b.uid')
            ->join('family c', 'a.family_id=c.id', 'left')
            ->select()
            ;
        /*
        $userMutTbs2 = Sqler::table(['sys_user', 'a'])
            ->field(['a.name', 'b.login_date', 'b.login_count', 'c.family_name'])
            ->join(['login', 'b'], 'a.uid = b.uid')
            ->join('family c', 'a.family_id=c.id', 'left')
            ->select()
        ;
        */
        print_r([
            /*$user
            , */
            $userMutTbs
            //, $userMutTbs2
        ]);
    }
}
new TestCase;
 