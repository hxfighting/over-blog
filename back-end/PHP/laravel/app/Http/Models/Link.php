<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;
use Watson\Rememberable\Rememberable;

class Link extends BasicModel
{
    use Filterable,Rememberable;
    protected $table = 'link';
    protected $dateFormat = 'U';
    protected $guarded = [];
    protected $rememberCacheTag = 'h_link_cache';

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    //统计相关友联
    public function getLinkCount(array $data)
    {
        $name = $data['name'] ?? null;
        return $this->when($name,function ($q) use ($name){
            return $q->where('name','like','%'.$name.'%');
        })->count();
    }
}
