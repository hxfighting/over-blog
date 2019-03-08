<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-08
 * Time: 11:51
 */

namespace console\models;


use yii\db\ActiveRecord;

class Article extends ActiveRecord
{
    public static function tableName()
    {
        return 'article';
    }
}
