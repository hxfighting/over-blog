<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-20
 * Time: 14:06
 */

namespace frontend\controllers;


use yii\log\Target;

class SqlController extends Target
{
    public function export(){
        array_pop($this->messages); //去掉最后一个消息，因为这个与SQL无关，只是底层自动额外追加的一些请求和运行时相关信息

        $sqlList = [];
        foreach($this->messages as $message){
            $sqlList[] = $message[0];
        }

        $listContent = implode(PHP_EOL, $sqlList);
        $sqlNums = count($sqlList);
        $logContent = <<<EOL
执行的SQL次数：$sqlNums

执行过的SQL语句：
$listContent
EOL;
        $logFile = \Yii::getAlias('@frontend/runtime/logs/sql.log');
        file_put_contents($logFile, $logContent);
    }
}
