<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-27
 * Time: 09:41
 */

namespace backend\controllers;


use app\models\Chat;
use backend\exception\ValidateException;

class ChatController extends BasicController
{
    public $enableCsrfValidation = false;
    private $chat;

    public function init()
    {
        parent::init();
        $this->chat = new Chat();
    }

    /**
     * 获取说说列表
     * Date: 2019-02-27 10:09
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionChatList()
    {
        $this->basicValidate($this->chat, 'chatList');
        $list = $this->chat->find()
            ->offset((($this->request_data)['pageNum'] - 1) * ($this->request_data)['pageSize'])
            ->limit(($this->request_data)['pageSize'])
            ->orderBy('created_at DESC')
            ->all();
        return !empty($list) ? $this->success('获取说说列表成功！', $list)
            : $this->error('暂无说说列表数据！');
    }

    /**
     * 添加说说
     * Date: 2019-02-27 10:09
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionChatAdd()
    {
        $chat = $this->basicValidate($this->chat, 'addChat');
        $res = $chat->save(false);
        return $res ? $this->success('添加标签成功！')
            : $this->error('添加标签失败，请稍后再试！');
    }

    /**
     * 更新说说
     * Date: 2019-02-27 10:09
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionChatUpdate()
    {
        $this->basicValidate($this->chat, 'updateChat');
        $exist_chat = $this->chat->findOne(($this->request_data)['id']);
        $exist_chat->content = ($this->request_data)['content'];
        $exist_chat->is_show = ($this->request_data)['is_show'];
        $res = $exist_chat->save(false, ['content', 'is_show', 'updated_at']);
        return $res ? $this->success('修改标签成功！')
            : $this->error('修改标签失败，请稍后再试！');
    }

    /**
     * 删除说说
     * Date: 2019-02-27 10:10
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionDelChat()
    {
        $this->basicValidate($this->chat, 'chatDelete');
        $exist_chat = $this->chat->findOne(($this->request_data)['id']);
        $res = $exist_chat->delete();
        return $res ? $this->success('删除标签成功！')
            : $this->error('删除标签失败，请稍后再试！');
    }
}
