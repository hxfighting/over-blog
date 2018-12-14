<?php namespace App\Http\ModelFilters;

use EloquentFilter\ModelFilter;

class ArticleCommentFilter extends ModelFilter
{
    /**
    * Related Models that have ModelFilters as well as the method on the ModelFilter
    * As [relationMethod => [input_key1, input_key2]].
    *
    * @var array
    */
    public $relations = [];

    public function article($value)
    {
        return $this->where('article_id',$value);
    }

    public function order($value)
    {
        return $this->latest($value);
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
