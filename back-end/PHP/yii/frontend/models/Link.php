<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-20
 * Time: 13:46
 */

namespace frontend\models;


use yii\db\ActiveRecord;

class Link extends ActiveRecord
{
    public static function tableName()
    {
        return "link";
    }
}
