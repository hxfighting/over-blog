<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-08
 * Time: 09:15
 */

namespace backend\controllers;


class CommentController extends BasicController
{
    public $enableCsrfValidation = false;

    public function actionCommentList()
    {
        $data = $this->get();
        $this->user->scenario = 'userList';
        $this->user->attributes = $data;
        if($this->user->validate()){
            $query = $this->user->find();
            if(isset($data['search']) && !empty($data['search'])){
                $query = $query->where(['like','name',$data['search']])
                    ->orWhere(['like','email',$data['search']]);
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
        return $this->error(current($this->user->firstErrors));
    }
}
