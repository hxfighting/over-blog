<?php

namespace App\Http\Models;

use EloquentFilter\Filterable;
use Illuminate\Notifications\Notifiable;
use Illuminate\Contracts\Auth\MustVerifyEmail;
use Illuminate\Foundation\Auth\User as Authenticatable;
use Tymon\JWTAuth\Contracts\JWTSubject;

class User extends Authenticatable implements JWTSubject
{
    use Notifiable,Filterable;
    protected $table = 'user';
    protected $dateFormat = 'U';
    protected $hidden = ['access_token','remember_token','openid'];

    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    /**
     * Return a key value array, containing any custom claims to be added to the JWT.
     *
     * @return array
     */
    public function getJWTCustomClaims()
    {
        return [];
    }

    public function getCreatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    public function getUpdatedAtAttribute($value)
    {
        return date('Y/m/d H:i:s',$value);
    }

    /**
     * 利用扩展包torann/geoip转换IP地址
     * https://github.com/Torann/laravel-geoip
     * @param $value
     * @return string
     */
    public function getLastLoginIpAttribute($value)
    {
        $data = geoip($value)->toArray();
        return $this->handleAddress($data);
    }

    //处理转换后的地址数据
    private function handleAddress(array $address): string
    {
        switch ($address['iso_code']){
            case 'CN':
                return $address['country'].'-'.$address['state_name'].'-'.$address['city'];
            case 'TW':
                return '中国-台湾';
            case 'MO':
                return '中国-澳门';
            default:
                return $address['country'].'-'.$address['state_name'].'-'.$address['city'];
        }
    }

    //列表查询统计相关用户
    public function getUserCount(array $data):int
    {
        $search = $data['search'] ?? null;
        return $this->when($search,function ($q) use ($search){
            return $q->where('name','like','%'.$search.'%')
                ->orWhere('email','like','%'.$search.'%');
        })->count();
    }
}
