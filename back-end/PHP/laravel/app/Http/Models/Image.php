<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;

class Image extends BasicModel
{
    use Filterable;
    protected $table = 'image';
    protected $dateFormat = 'U';
    protected $guarded = [];
    //多态关联时轮播图对应的图片类型
    private $rotation_image_type = 'App\Http\Models\RotationImage';
    //多态关联时照片对应的图片类型
    private $photo_image_type = 'App\Http\Models\Photo';

    public function image()
    {
        return $this->morphTo();
    }

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    //获取轮播图
    public function rotationImage()
    {
        return $this->with('image')
            ->filter(['image_type'=>$this->rotation_image_type])
            ->get();
    }

    //获取照片
    public function photoImage()
    {
        return $this->with('image')
            ->filter(['image_type'=>$this->photo_image_type])
            ->get();
    }
}
