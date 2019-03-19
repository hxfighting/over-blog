<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-19
 * Time: 11:26
 */

namespace frontend\models;


use yii\db\ActiveRecord;

class Article extends ActiveRecord
{
    public static function tableName()
    {
        return "article";
    }
}
