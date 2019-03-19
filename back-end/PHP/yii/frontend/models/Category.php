<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-15
 * Time: 10:39
 */

namespace frontend\models;


use yii\db\ActiveRecord;

class Category extends ActiveRecord
{
    public static function tableName()
    {
        return "category";
    }

    public function getChildren()
    {
        return $this->hasMany(self::class,['pid'=>'id']);
    }
}
