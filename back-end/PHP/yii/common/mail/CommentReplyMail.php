<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-08
 * Time: 11:24
 */

use yii\helpers\Html;
?>
<p><?= $reply_name ?> 你好</p>

<p><?= $comment_name?>在文章<< <?=Html::encode($title)?> >>中回复了你:</p>

<p><?= $reply_content?></p>

<p><?= Html::a(Html::encode($title),$url) ?></p>

<p><?= Html::a(Yii::$app->params['blog_name'],Yii::$app->params['blog_index_url']) ?></p>
