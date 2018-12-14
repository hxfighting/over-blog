<?php namespace App\Http\ModelFilters;

use EloquentFilter\ModelFilter;

class ImageFilter extends ModelFilter
{
    /**
    * Related Models that have ModelFilters as well as the method on the ModelFilter
    * As [relationMethod => [input_key1, input_key2]].
    *
    * @var array
    */
    public $relations = [];

    public function imageType($value)
    {
        return $this->where('image_type',$value);
    }

    public function order($value)
    {
        return $this->latest($value);
    }

    //初始化按时间排序
    public function setUp()
    {
        if (!$this->input('order'))  {
            $this->push('order', 'created_at');
        }
        if(!$this->input('pageSize')){
            $this->push('pageSize', 10);
        }
        if(!$this->input('pageNum')){
            $this->push('pageNum', 1);
        }
    }
}
