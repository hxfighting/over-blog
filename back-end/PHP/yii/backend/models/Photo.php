<?php

namespace app\models;

use Yii;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "photo".
 *
 * @property int $id
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class Photo extends \yii\db\ActiveRecord
{
    public $image_url;
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'photo';
    }

    public function behaviors()
    {
        return [
            [
                'class' => TimestampBehavior::class
            ]
        ];
    }

    public function rules()
    {
        return [
            ['id','required','on' => 'delPhoto'],
            ['id','integer','on' => 'delPhoto'],
            ['image_url','required','on' => ['photoAdd','photoUpdate']],
            ['image_url','url','on' => ['photoAdd','photoUpdate']]
        ];
    }


}
