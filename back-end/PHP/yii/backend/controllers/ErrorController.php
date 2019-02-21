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

    public function actionIndex()
    {
        $exception = \Yii::$app->errorHandler->exception;
        if ($exception instanceof NotFoundHttpException)
        {
            return $this->error('路由不存在！');
        }
        if ($exception instanceof MethodNotAllowedHttpException)
        {
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
        $error->scenario = 'logError';
        $error->attributes = $data;
        $res = $error->save();
        return $res ? $this->success('记录错误成功！')
            : $this->error(current($error->firstErrors));
    }

    /**
     * 删除错误信息
     * Date: 2019-02-21 16:06
     * @return \yii\web\Response
     */
    public function actionDelError()
    {
        $data = $this->post();
        $error = new WebError();
        $error->scenario = 'delError';
        $error->attributes = $data;
        $res = $error->find()->where(['in', 'id', $data['ids']])->delete();
        return $res ? $this->success('错误信息删除成功！')
            : $this->error('错误信息删除失败，请稍后再试！');
    }
}
