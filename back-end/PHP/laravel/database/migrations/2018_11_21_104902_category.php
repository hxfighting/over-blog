<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Category extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('category', function (Blueprint $table) {
            $table->increments('id');
            $table->string('title',50)->comment('分类名称');
            $table->string('url',100)->comment('分类URL');
            $table->unsignedInteger('pid')->comment('父ID');
            $table->unsignedInteger('type')->default(1)->comment('分类类型:1文章分类,2联系我,3说说');
            $table->unsignedInteger('is_article_category')->default(0)->comment('是否是文章分类');
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
        Schema::dropIfExists('category');
    }
}
