<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-04
 * Time: 09:39
 */

namespace backend\controllers;


use app\models\Link;

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
     */
    public function actionLinkList()
    {
        $data = $this->get();
        $this->link->scenario = 'linkList';
        $this->link->attributes = $data;
        if($this->link->validate()){
            $query = $this->link->find();
            $total = $query->count();
            if(isset($data['name']) && !empty($data['name'])){
                $query = $query->where(['like','name',$data['name']]);
            }
            $list = $query
                ->offset(($data['pageNum'] - 1) * $data['pageSize'])
                ->limit($data['pageSize'])
                ->orderBy('created_at DESC')
                ->all();
            $data = compact('list','total');
            return !empty($list)?$this->success('获取友联列表成功！',$data)
                :$this->error('暂无友联列表数据！');
        }
        return $this->error(current($this->link->firstErrors));
    }

    /**
     * 添加友联
     * Date: 2019-03-04 09:54
     * @return \yii\web\Response
     */
    public function actionLinkAdd()
    {
        $data = $this->post();
        $this->link->scenario = 'addLink';
        $this->link->attributes = $data;
        if($this->link->validate()){
            $attr = ['name','url','description','created_at','updated_at','order'];
            $res = $this->link->save(false,$attr);
            return $res?$this->success('添加友联成功！')
                :$this->error('添加友联失败，请稍后再试！');
        }
        return $this->error(current($this->link->firstErrors));
    }

    /**
     * 修改友联
     * Date: 2019-03-04 09:56
     * @return \yii\web\Response
     */
    public function actionLinkUpdate()
    {
        $data = $this->post();
        $this->link->scenario = 'addUpdate';
        $this->link->attributes = $data;
        if($this->link->validate()){
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
        return $this->error(current($this->link->firstErrors));
    }

    /**
     * 删除标签
     * Date: 2019-03-04 10:01
     * @return \yii\web\Response
     */
    public function actionDelLink()
    {
        $data = $this->post();
        $this->link->scenario = 'delLink';
        $this->link->attributes = $data;
        if($this->link->validate()){
            $res = $this->link->deleteAll(['id'=>$data['id']]);
            return $res?$this->success('删除标签成功！')
                :$this->error('删除标签失败，请稍后再试！');
        }
        return $this->error(current($this->link->firstErrors));
    }
}
