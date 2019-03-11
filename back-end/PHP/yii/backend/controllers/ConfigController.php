<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-11
 * Time: 09:34
 */

namespace backend\controllers;


use app\models\WebConfig;

class ConfigController extends BasicController
{
    public $enableCsrfValidation = false;
    private $config;
    //社交地址type
    private const SOCIAL_TYPE = 1;
    //footer内容配置
    private const FOOTER_TYPE = 2;
    //其他配置类型
    private const OTHER_CONFIG_TYPE = 3;

    public function init()
    {
        parent::init();
        $this->config = new WebConfig();
    }

    /**
     * 获取配置列表
     * Date: 2019-03-11 09:42
     * @return \yii\web\Response
     */
    public function actionConfigList()
    {
        $data = $this->get();
        $this->config->scenario = 'configList';
        $this->config->attributes = $data;
        if($this->config->validate()){
            $query = $this->config->find();
            if(isset($data['type']) && !empty($data['type'])){
                $query = $query->where(['type'=>$data['type']]);
            }
            $total = (int)($query->count());
            $list = $query
                ->offset(($data['pageNum'] - 1) * $data['pageSize'])
                ->limit($data['pageSize'])
                ->orderBy('created_at DESC')
                ->all();
            $data = compact('list','total');
            return !empty($list)?$this->success('获取配置列表成功！',$data)
                :$this->error('暂无配置列表数据！');
        }
        return $this->error(current($this->config->firstErrors));
    }

    /**
     * 添加配置项
     * Date: 2019-03-11 09:50
     * @return \yii\web\Response
     */
    public function actionAddConfig()
    {
        $data = $this->post();
        $this->config->scenario = 'addConfig';
        $this->config->attributes = $data;
        if($this->config->validate()){
            $check = $this->checkConfigVal($data);
            if(!$check['flag']){
                return $this->error($check['msg']);
            }
            $exist = $this->config->where(['name' => $data['name'], 'type' => $data['type']])->one();
            if (!isset($exist['id']))
            {
                $this->config->name = $data['name'];
                $this->config->type = $data['type'];
                $this->config->title = $data['title'];
                $this->config->val = $data['val'];
                $res = $this->config->save(false);
                return $res ? $this->success('添加配置项成功!')
                    : $this->error('添加配置项失败,请稍后再试!');
            } else
            {
                return $this->error('已添加该配置,请勿重复添加!');
            }
        }
        return $this->error(current($this->config->firstErrors));
    }

    //检查数据
    private function checkConfigVal($data): array
    {
        $count = $this->config->where(['name' => $data['name'], 'type' => $data['type']])->count();
        if ($data['type'] == self::FOOTER_TYPE)
        {
            if ($data['name'] == 'copyright' && $count >= 1)
            {
                return ['msg'=>'版权信息只能有一条!','flag'=>false];
            }
            if ($count >= 4)
            {
                return ['msg'=>'footer内容每一项最多添加4个!','flag'=>false];
            }
        }
        return ['flag'=>true];
    }

    /**
     * 修改配置项
     * Date: 2019-03-11 09:56
     * @return \yii\web\Response
     */
    public function actionUpdateConfig()
    {
        $data = $this->post();
        $this->config->scenario = 'configUpdate';
        $this->config->attributes = $data;
        if($this->config->validate()){
            $check = $this->checkConfigVal($data);
            if(!$check['flag']){
                return $this->error($check['msg']);
            }
            $exist = $this->config->where(['name' => $data['name'], 'type' => $data['type']])->one();
            if (isset($exist['id']) && $exist['id'] != $data['id'])
            {
                return $this->error('已添加该配置,请勿重复添加!');
            }
            $exist_config = $this->config->findOne($data['id']);
            $exist_config->name = $data['name'];
            $exist_config->type = $data['type'];
            $exist_config->title = $data['title'];
            $exist_config->val = $data['val'];
            $res = $exist_config->save(false,['updated_at','name','type','title','val']);
            return $res ? $this->success('修改配置项成功!')
                : $this->error('修改配置项失败,请稍后再试!');
        }
        return $this->error(current($this->config->firstErrors));
    }

    /**
     * 删除配置
     * Date: 2019-03-11 10:06
     * @return \yii\web\Response
     */
    public function actionDelConfig()
    {
        $data = $this->post();
        $this->config->scenario = 'configDelete';
        $this->config->attributes = $data;
        if($this->config->validate()){
            $res = $this->config->deleteAll(['id'=>$data['id']]);
            return $res ? $this->success('删除配置成功!')
                : $this->error('删除配置失败,请稍后再试!');
        }
        return $this->error(current($this->config->firstErrors));
    }
}
