<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;
use Watson\Rememberable\Rememberable;

class Tag extends BasicModel
{
    use Filterable,Rememberable;
    protected $table = 'tag';
    protected $dateFormat = 'U';
    protected $fillable = ['name'];
    protected $rememberCacheTag = 'h_tag_cache';

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s,',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }
}
