<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <link rel="shortcut icon" type="image/ico" href="image/favicon.ico">
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    {use class="yii\helpers\Html"}
    {Html::csrfMetaTags()}
    <title>websiteConfig.web_title</title>
    <meta name="keywords" content="websiteConfig.keywords">
    <meta name="description" content="'websiteConfig.description">
    <meta name="author" content="websiteConfig.seo_title">
    {* Bootstrap Core CSS *}
    <link href="css/home/bootstrap.min.css" rel="stylesheet">

    <!-- Owl Carousel Assets -->
    <link href="css/home/owl-carousel/owl.carousel.css" rel="stylesheet">
    <link href="css/home/owl-carousel/owl.theme.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link rel="stylesheet" href="css/home/style.css">
    <link href="css/home/bootstrap-datetimepicker.min.css" rel="stylesheet" media="screen">

    <!-- Custom Fonts -->
    <link rel="stylesheet" href="css/home/font-awesome-4.4.0/css/font-awesome.min.css" type="text/css">

    <!-- jQuery and Modernizr-->
    <script src="js/jquery-3.3.1.min.js"></script>

    <!-- Core JavaScript Files -->
    <script src="js/bootstrap.min.js"></script>

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <![endif]-->
    <link rel="stylesheet" href="css/home/animate.css">
    <link rel="stylesheet" href="css/home/index.css">
    <link rel="stylesheet" href="css/sweetalert2.min.css">

</head>
<body>
<header>
    <!--Top-->
    <nav id="top">
        <div class="container">
            <div class="row">
                <div class="col-md-6" title="Welcome">
                    <strong>Welcome !</strong>
                </div>
                <div class="col-md-6">
                    <ul class="list-inline top-link link">
                        @if(session('user.is_admin')==0 || is_null(session('user')))
                        <li><a title="首页" href="/"><i class="fa fa-home"></i> 首页</a></li>
                        @else
                        <li style="cursor: pointer"><a title="后台" href="blog.blog_admin_url"
                                                       target="_blank"><i
                                        class="fa fa-home"></i> 后台</a></li>
                        @endif
                        @guest
                        <li style="cursor: pointer"><a title="登录" class="oauth_login"><i class="fa fa-user"></i> 登录</a>
                        </li>
                        @endguest
                        @auth
                        <li style="cursor: pointer"><a title="退出登录" class="oauth_quit"><i class="fa fa-user"></i>
                                退出登录</a></li>
                        @endauth
                        <li><a title="联系我" href="/contact"><i class="fa fa-comments"></i>↓联系我↓</a></li>
                    </ul>
                </div>
            </div>
        </div>
    </nav>

    <div class="modal fade" id="b-modal-login" tabindex="-1" role="dialog" aria-labelledby="myModalLabel"
         aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content row">
                <div class="col-xs-12 col-md-12 col-lg-12">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                    aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title b-ta-center" id="myModalLabel">无需注册，用以下帐号即可直接登录</h4>
                    </div>
                </div>
                <div class="col-xs-12 col-md-12 col-lg-12 b-login-row">
                    <ul class="row">
                        <li class="col-xs-6 col-md-4 col-lg-4 b-login-img">
                            <a href="/oauth/redirectToProvider/weibo"><img
                                        src="image/sina-login.png" alt="微博登录" title="微博登录"></a>
                        </li>
                        <li class="col-xs-6 col-md-4 col-lg-4 b-login-img">
                            <a href="/oauth/redirectToProvider/qq"><img
                                        src="image/qq-login.png" alt="QQ登录" title="QQ登录"></a>
                        </li>
                        <li class="col-xs-6 col-md-4 col-lg-4 b-login-img">
                            <img src="image/wechat.png" alt="微信登录" title="微信登录" style="cursor: pointer"
                                 class="hx-wechat" tabindex="0" role="button" data-toggle="popover" data-trigger="focus"
                                 data-content="" data-placement="bottom" data-html="true">
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
    <div class="wechat-modal fade bs-example-modal-sm" tabindex="-1" role="dialog" aria-labelledby="mySmallModalLabel"
         hidden>
        <div class="modal-dialog modal-sm" role="document">
            <div class="modal-content">
                <img class="wechat-image" src="">
            </div>
        </div>
    </div>
    <!--Navigation-->
    <nav id="menu" class="navbar container">
        <div class="navbar-header">
            <button type="button" class="btn btn-navbar navbar-toggle" data-toggle="collapse"
                    data-target=".navbar-ex1-collapse"><i class="fa fa-bars"></i></button>
            <a class="navbar-brand" href="/">
                <div class="logo" title="拖油瓶"><span>拖油瓶</span></div>
            </a>
        </div>
        <div class="collapse navbar-collapse navbar-ex1-collapse">
            <ul class="nav navbar-nav">
                {foreach $this->params['dh'] as $k => $v }
                    {if !empty($v.children) }
                        <li class="dropdown">
                            <a href="{$v.url}" class="dropdown-toggle"
                               title="{$v.title}">{$v.title}</a>
                            <div class="dropdown-menu">
                                <div class="dropdown-inner">
                                    <ul class="list-unstyled">
                                        {foreach $v.children as $child}
                                            <li><a href="{$child.url}">{$child.title}</a></li>
                                        {/foreach}
                                    </ul>
                                </div>
                            </div>
                        </li>
                    {else}
                        <li>
                            <a href="{$v.url}" title="{$v.title}">{$v.title}</a>
                        </li>
                    {/if}
                {/foreach}
            </ul>
            <ul class="list-inline navbar-right top-social">
                {foreach $this->params['socialData'] as $so}
                    <li><a title="博主" href="{$so.val}"
                           target="_blank"><i class="fa fa-{$so.name}"></i></a></li>
                {/foreach}
            </ul>
        </div>
    </nav>
