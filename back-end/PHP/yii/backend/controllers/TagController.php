<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-22
 * Time: 09:18
 */

namespace backend\controllers;


use app\models\Tag;
use backend\exception\ValidateException;

class TagController extends BasicController
{
    public $enableCsrfValidation = false;
    private $tag;

    public function init()
    {
        parent::init();
        $this->tag = new Tag();
    }

    /**
     * 获取标签列表
     * Date: 2019-02-22 10:10
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionTagList()
    {
        $this->basicValidate($this->tag, 'tagList');
        $data = $this->request_data;
        $query = $this->tag->find();
        $total = $query->count();
        $list = $query
            ->offset(($data['pageNum'] - 1) * $data['pageSize'])
            ->limit($data['pageSize'])
            ->orderBy('created_at DESC')
            ->all();
        $data = compact('list', 'total');
        return !empty($list) ? $this->success('获取标签列表成功！', $data)
            : $this->error('暂无标签列表数据！');
    }

    /**
     * 添加标签
     * Date: 2019-02-22 10:15
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionAddTag()
    {
        $tag = $this->basicValidate($this->tag, 'tagAdd');
        $res = $tag->save(false);
        return $res ? $this->success('添加标签成功！')
            : $this->error('添加标签失败，请稍后再试！');
    }

    /**
     * 修改标签
     * Date: 2019-02-22 10:21
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionUpdateTag()
    {
        $this->basicValidate($this->tag, 'tagUpdate');
        $data = $this->request_data;
        $exist_tag = $this->tag->findOne($data['id']);
        $exist_tag->name = $data['name'];
        $res = $exist_tag->save(false, ['name', 'updated_at']);
        return $res ? $this->success('修改标签成功！')
            : $this->error('修改标签失败，请稍后再试！');
    }

    /**
     * 删除标签
     * Date: 2019-02-22 10:21
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionDelTag()
    {
        $this->basicValidate($this->tag, 'tagDelete');
        $data = $this->request_data;
        $res = $this->tag->deleteAll(['id' => $data['id']]);
        return $res ? $this->success('删除标签成功！')
            : $this->error('删除标签失败，请稍后再试！');
    }
}
