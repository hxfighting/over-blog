<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-06
 * Time: 09:43
 */

namespace backend\job;


use yii\base\BaseObject;
use yii\queue\JobInterface;

class ReplyMailJob extends BaseObject implements JobInterface
{
    public $reply_content;

    public $name;

    public $email;

    /**
     * 发送邮件队列
     * Date: 2019-03-06 10:30
     * @param \yii\queue\Queue $queue
     */
    public function execute($queue)
    {
        try
        {
            $subject = \Yii::$app->params['blog_name'] . '留言回复';
            $reply_content = $this->reply_content ? $this->reply_content : '非常感谢你的留言,我会尽快回复你的.';
            \Yii::$app->mailer->compose('ReplyUserMail', ['name' => $this->name, 'reply_content' => $reply_content])
                ->setTo($this->email)
                ->setSubject($subject)
                ->send();
        } catch (\Exception $e)
        {
            \Yii::$app->log->getLogger()->log($e->getMessage(),1);
        }
    }
}
