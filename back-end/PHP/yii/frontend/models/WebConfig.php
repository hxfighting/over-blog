<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-19
 * Time: 09:32
 */

namespace frontend\models;


use yii\db\ActiveRecord;

class WebConfig extends ActiveRecord
{
    public static function tableName()
    {
        return "web_config";
    }
}
