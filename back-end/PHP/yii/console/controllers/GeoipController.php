<?php
/**
 * Created by huxin.
 * User: huxin
 * Date: 2019-03-07
 * Time: 11:14
 */

namespace console\controllers;


use yii\console\Controller;
use yii\helpers\Console;

class GeoipController extends Controller
{
    //geoip database 下载地址
    const GEOIP_DB_URL = 'https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz';

    public function actionUpdate()
    {
        $this->stdout("开始更新本地geoip数据库...\n", Console::BOLD);
        $path = __DIR__.'/../geoip.mmdb';

        try
        {// Get header response
            $headers = get_headers(self::GEOIP_DB_URL);
            if (substr($headers[0], 9, 3) != '200')
            {
                $message = '下载数据库失败. (' . substr($headers[0], 13) . ')';
                $message = $this->ansiFormat($message,Console::FG_RED);
                echo $message;
                return 1;
            }// Download zipped database to a system temp file
            $tmpFile = tempnam(sys_get_temp_dir(), 'maxmind');
            file_put_contents($tmpFile, fopen(self::GEOIP_DB_URL, 'r'));// Unzip and save database
            file_put_contents($path, gzopen($tmpFile, 'r'));// Remove temp file
            @unlink($tmpFile);
        } catch (\Exception $e)
        {
            $message = $this->ansiFormat($e->getMessage(),Console::FG_RED);
            echo $message;
            return 1;
        }
        $this->stdout("本地数据库文件更新成功！ ({$path})\n", Console::BOLD);
        return 0;
    }
}
