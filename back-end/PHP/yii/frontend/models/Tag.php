<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-19
 * Time: 09:41
 */

namespace frontend\models;


use yii\db\ActiveRecord;

class Tag extends ActiveRecord
{
    public static function tableName()
    {
        return "tag";
    }
}
