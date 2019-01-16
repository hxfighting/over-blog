<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Link extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('link', function (Blueprint $table) {
            $table->increments('id');
            $table->string('url',255)->comment('友联url');
            $table->string('name',50)->comment('友联名称');
            $table->string('description',50)->comment('友联描述');
            $table->unsignedTinyInteger('order')->default(0)->comment('友联排序');
            $table->unsignedTinyInteger('is_show')->default(0)->comment('友联是否显示,1是,0否');
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
        Schema::dropIfExists('link');
    }
}
