<?php

use Illuminate\Support\Facades\Schema;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Database\Migrations\Migration;

class User extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('user', function (Blueprint $table) {
            $table->increments('id');
            $table->unsignedTinyInteger('type')->comment('用户类型 1：QQ  2：微信 3：新浪微博');
            $table->string('name',30)->comment('第三方用户名称');
            $table->string('avatar',255)->comment('第三方用户头像');
            $table->string('openid',100)->comment('第三方用户openid');
            $table->string('access_token',100)->comment('access_token token（当是微信时存的是session_key）');
            $table->ipAddress('last_login_ip')->comment('最后登录IP');
            $table->unsignedInteger('login_times')->default(0)->comment('登录次数');
            $table->string('email',100)->nullable()->comment('邮箱');
            $table->unsignedTinyInteger('is_admin')->default(0)->comment('是否是管理员,1是,0否');
            $table->rememberToken();
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
        Schema::dropIfExists('user');
    }
}
