<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;

class Chat extends BasicModel
{
    use Filterable;
    protected $table = 'chat';
    protected $dateFormat = 'U';
    protected $fillable = ['content'];

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

}
