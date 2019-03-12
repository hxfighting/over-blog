<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-02-22
 * Time: 11:26
 */

namespace backend\controllers;


use app\models\Category;
use backend\exception\ValidateException;
use yii\db\Exception;

class CategoryController extends BasicController
{
    //文章分类type
    private const ARTICLE_TYPE = 1;

    public $enableCsrfValidation = false;
    private $category;

    public function init()
    {
        parent::init();
        $this->category = new Category();
    }

    /**
     * 获取分类列表
     * Date: 2019-02-25 11:07
     * @return \yii\web\Response
     */
    public function actionCategoryList()
    {
        $list = $this->category->find()
            ->with('children')
            ->where(['pid'=>0])
            ->asArray()
            ->all();
        if(!empty($list)){
            foreach ($list as &$item)
            {
                $item['pid'] = (int)$item['pid'];
                $item['type'] = (int)$item['type'];
                $item['expand'] = true;
            }
            unset($item);
            return $this->success('获取分类列表成功！',$list);
        }
        return $this->error('获取分类列表失败，请稍后再试！');
    }

    /**
     * 添加分类
     * Date: 2019-02-25 11:21
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionCategoryAdd()
    {
        $category = $this->basicValidate($this->category,'categoryAdd');
        $res = $category->save(false);
        return $res?$this->success('添加分类成功！')
            :$this->error('添加分类失败，请稍后再试！');
    }

    /**
     * 修改分类信息
     * Date: 2019-02-25 11:23
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionCategoryUpdate()
    {
        $this->basicValidate($this->category,'categoryUpdate');
        $exist_category = $this->category->findOne(($this->request_data)['id']);
        $exist_category->title = ($this->request_data)['title'];
        $res = $exist_category->save(false,['title','updated_at']);
        return $res?$this->success('修改分类成功！')
            :$this->error('修改分类失败，请稍后再试！');
    }

    /**
     * 删除分类
     * Date: 2019-02-25 11:25
     * @return \yii\web\Response
     * @throws ValidateException
     */
    public function actionDelCategory()
    {
        $this->basicValidate($this->category,'delCategory');
        $transaction = \Yii::$app->db->beginTransaction();
        try
        {
            $exist_tag = $this->category->findOne(($this->request_data)['id']);
            if ($exist_tag->type != self::ARTICLE_TYPE)
            {
                return $this->error('不能删除此分类！');
            }
            $this->category->deleteAll(['pid' => ($this->request_data)['id']]);
            $exist_tag->delete();
            $transaction->commit();
            return $this->success('删除分类成功！');
        } catch (Exception $e)
        {
            $transaction->rollBack();
            return $this->error('删除分类失败，请稍后再试！');
        }
    }
}
