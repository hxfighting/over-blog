<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-19
 * Time: 13:29
 */

namespace backend\controllers;


use app\models\WebError;
use yii\filters\VerbFilter;
use yii\helpers\ArrayHelper;
use yii\web\MethodNotAllowedHttpException;
use yii\web\NotFoundHttpException;

class ErrorController extends BasicController
{
    public $enableCsrfValidation = false;

    public function behaviors()
    {
        return ArrayHelper::merge([
            'verbs' => [
                'class' => VerbFilter::className(),
                'actions' => [
                    'log-error'  => ['post'],
                ],
            ],
        ], parent::behaviors());
    }

    public function actionIndex()
    {
        $exception = \Yii::$app->errorHandler->exception;
        if($exception instanceof NotFoundHttpException){
            return $this->error('路由不存在！');
        }
        if($exception instanceof MethodNotAllowedHttpException){
            return $this->error('请求方式错误！');
        }
        return $this->error($exception->getMessage());
    }

    /**
     * 记录错误
     * Date: 2019-02-19 14:18
     * @return \yii\web\Response
     */
    public function actionLogError()
    {
        $data = $this->post();
        $error = new WebError();
        $error->attributes = $data;
        $res = $error->save();
        return $res?$this->success('记录错误成功！')
            :$this->error(current($error->firstErrors));
    }
}
