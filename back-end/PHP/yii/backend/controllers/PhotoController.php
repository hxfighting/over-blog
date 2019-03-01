<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-01
 * Time: 10:00
 */

namespace backend\controllers;


use app\models\Photo;
use yii\db\Query;

class PhotoController extends BasicController
{
    public $enableCsrfValidation = false;
    //照片type
    private $photo_image_type = 'App\Http\Models\Photo';

    private $photo;

    public function init()
    {
        parent::init();
        $this->photo = new Photo();
    }

    /**
     * 获取列表数据
     * Date: 2019-03-01 10:20
     * @return \yii\web\Response
     */
    public function actionPhotoList()
    {
        $query = new Query();
        $list = $query->from('photo as r')->leftJoin('image as i',"r.id = i.image_id")
            ->select('r.*,i.image_url,i.id as img_id')
            ->where(['i.image_type'=>$this->photo_image_type])
            ->orderBy('r.created_at desc')
            ->all();
        $list = $this->handlePhotoListData($list);
        return !empty($list)?$this->success('获取照片成功！',$list)
            :$this->error('暂无照片数据！');
    }

    /**
     * 处理列表数据
     * Date: 2019-03-01 10:19
     * @param array $list
     * @return array
     */
    private function handlePhotoListData(array $list)
    {
        if(!empty($list)){
            foreach ($list as &$item)
            {
                $item['created_at'] = date('Y/m/d H:i:s',$item['created_at']);
                $item['updated_at'] = date('Y/m/d H:i:s',$item['updated_at']);
                $item['image'] = ['id'=>$item['id']];
                $item['id'] = $item['img_id'];
            }
            unset($item);
        }
        return $list;
    }

    /**
     * 添加照片
     * Date: 2019-03-01 10:19
     * @return \yii\web\Response
     */
    public function actionPhotoAdd()
    {
        $data = $this->post();
        $this->photo->scenario = 'photoAdd';
        $this->photo->attributes = $data;
        if($this->photo->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $this->photo->save(false, ['created_at','updated_at']);
                $image_data = $this->getImageData($this->photo->id, $data['image_url']);
                \Yii::$app->db->createCommand()->insert('image',$image_data)->execute();
                $tr->commit();
                return $this->success('添加照片成功！');
            } catch (\Exception $e)
            {
                $tr->rollBack();
                return $this->error('添加照片失败，请稍后再试！');
            }
        }
        return $this->error(current($this->photo->firstErrors));
    }

    /**
     * 获取图片数据
     * Date: 2019-03-01 10:19
     * @param int    $photo_id
     * @param string $image_url
     * @return array
     */
    private function getImageData(int $photo_id,string $image_url)
    {
        $image_data = [
            'image_id'=>$photo_id,
            'image_type'=>$this->photo_image_type,
            'created_at'=>time(),
            'updated_at'=>time(),
            'image_url'=>$image_url
        ];
        return $image_data;
    }

    /**
     * 修改照片
     * Date: 2019-03-01 10:19
     * @return \yii\web\Response
     */
    public function actionPhotoUpdate()
    {
        $data = $this->post();
        $this->photo->scenario = 'photoUpdate';
        $this->photo->attributes = $data;
        if($this->photo->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $exist_photo = $this->photo->findOne($data['id']);
                $exist_photo->save(false, ['updated_at']);
                $type = implode('\\\\',explode('\\',$this->photo_type));
                \Yii::$app->db
                    ->createCommand()
                    ->update('image',['image_url'=>$data['image_url']],"image_id = {$data['id']} and image_type = '{$type}'")
                    ->execute();
                $tr->commit();
                return $this->success('修改照片成功！');
            } catch (\Exception $e)
            {
                $tr->rollBack();
                return $this->error('修改照片失败，请稍后再试！');
            }
        }
        return $this->error(current($this->photo->firstErrors));
    }

    /**
     * 删除照片
     * Date: 2019-03-01 10:18
     * @return \yii\web\Response
     */
    public function actionDelPhoto()
    {
        $data = $this->post();
        $this->photo->scenario = 'delPhoto';
        $this->photo->attributes = $data;
        if($this->photo->validate()){
            $tr = \Yii::$app->db->beginTransaction();
            try
            {
                $this->photo->deleteAll(['id'=>$data['id']]);
                $type = implode('\\\\',explode('\\',$this->photo_type));
                \Yii::$app->db
                    ->createCommand()
                    ->delete('image',"image_id = {$data['id']} and image_type = '{$type}'")
                    ->execute();
                $tr->commit();
                return $this->success('删除照片成功！');
            } catch (\Exception $e)
            {
                $tr->rollBack();
                return $this->error('删除照片失败，请稍后再试！');
            }
        }
        return $this->error(current($this->photo->firstErrors));
    }
}
