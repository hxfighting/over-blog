<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-13
 * Time: 10:05
 */

namespace backend\controllers;


use backend\exception\ValidateException;
use yii\web\UploadedFile;

class UploadController extends BasicController
{
    public $enableCsrfValidation = false;

    private $allow_ext = ['png', 'jpg', 'jpeg', 'gif'];

    private $max_size = 20 * 1024 * 1024;

    /**
     * 图片上传
     * Date: 2019-03-13 11:49
     * @return \yii\web\Response
     * @throws ValidateException
     * @throws \yii\base\InvalidConfigException
     */
    public function actionUpload()
    {
        $file = UploadedFile::getInstanceByName('file');
        $this->checkFile($file);
        $file->getBaseName();
        $qiniu = \Yii::$app->get('qiniu');
        $disk = $qiniu->getDisk();
        $contents = file_get_contents($file->tempName);
        $image_key = md5($contents);
        if ($disk->has($image_key))
        {
            return $this->success('上传图片成功', $image_key);
        } else
        {
            $picture_path = $disk->put($image_key, $contents);
            return $picture_path ? $this->success('上传图片成功', $image_key)
                : $this->error('上传图片失败,请稍后再试');
        }
    }

    /**
     * 检查图片
     * Date: 2019-03-13 11:49
     * @param $file
     * @return bool
     * @throws ValidateException
     */
    private function checkFile($file)
    {
        $size = $file->size;
        $ext = $file->getExtension();
        if (!in_array($ext, $this->allow_ext))
        {
            throw new ValidateException('图片允许的格式有png、jpg、jpeg、gif');
        }
        if ($size > $this->max_size)
        {
            throw new ValidateException('图片最大20M');
        }
        return true;
    }
}
