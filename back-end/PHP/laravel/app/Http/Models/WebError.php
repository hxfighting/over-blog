<?php

namespace App\Http\Models;

use EloquentFilter\Filterable;

class WebError extends BasicModel
{
    use Filterable;
    protected $table = 'web_error';
    protected $dateFormat = 'U';
    protected $guarded = [];

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s', $value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s', $value);
    }

}
