<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Support\Facades\Log;
use Overtrue\EasySms\EasySms;
use Overtrue\EasySms\Exceptions\NoGatewayAvailableException;

class SendSms extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'sms:send';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'send sms';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $config = config('sms');
        $sms = new EasySms($config);
        try
        {
            for ($i = 0; $i < 2; $i++)
            {
                if ($i == 0)
                {
                    $number = mt_rand(0, 9);
                    $phone = config('blog.phone_one');
                    $sms->send($phone, [
                        'template' => config('blog.phone_one_template'),
                        'data' => [
                            'code' => $number
                        ],
                    ]);
                } else
                {
                    $date = strtotime('1994-03-10');
                    $date2 = time();
                    $days = intval(floor(($date2 - $date) / 3600 / 24));
                    $phone = config('blog.phone_two');
                    $sms->send($phone, [
                        'template' => config('blog.phone_two_template'),
                        'data' => [
                            'number' => $days
                        ],
                    ]);
                }

            }
        } catch (NoGatewayAvailableException $e)
        {
            Log::error((string)$e->results);
        }
    }
}
