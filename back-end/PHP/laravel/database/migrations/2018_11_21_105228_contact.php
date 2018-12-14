<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class Contact extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('contact', function (Blueprint $table) {
            $table->increments('id');
            $table->string('name',30)->comment('留言人姓名');
            $table->string('content',255)->comment('留言内容');
            $table->string('email',60)->comment('留言人email');
            $table->unsignedTinyInteger('is_reply')->default(0)->comment('是否回复,1是0否');
            $table->string('reply_content',255)->nullable()->comment('回复内容');
            $table->unsignedInteger('created_at')->nullable()->comment('创建时间');
            $table->unsignedInteger('updated_at')->nullable()->comment('修改时间');
            $table->unsignedInteger('replied_at')->nullable()->comment('回复时间');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('contact');
    }
}
