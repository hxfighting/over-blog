<?php

namespace App\Http\Models;


use EloquentFilter\Filterable;

class Contact extends BasicModel
{
    use Filterable;
    protected $table = 'contact';
    protected $dateFormat = 'U';
    protected $fillable = ['name','email','content'];

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    //统计相关留言
    public function getContactCount(array $data)
    {
        $name = $data['search'] ?? null;
        return $this->when($name,function ($q) use ($name){
            return $q->where('name','like','%'.$name.'%')
                ->orWhere('email','like','%'.$name.'%');
        })->count();
    }
}
