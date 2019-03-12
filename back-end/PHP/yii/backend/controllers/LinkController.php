<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-04
 * Time: 09:39
 */

namespace backend\controllers;


use app\models\Link;
use backend\exception\ValidateException;

class LinkController extends BasicController
{
    public $enableCsrfValidation = false;
    private $link;

    public function init()
    {
        parent::init();
        $this->link = new Link();
    }

    /**
     * 获取友联列表
     * Date: 2019-03-04 09:49
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionLinkList()
    {
        $this->basicValidate($this->link, 'linkList');
        $data = $this->request_data;
        $query = $this->link->find();
        if(isset($data['name']) && !empty($data['name'])){
            $query = $query->where(['like','name',$data['name']]);
        }
        $total = $query->count();
        $list = $query
            ->offset(($data['pageNum'] - 1) * $data['pageSize'])
            ->limit($data['pageSize'])
            ->orderBy('created_at DESC')
            ->all();
        $data = compact('list','total');
        return !empty($list)?$this->success('获取友联列表成功！',$data)
            :$this->error('暂无友联列表数据！');
    }

    /**
     * 添加友联
     * Date: 2019-03-04 09:54
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionLinkAdd()
    {
        $link = $this->basicValidate($this->link, 'addLink');
        $attr = ['name','url','description','created_at','updated_at','order'];
        $res = $link->save(false,$attr);
        return $res?$this->success('添加友联成功！')
            :$this->error('添加友联失败，请稍后再试！');
    }

    /**
     * 修改友联
     * Date: 2019-03-04 09:56
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionLinkUpdate()
    {
        $this->basicValidate($this->link, 'updateLink');
        $data = $this->request_data;
        $exist_link = $this->link->findOne($data['id']);
        $exist_link->name = $data['name'];
        $exist_link->url = $data['url'];
        $exist_link->description = $data['description'];
        $exist_link->order = $data['order'];
        $exist_link->is_show = $data['is_show'];
        $res = $exist_link->save(false);
        return $res?$this->success('修改友联成功！')
            :$this->error('修改友联失败，请稍后再试！');
    }

    /**
     * 删除标签
     * Date: 2019-03-04 10:01
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionDelLink()
    {
        $this->basicValidate($this->link, 'delLink');
        $data = $this->request_data;
        $res = $this->link->deleteAll(['id'=>$data['id']]);
        return $res?$this->success('删除标签成功！')
            :$this->error('删除标签失败，请稍后再试！');
    }
}