</header>
<body>
@yield('content')
<div id="page-content" class="archive-page container">
    <div class="">
        <div class="row">
            @section('main-content')
            @show
            <div id="sidebar" class="col-md-4">
                <!---- Start Widget ---->
                <div class="widget wid-tags">
                    <div class="heading"><h4>搜索</h4></div>
                    <br/>
                    <div class="content">
                        <form role="form" class="form-horizontal" method="get" action="search">
                            <input type="text" placeholder="回车键搜索文章" value="" name="search"
                                   id="v_search" class="form-control">
                        </form>
                    </div>
                </div>
                <!---- Start Widget ---->
                <div class="widget wid-tags">
                    <div class="heading"><h4>标签云</h4></div>
                    <br/>
                    <div class="content">
                        {foreach $this->params['tagCloud'] as $tag}
                            <a class="label" href="/tag/{$tag.id}">{$tag.name}</a>
                        {/foreach}
                    </div>
                </div>
                <div class="widget ">
                    <div class="heading"><h4>最热文章</h4></div>
                    <br/>
                    <div class="content">
                        {foreach $this->params['hotArticle'] as $hot}
                            <div class="wrap-vid">
                                <h3 class="vid-name"><a href="/article/{$hot.id}"
                                                        title="{$hot.title}">{mb_strimwidth($hot.title,0,30,'...','utf8')}</a>
                                </h3>
                                <div class="info">
                                    <span><i class="fa fa-calendar"></i>{date('Y/m/d',intval($hot.created_at))}</span>
                                    <span><i class="fa fa-eye"></i>{$hot.click}</span>
                                    <span><i class="fa fa-comments-o"></i>{$hot.comment_count}&nbsp;Comments</span>
                                </div>
                            </div>
                        {/foreach}
                    </div>
                </div>
                <!---- Start Widget ---->
                <div class="widget wid-comment">
                    <div class="heading"><h4>最新评论</h4></div>
                    <br/>
                    <div class="content">
                        {if empty($this->params['comment_t']) }
                            <h6 class="text-center">暂时没有评论</h6>
                        {else}
                            {foreach $this->params['comment_t'] as $v}
                                <div class="post">
                                    <a> <img src="{$v.avatar}" class="img-circle img-responsive"
                                             title="{$v.name}"/></a>
                                    <div class="wrapper">
                                        <a href="/article/{$v.article_id}"
                                           title="{$v.content}"><span>{mb_strimwidth($v.content,0,20,'...','utf8')}</span></a>
                                        <ul class="list-inline">
                                            <li><i class="fa fa-calendar"></i>&nbsp;{date('Y/m/d',$v.created_at)}</li>
                                        </ul>
                                    </div>
                                </div>
                            {/foreach}
                        {/if}
                    </div>
                </div>
                <div class="widget wid-tags">
                    <div class="heading"><h4>友情链接</h4></div>
                    <p><a style="color: black;cursor: pointer" class="link_modal">申请友链</a></p>
                    <div class="content">
                        {if !empty($this->params['friendLink']) }
                            {foreach $this->params['friendLink'] as $link }
                                <a class="label" href="{$link.url}" style="color: black"
                                   title="{$link.description}" target="_blank">{$link.name}</a>
                            {/foreach}
                        {else}
                            <span>暂无友联数据</span>
                        {/if}
                    </div>
                </div>

                <!-- link --->
                <div class="modal fade" id="hx_link_modal" tabindex="-1" role="dialog"
                     aria-labelledby="exampleModalLabel">
                    <div class="modal-dialog" role="document">
                        <div class="modal-content">
                            <div class="modal-header">
                                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                            aria-hidden="true">&times;</span></button>
                                <h4 class="modal-title" id="exampleModalLabel">友链申请</h4>
                            </div>
                            <div class="modal-body">
                                <div>
                                    <div class="form-group">
                                        <label for="ur_name" class="control-label">友链名字:</label>
                                        <input type="text" class="form-control" id="ur_name" name="ur_name">
                                    </div>
                                    <div class="form-group">
                                        <label for="ur_url" class="control-label">友链URL:</label>
                                        <input type="text" class="form-control" id="ur_url" name="ur_url"
                                               value="http://">
                                    </div>
                                    <div class="form-group">
                                        <label for="ur_description" class="control-label">友链描述:</label>
                                        <textarea class="form-control" id="ur_description"
                                                  name="ur_description"></textarea>
                                    </div>
                                </div>
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                                <button type="button" class="btn btn-primary save">提交</button>
                            </div>
                        </div>
                    </div>
                </div>

            </div>
        </div>
    </div>
