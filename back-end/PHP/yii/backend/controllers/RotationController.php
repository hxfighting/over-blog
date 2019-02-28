<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-28
 * Time: 10:11
 */

namespace backend\controllers;


use app\models\RotationImage;
use yii\db\Exception;
use yii\db\Query;

class RotationController extends BasicController
{
    public $enableCsrfValidation = false;
    //轮播图type
    private $rotation_type = 'App\Http\Models\RotationImage';

    /**
     * 获取轮播图列表
     * Date: 2019-02-28 11:26
     * @return \yii\web\Response
     */
    public function actionRotationList()
    {
        $query = new Query();
        $list = $query->from('rotation_image as r')->leftJoin('image as i',"r.id = i.image_id")
            ->select('r.*,i.image_url,i.id as img_id')
            ->where(['i.image_type'=>$this->rotation_type])
            ->orderBy('r.created_at desc')
            ->all();
        $list = $this->handleRotationListData($list);
        return !empty($list)?$this->success('获取轮播图成功！',$list)
            :$this->error('暂无轮播图数据！');
    }

    /**
     * 处理轮播图数据
     * Date: 2019-02-28 11:26
     * @param array $list
     * @return array
     */
    private function handleRotationListData(array $list)
    {
        if(!empty($list)){
            foreach ($list as &$item)
            {
                $item['created_at'] = date('Y/m/d H:i:s',$item['created_at']);
                $item['updated_at'] = date('Y/m/d H:i:s',$item['updated_at']);
                $item['image'] = ['words'=>$item['words'],'id'=>$item['id']];
                $item['id'] = $item['img_id'];
            }
            unset($item);
        }
        return $list;
    }

    /**
     * 添加轮播图
     * Date: 2019-02-28 11:46
     * @return \yii\web\Response
     */
    public function actionRotationAdd()
    {
        $data = $this->post();
        $rotation = new RotationImage();
        $rotation->scenario = 'rotationAdd';
        $rotation->attributes = $data;
        if($rotation->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $rotation->save(false, ['words','created_at','updated_at']);
                $image_data = $this->getImageData($rotation->id, $data['image_url']);
                \Yii::$app->db->createCommand()->insert('image',$image_data)->execute();
                $tr->commit();
                return $this->success('添加轮播图成功！');
            } catch (Exception $e)
            {
                $tr->rollBack();
                return $this->error('添加轮播图失败，请稍后再试！');
            }
        }
        return $this->error(current($rotation->firstErrors));
    }

    /**
     * 获取轮播图图片数据
     * Date: 2019-02-28 11:46
     * @param int    $rotation_id
     * @param string $image_url
     * @return array
     */
    private function getImageData(int $rotation_id,string $image_url)
    {
        $image_data = [
            'image_id'=>$rotation_id,
            'image_type'=>$this->rotation_type,
            'created_at'=>time(),
            'updated_at'=>time(),
            'image_url'=>$image_url
        ];
        return $image_data;
    }

    /**
     * 修改轮播图信息
     * Date: 2019-02-28 13:22
     * @return \yii\web\Response
     */
    public function actionRotationUpdate()
    {
        $data = $this->post();
        $rotation = new RotationImage();
        $rotation->scenario = 'rotationUpdate';
        $rotation->attributes = $data;
        if($rotation->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $exist_rotation = $rotation->findOne($data['id']);
                $exist_rotation->words = $data['words'];
                $exist_rotation->save(false, ['words','updated_at']);
                $type = implode('\\\\',explode('\\',$this->rotation_type));
                \Yii::$app->db
                    ->createCommand()
                    ->update('image',['image_url'=>$data['image_url']],"image_id = {$data['id']} and image_type = '{$type}'")
                    ->execute();
                $tr->commit();
                return $this->success('修改轮播图成功！');
            } catch (Exception $e)
            {
                $tr->rollBack();
                return $this->error('修改轮播图失败，请稍后再试！');
            }
        }
        return $this->error(current($rotation->firstErrors));
    }

    /**
     * 删除轮播图
     * Date: 2019-02-28 13:57
     * @return \yii\web\Response
     */
    public function actionDelRotation()
    {
        $data = $this->post();
        $rotation = new RotationImage();
        $rotation->scenario = 'delRotation';
        $rotation->attributes = $data;
        if($rotation->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $rotation->deleteAll(['id'=>$data['id']]);
                $type = implode('\\\\',explode('\\',$this->rotation_type));
                \Yii::$app->db
                    ->createCommand()
                    ->delete('image',"image_id = {$data['id']} and image_type = '{$type}'")
                    ->execute();
                $tr->commit();
                return $this->success('删除轮播图成功！');
            } catch (Exception $e)
            {
                $tr->rollBack();
                return $this->error('删除轮播图失败，请稍后再试！');
            }
        }
        return $this->error(current($rotation->firstErrors));
    }
}
