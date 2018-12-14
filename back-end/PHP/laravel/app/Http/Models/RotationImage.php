<?php

namespace App\Http\Models;

use Illuminate\Database\Eloquent\Model;

class RotationImage extends Model
{
    protected $table = 'rotation_image';
    protected $dateFormat = 'U';
    protected $fillable = ['words'];

    //多态关联
    public function images()
    {
        return $this->morphMany(Image::class,'image');
    }
}