</div>
</body>
<footer>
    <div class="wrap-footer">
        <div class="container">
            <div class="row">
                <div class="col-md-4 col-footer footer-1">
                    <div class="footer-heading"><h1><span style="color: #fff;">拖油瓶</span></h1></div>
                    <br/>
                    <div class="content">
                        {if isset($this->params['footerData']['famous_remark']) && !empty($this->params['footerData']['famous_remark']) }
                            {foreach $this->params['footerData']['famous_remark'] as $v }
                                <p>{$v.val}</p>
                            {/foreach}
                        {else}
                            <p>暂无名言</p>
                        {/if}
                    </div>
                </div>
                <div class="col-md-4 col-footer footer-3">
                    <div class="footer-heading"><h4>博客相关</h4></div>
                    <br/>
                    <div class="content">
                        <ul>
                            <li>
                                <i class="fa fa-cogs"></i>&nbsp;框架：{$this->params['footerData']['blog_related'][0]['val']}
                            </li>
                            <li><i class="fa fa-list-ol"></i>&nbsp;文章总计：&nbsp;&nbsp;篇</li>
                            <li>
                                <i class="fa fa-group"></i>&nbsp;访问总计：&nbsp;
                                &nbsp;次
                            </li>
                            <li><i></i></li>
                        </ul>
                    </div>
                </div>
                <div class="col-md-4 col-footer footer-3">
                    <div class="footer-heading"><h4>博主长逛</h4></div>
                    <br/>
                    <div class="content">
                        <ul>
                            {if isset($this->params['footerData']['blogger_collection']) && !empty($this->params['footerData']['blogger_collection'])}
                                {foreach $this->params['footerData']['blogger_collection'] as $v}
                                    <li><a href="{$v.val}" target="_blank" title="{$v.title}"
                                           style="color: whitesmoke">{$v.title}</a>
                                    </li>
                                {/foreach}
                            {else}
                                <li>暂无长逛内容</li>
                            {/if}
                            <li><a></a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="copy-right">
        <p></p>
    </div>
</footer>
<script src="css/home/owl-carousel/owl.carousel.js"></script>
<script>
    $(document).ready(function () {
        $("#owl-demo-1").owlCarousel({
            autoPlay: 3000,
            items: 1,
            itemsDesktop: [1199, 1],
            itemsDesktopSmall: [400, 1]
        });
        $("#owl-demo-2").owlCarousel({
            autoPlay: 3000,
            items: 5,
            itemsDesktop: [1199, 4],
            itemsDesktopSmall: [979, 4]

        });
    });
    let APP_URL = "app.url"
</script>
<script src="js/layer/layer.js"></script>
<script src="js/sweetalert2.all.min.js"></script>
<script src="js/common.js"></script>
<script src="js/home/jquery.goup.min.js"></script>
<script src="js/home/a_index.js"></script>
@stack('scripts')
</body>
</html>
