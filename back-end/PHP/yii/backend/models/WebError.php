<?php

namespace app\models;

use Yii;
use yii\behaviors\TimestampBehavior;
use yii\db\ActiveRecord;

/**
 * This is the model class for table "web_error".
 *
 * @property int    $id
 * @property string $type       类型
 * @property string $code       状态code
 * @property string $mes        错误信息
 * @property string $url        请求URL
 * @property int    $created_at 创建时间
 * @property int    $updated_at 修改时间
 */
class WebError extends \yii\db\ActiveRecord
{
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'web_error';
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['type', 'code', 'mes', 'url'], 'required'],
            ['type', 'string', 'max' => 20],
            [['mes', 'url'], 'string', 'max' => 255],
            ['url', 'url']
        ];
    }

    public function behaviors()
    {
        return [
            [
                'class' => TimestampBehavior::className(),
                'attributes' => [
                    ActiveRecord::EVENT_BEFORE_INSERT => ['created_at', 'updated_at'],
                    ActiveRecord::EVENT_BEFORE_UPDATE => ['updated_at'],
                ],
            ],
        ];
    }
}
