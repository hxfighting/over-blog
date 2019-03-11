<?php

namespace app\models;

use Yii;
use yii\behaviors\AttributeTypecastBehavior;
use yii\behaviors\TimestampBehavior;

/**
 * This is the model class for table "web_config".
 *
 * @property int $id
 * @property string $title 配置名称
 * @property string $name 配置变量名
 * @property string $val 配置值
 * @property int $type 字段类型,1:社交地址,2footer内容,3其他配置
 * @property int $created_at 创建时间
 * @property int $updated_at 修改时间
 */
class WebConfig extends \yii\db\ActiveRecord
{
    public $pageSize;
    public $pageNum;
    /**
     * {@inheritdoc}
     */
    public static function tableName()
    {
        return 'web_config';
    }

    public function behaviors()
    {
        return [
            'typecast' => [
                'class' => AttributeTypecastBehavior::class,
                'attributeTypes' => [
                    'created_at' => function ($value) {
                        return date('Y-m-d H:i:s', $value);
                    },
                    'updated_at' => function ($value) {
                        return date('Y-m-d H:i:s', $value);
                    }
                ],
                'typecastAfterValidate' => true,
                'typecastBeforeSave' => false,
                'typecastAfterFind' => true,
            ],
            [
                'class' => TimestampBehavior::class
            ]
        ];
    }

    public function init()
    {
        parent::init();
        $this->on(self::EVENT_AFTER_INSERT,[$this,'putConfigFile']);
    }

    public function putConfigFile($event)
    {
        $config_data = $this->find()
            ->where(['type'=>3])
            ->select('name', 'val')
            ->asArray()
            ->all();
        $config_data = array_combine(array_column($config_data,'name'),array_column($config_data,'val'));
        $path = __DIR__.'../config/webConfig.php';
        chmod($path,777);
        $str = '<?php return ' . var_export($config_data, true) . ';';    //将得到数组转换成字符串
        file_put_contents($path, $str); //写入文件
    }

    /**
     * {@inheritdoc}
     */
    public function rules()
    {
        return [
            [['pageSize', 'pageNum'], 'required', 'on' => 'configList'],
            [['pageSize', 'pageNum'], 'integer', 'on' => 'configList'],
            [['pageSize', 'pageNum'], 'number', 'min' => 1, 'on' => 'configList'],
            ['id', 'required', 'on' => ['configDelete','configUpdate']],
            ['id', 'integer', 'on' => ['configDelete','configUpdate']],
            ['id', 'exist', 'on' => ['configDelete','configUpdate']],
            [['title', 'name', 'val', 'type'], 'required','on' => 'addConfig'],
            ['val', 'string','on' => 'addConfig'],
            ['type', 'integer','on' => 'addConfig'],
            [['title', 'name'], 'string', 'max' => 100,'min' => 2,'on' => 'addConfig'],
        ];
    }

    /**
     * {@inheritdoc}
     */
    public function attributeLabels()
    {
        return [
            'id' => 'ID',
            'title' => 'Title',
            'name' => 'Name',
            'val' => 'Val',
            'type' => 'Type',
            'created_at' => 'Created At',
            'updated_at' => 'Updated At',
        ];
    }
}
