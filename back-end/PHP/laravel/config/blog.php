<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2018/12/5
 * Time: 10:53
 */

return [
    /**
     * 博客首页URL
     */
    'blog_home'=>env('BLOG_HOME_URL',''),
    /**
     * 首页分类cache的key
     */
    'category_cache_key'=>'category_cache',
    /**
     * 首页轮播图cache的key
     */
    'rotation_cache_key'=>'rotation_cache',
    /**
     * 标签缓存cache
     */
    'tag_cache_key'=>'typ_tag_cache',
    /**
     * 评论缓存cache
     */
    'comment_cache_key'=>'comment_cache',
    /**
     * 热门文章cache
     */
    'hot_article_cache_key'=>'hot_article_cache',
    /**
     * 友联cache
     */
    'link_cache_key'=>'link_cache'
];
