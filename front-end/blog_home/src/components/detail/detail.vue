<template>
    <div>
        <div id="main-content">
            <div class="box single_article">
                <div class="line"></div>
                <h2 style="text-align: center;color: black;" title="123">123</h2>
                <div class="info">

                    <span class="center1"><i class="fa fa-user"></i>12</span>
                    <span class="center1"><i
                            class="fa fa-calendar"></i>123</span>
                    <span class="center1"><i class="fa fa-eye"></i>123</span>
                    <span class="center1"><i class="fa fa-comments-o"></i>123&nbsp;comments</span>
                </div>
                <p style="margin-top: 20px;">
                    <blockquote>
                <p><strong>序言：本文主要介绍了laravel的扩展包laravel-voyager的使用，使用它快速搭建web后台管理系统，让你几乎不用敲一行代码就能实现？不信？那就跟着我的步伐浪起来…</strong>
                </p>
                </blockquote>
                <p>（1）. 安装laravel-voyager的扩展包</p>
                <p><code>composer require tcg/voyager</code></p>
                <p>（2）. 安装voyager配置和数据表，请提前在.env中配置好你的数据库连接，用一个新的数据库最好。</p>
                <p><code>php artisan voyager:install</code></p>
                <blockquote>
                    <p><strong>这条命令会产生voyager的配置项和一系列数据表</strong></p>
                </blockquote>
                <p>（3）. 创建管理员账号</p>
                <p><code>php artisan voyager:admin <a href="mailto:your@email.com">your@email.com</a> --create</code></p>
                <blockquote>
                    <p><strong>将上面的邮箱换成你自己的邮箱，接着它会叫你设置用户名和密码，你按着自己的设置就行了。</strong></p>
                </blockquote>
                <p>（4）. 因为是外国人的作品，所以后台是全英文的，不过好在有很多贡献者，所以voyager支持中文，按着官方文档，我设置了多语言支持，然而并没有看到中文，所以我简单粗暴的把它的中文包和默认英文包对换了，你也可以这样做。记得做好备份。</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/24f11b8fe5b7e8373e8bbb56ca1a63a0.png" alt="对换语言包"></a>
                </p><p>（5）. voyager已经内置了RBAC权限体系，你可以按着自己的业务需求进行更改，强大的BREAD Builder能让你不敲代码完成CURD操作，这真是so cool! <a class='hx_article_images' href="javascript:void(0);"><a class='hx_article_images' href="javascript:void(0);"><a class='hx_article_images' href="javascript:void(0);"></a></a></a>,下面以一个小例子来演示强大的BREAD Builder。
            </p><p>（6）. 点击数据库操作菜单，点击左上角的创建新表</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/e4ef568fd2b502318e7b905b7f307afd.png" alt="创建新表"></a>
                </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/91ab93e44e23b6fb7bdd05dd14e37bb0.png" alt="新数据表"></a>
            </p><p>（7）. 为创建好的articles表添加CURD操作</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/de0cf66a2e83ba3b8e2d6fa101f2f504.png" alt="curd"></a>
                </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/bf3232574264be7d563a49be013ec6e4.png" alt="设置"></a>
            </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/484298df90c451808f03fad052f809ca.png" alt="选项"></a>
            </p><p>（8）. 点击发布后，会自动创建对应的控制器CURD方法和对应的视图，但是现在你还看不了，因为没有授权，所以要给你的角色对应的权限</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/6ee08340672ddd7d7c5a2a99162ab8c6.png" alt="role"></a>
                </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/87a47c0c59541d9608416750cb715a36.png" alt="roles"></a>
            </p><p>（9）. 进入文章列表</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/927b32b965d628b20c464d973e44bf6d.png" alt="list"></a>
                </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/5c454e13b58443085c193dfbc004c67d.png" alt="list2"></a>
            </p><p>（10）. 新建一篇文章，然后查看你的数据表是不是多了一条记录</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/4bf9ce294ef4aea09d972e4db0d762ca.png" alt=""></a>
                </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/a5dc0f50fb901b0cdeb0a300306d133b.png" alt=""></a>
            </p><blockquote>
                <p><strong>是不是相当的强大，你并没有写一行代码，就完成了文章的添加。非常适合懒人..</strong> <a class='hx_article_images' href="javascript:void(0);"><a class='hx_article_images' href="javascript:void(0);"><a class='hx_article_images' href="javascript:void(0);"></a></a></a>
                </p></blockquote>
                <p>（11）. 将文章列表添加到菜单栏</p>
                <p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/7dc755a6b74622f2bb602ee9c1b08cb2.png" alt="1"></a>
                </p><p><a class='hx_article_images' href="javascript:void(0);"><img src="/upload/article/20171021/5b853a8dff308bbbdbf2725e85958d4c.png" alt="2"></a>
            </p><p>
                <a class='hx_article_images' href='javascript:void(0)'><img src="/upload/article/20171021/e773399f377367482aa2cb1d19c3dd33.png" alt="3"></a>
            </p><p>
                <a class='hx_article_images' href='javascript:void(0)'><img src="/upload/article/20171021/ff72191b9a4e8f5a30ac4f23656564aa.png" alt="4"></a></p>
                <p>
                    <a class='hx_article_images' href='javascript:void(0)'><img src="/upload/article/20171021/cc0a6dbb29aa592b593a2c03bf872012.png" alt="5"></a>
                </p><blockquote>
                <p><strong>到这里就完成了基本的操作，除此之外，voyager还有很多强大的地方，建议你去官网仔细看文档，如果你不喜欢自带的页面，完全可以重构，官方文档有重构具体方法，其实就是覆盖原来的模板而已<br>扩展地址：<a href="https://github.com/the-control-group/voyager" target="_blank">voyager</a></strong></p>
                <p>转载请注明出处，谢谢！
                </p></blockquote>
                </p>
                <br/>
                <ul class="b-prev-next">
                    <li class="b-prev">
                        <span>没有了</span>
                        <a href="" style="color: black">234</a>
                    </li>
                    <li class="b-next">
                        下一篇：
                        <span>没有了</span>
                        <a href="" style="color: black">fsdfs</a>
                    </li>
                </ul>
                <div class="share">
                    <div class="bdsharebuttonbox"><a href="#" class="bds_more" data-cmd="more">
                    </a><a href="#" class="bds_qzone" data-cmd="qzone" title="分享到QQ空间"></a>
                        <a href="#" class="bds_tsina" data-cmd="tsina" title="分享到新浪微博"></a>
                        <a href="#" class="bds_sqq" data-cmd="sqq" title="分享到QQ好友"></a>
                        <a href="#" class="bds_renren" data-cmd="renren" title="分享到人人网"></a>
                        <a href="#" class="bds_weixin" data-cmd="weixin" title="分享到微信"></a>
                        <a href="#" class="bds_twi" data-cmd="twi" title="分享到Twitter"></a>
                    </div>
                </div>
                <br/>
                <h6>标签</h6>
                <div class="vid-tags">
                    <a href="">123123</a>
                    <a href="">123123</a>
                </div>
                <div class="line"></div>
                <div class="comment a_comment">
                    <h5>评论一下</h5>
                    <div>
                        <div class="row">
                            <div class="col-md-6">
                                <div class="form-group">
                                    <input type="email" class="form-control input-lg" name="email" id="email"
                                           placeholder="接收回复的邮箱" required="required"/>
                                </div>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-md-12">
                                <div class="form-group">
                                    <textarea name="message" id="message" class="form-control" rows="4" cols="25"
                                              required="required" placeholder="评论内容"></textarea>
                                </div>
                                <button type="submit" class="btn btn-4 btn-block save" data="123">
                                    立即评论
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
                <br/>
                <h5>评论列表</h5>
                <hr/>
                <div class="widget wid-comment hx-comments">
                    <div class="content">
                        <div class="post">
                            <img src="" class="img-circle" title="123"/>
                            <div class="wrapper" style="max-width: 620px;display: inline-block;text-align: left;margin-left: 5px">
                                <ul class="list-inline">
                                    <li>123&emsp;评论:</li>
                                </ul>
                                <h5>123</h5>
                                <ul class="list-inline">
                                    <li><i class="fa fa-calendar"></i>123</li>
                                    <li style="cursor: pointer" class="c_reply"><i class="fa fa-comment"></i>回复</li>
                                </ul>
                            </div>
                            <div class="form-group hx_group" style="display: none">
                                <textarea id="co_content" name="co_content" class="form-control co_content" required="required" placeholder="回复内容"></textarea>
                                <input class="pull-left co_email" type="email" name="co_email" placeholder="请输入接收回复的邮箱" required="required"/>
                                <button type="submit" class="btn btn-success co_reply" data-pid="" data-aid="" data-rid="">回复</button>
                                <button type="submit" class="btn btn-success co_cancel">取消</button>
                            </div>
                            <div class="post hx_child_comments" style="margin-left: 55px;margin-bottom: 0; border-bottom: 1px solid #ddd;padding-top: 10px;">
                                <img src="" class="img-circle" title=""/>
                                <div class="wrapper child_wrapper" style="max-width: 580px;display: inline-block;text-align: left;margin-left: 5px">
                                    <ul class="list-inline">
                                        <li>234&nbsp;回复&nbsp;234</li>
                                    </ul>
                                    <h5>234</h5>
                                    <ul class="list-inline">
                                        <li><i class="fa fa-calendar"></i>234</li>
                                        <li style="cursor: pointer" class="hx_child_co_reply"><i class="fa fa-comments"></i>回复</li>
                                    </ul>
                                </div>
                                <div class="form-group hx_child_group" style="display: none">
                                    <textarea id="coo_content" name="co_content" class="form-control co_content" required="required" placeholder="回复内容"></textarea>
                                    <input class="pull-left co_email" type="email" name="co_email" placeholder="请输入接收回复的邮箱" required="required"/>
                                    <button type="submit" class="btn btn-success child_co_reply" data-pid="" data-aid="" data-rid="">回复</button>
                                    <button type="submit" class="btn btn-success child_co_cancel">取消</button>
                                </div>
                            </div>

                        </div>

                        <h6 style="text-align: center">暂时没有评论,快来评论吧..</h6>

                    </div>
                </div>

            </div>
        </div>

    </div>
</template>

<script>
    export default {
        name: "detail"
    }
</script>

<style scoped>
    .vid-tags a {
        background: #FD0005;
        padding: 5px 10px;
        color: #fff;
        margin: 2px;
    }
</style>
