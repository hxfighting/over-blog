<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-19
 * Time: 13:29
 */

namespace backend\controllers;


use app\models\WebError;
use backend\exception\ValidateException;
use yii\filters\VerbFilter;
use yii\helpers\ArrayHelper;
use yii\web\MethodNotAllowedHttpException;
use yii\web\NotFoundHttpException;

class ErrorController extends BasicController
{
    public $enableCsrfValidation = false;
    private $error;

    public function init()
    {
        parent::init();
        $this->error = new WebError();
    }

    /**
     * 统一错误处理
     * Date: 2019-03-12 10:15
     * @return \yii\web\Response
     */
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
        if ($exception instanceof ValidateException)
        {
            return $this->error($exception->getMessage());
        }
        \Yii::$app->log->getLogger()->log($exception->getMessage(),1);
        return $this->error('服务器错误！');
    }

    /**
     * 记录错误
     * Date: 2019-02-19 14:18
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionLogError()
    {
        $error = $this->basicValidate($this->error, 'logError');
        $res = $error->save(false);
        return $res ? $this->success('记录错误成功！')
            : $this->error('记录错误失败，请稍后再试！');
    }

    /**
     * 删除错误信息
     * Date: 2019-02-21 16:06
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionDelError()
    {
        $this->basicValidate($this->error, 'delError');
        $data = $this->request_data;
        $exist_error = $this->error->findAll($data['ids']);
        $res = $exist_error->delete();
        return $res ? $this->success('错误信息删除成功！')
            : $this->error('错误信息删除失败，请稍后再试！');
    }
}
