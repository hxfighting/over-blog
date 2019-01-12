<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;
use Watson\Rememberable\Rememberable;

class WebConfig extends BasicModel
{
    use Filterable,Rememberable;
    protected $table = 'web_config';
    protected $dateFormat = 'U';
    protected $guarded = [];
    protected $rememberCacheTag = 'h_web_config_cache';

    public function scopeType($query,int $type)
    {
        return $query->where('type',$type);
    }

    //获取列表统计
    public function getListCount(array $data): int
    {
        $search = $data['type'] ?? null;
        return $this->when($search,function ($q) use ($search){
            return $q->where('type',$search);
        })->count();
    }
}
