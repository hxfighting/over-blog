<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class WebConfig extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('web_config', function (Blueprint $table) {
            $table->increments('id');
            $table->string('title',100)->comment('配置名称');
            $table->string('name',100)->comment('配置变量名');
            $table->text('val')->comment('配置值');
            $table->unsignedTinyInteger('order')->comment('配置排序');
            $table->string('description',255)->comment('配置描述');
            $table->unsignedTinyInteger('field_type')->comment('字段类型,1:input,2:radio,3:textarea');
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
        Schema::dropIfExists('web_config');
    }
}
