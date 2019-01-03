<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;

class WebConfig extends BasicModel
{
    use Filterable;
    protected $table = 'web_config';
    protected $dateFormat = 'U';
    protected $guarded = [];

    //获取列表统计
    public function getListCount(array $data): int
    {
        $search = $data['type'] ?? null;
        return $this->when($search,function ($q) use ($search){
            return $q->where('type',$search);
        })->count();
    }
}
