<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;

class Tag extends BasicModel
{
    use Filterable;
    protected $table = 'tag';
    protected $dateFormat = 'U';
    protected $fillable = ['name'];

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s,',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }
}
