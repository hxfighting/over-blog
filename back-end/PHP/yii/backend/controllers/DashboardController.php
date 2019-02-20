<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-18
 * Time: 18:01
 */

namespace backend\controllers;


class DashboardController extends BasicController
{
    public function actionIndex()
    {
        return ['data'=>'123','code'=>200,'msg'=>'哈哈哈'];
    }
}
