<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-15
 * Time: 09:46
 */

namespace frontend\controllers;


class IndexController extends BasicController
{
    public $layout = 'main.tpl';

    public function actionIndex()
    {
        return $this->render('index.tpl');
    }
}
