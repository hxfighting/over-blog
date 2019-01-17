<?php namespace App\Http\ModelFilters;

use EloquentFilter\ModelFilter;

class ArticleFilter extends ModelFilter
{
    /**
    * Related Models that have ModelFilters as well as the method on the ModelFilter
    * As [relationMethod => [input_key1, input_key2]].
    *
    * @var array
    */
    public $relations = ['tags'=>['tag_id']];

    public function search($title)
    {
        return $this->whereLike('title',$title);
    }

    public function category($value)
    {
        return $this->where('category_id',$value);
    }

    public function id($id)
    {
        return $this->where('id',$id);
    }

    public function show($value)
    {
        return $this->where('is_show',$value);
    }

    public function order($value)
    {
        return $value=='created_at'?$this->latest($value):$this->orderByDesc('is_top')->latest('created_at');
    }

    public function pageSize($value)
    {
        return $this->take($value);
    }

    public function pageNum($value)
    {
        $pageSize = $this->input('pageSize');
        return $this->skip(($value - 1) * $pageSize);
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
