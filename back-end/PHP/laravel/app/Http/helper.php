<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2018/6/6 0006
 * Time: 下午 15:26
 */


//统一返回格式
if (!function_exists('renderSuccess'))
{
    function renderSuccess($message, $data = null, $code = 200,$httpCode=200)
    {
        return response()->json([
            'data'      => $data ?? null,
            'msg'       => $message ?? null,
            'code'      => $code
        ],$httpCode,[],JSON_UNESCAPED_UNICODE);
    }
}

if (!function_exists('renderError'))
{
    function renderError($message, $data = null, $code = -1,$httpCode=200)
    {
        return response()->json([
            'data'      => $data ?? null,
            'msg'       => $message ?? null,
            'code'      => $code
        ],$httpCode,[],JSON_UNESCAPED_UNICODE);
    }
}


//转换db类获取的结果集
if(!function_exists('transferToArray')){
    function transferToArray(array $value){
        return array_map(function ($val){
            return (array)$val;
        },$value);
    }
}


