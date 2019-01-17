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

//简化数据库连接
if(!function_exists('db')){
    function db(string $table,string $connection='mysql'){
        return \Illuminate\Support\Facades\DB::connection($connection)->table($table);
    }
}

//获取IP
if(!function_exists('getIP')){
    function getIP(){
        $user_IP = isset($_SERVER["HTTP_VIA"]) ? $_SERVER["HTTP_X_FORWARDED_FOR"] : $_SERVER["REMOTE_ADDR"];
        $user_IP = $user_IP ?? $_SERVER["REMOTE_ADDR"];
        return $user_IP;
    }
}


