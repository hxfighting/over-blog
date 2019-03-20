<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-20
 * Time: 11:54
 */

namespace frontend\models;


use yii\db\ActiveRecord;

class Comment extends ActiveRecord
{
    public static function tableName()
    {
        return "article_comment";
    }
}
