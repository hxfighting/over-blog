<?php
use yii\helpers\Html;
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-06
 * Time: 09:59
 */
?>

<p><?= $name ?> 你好</p>

<p><?= $reply_content ?></p>

<p><?= Html::a(Yii::$app->params['blog_name'],Yii::$app->params['blog_index_url']) ?></p>
