<?php

namespace App\Console\Commands;

use Illuminate\Console\Command;
use Illuminate\Support\Facades\Cache;

class UpdateIndexViewCount extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'index:view';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'update index view count';

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
        $key = config('blog.blog_index_count_key');
        $count = Cache::get($key);
        if($count){
            db('web_config')->where('name','blog_view_count')->update(['val'=>$count]);
        }
    }
}
