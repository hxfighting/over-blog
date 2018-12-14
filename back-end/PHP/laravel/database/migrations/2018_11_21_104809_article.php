<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Article extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('article', function (Blueprint $table) {
            $table->increments('id');
            $table->string('title',100)->comment('文章标题');
            $table->string('author',20)->comment('文章作者');
            $table->longText('content_html')->comment('文章html内容');
            $table->longText('content_md')->nullable()->comment('文章markdown内容');
            $table->string('keywords',255)->comment('文章关键词');
            $table->string('description',255)->comment('文章描述');
            $table->string('thumb',255)->comment('文章缩略图');
            $table->unsignedTinyInteger('is_show')->default(1)->comment('是否显示,1是0否');
            $table->unsignedTinyInteger('is_original')->default(1)->comment('是否原创,1是0否');
            $table->unsignedTinyInteger('is_top')->default(0)->comment('是否置顶,1是0否');
            $table->unsignedInteger('click')->default(0)->comment('文章点击次数');
            $table->unsignedInteger('category_id')->comment('文章分类ID');
            $table->unsignedInteger('deleted_at')->nullable()->comment('删除时间');
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
        Schema::dropIfExists('article');
    }
}
