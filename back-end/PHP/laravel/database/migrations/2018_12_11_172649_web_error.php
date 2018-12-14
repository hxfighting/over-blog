<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class WebError extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('web_error', function (Blueprint $table) {
            $table->increments('id');
            $table->string('type',20)->comment('类型');
            $table->string('code',20)->comment('状态code');
            $table->string('mes',255)->comment('错误信息');
            $table->string('url',255)->comment('请求URL');
            $table->unsignedInteger('created_at')->nullable()->comment('创建时间');
            $table->unsignedInteger('updated_at')->nullable()->comment('修改时间');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('web_error');
    }
}
