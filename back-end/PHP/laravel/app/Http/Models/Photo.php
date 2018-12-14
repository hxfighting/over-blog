<?php

namespace App\Http\Models;

use Illuminate\Database\Eloquent\Model;

class Photo extends Model
{
    protected $table = 'photo';
    protected $dateFormat = 'U';

    //多态关联
    public function images()
    {
        return $this->morphMany(Image::class,'image');
    }
}
