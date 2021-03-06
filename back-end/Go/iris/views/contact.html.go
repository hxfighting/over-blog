// Code generated by hero.
// source: /Users/huxin/DockerVolumes/Nginx/www/over-blog/back-end/Go/iris/views/contact.html
// DO NOT EDIT!
package template

import (
	"blog/helper"
	"bytes"

	"github.com/shiyanhui/hero"
)

func ContactPage(AuthInfo map[string]string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <link rel="shortcut icon" type="image/ico" href="/image/favicon.ico">
    <meta name="viewport" content="width=device-width, initial-scale=1"/>
    <meta name="csrf-token" content="`)
	buffer.WriteString(CsrfToken)
	buffer.WriteString(`">
    `)
	buffer.WriteString(`
<title>`)
	hero.EscapeHTML(Config["web_title"], buffer)
	buffer.WriteString(`</title>
<meta name="keywords" content='`)
	hero.EscapeHTML(Config["keywords"], buffer)
	buffer.WriteString(`'>
<meta name="description" content='`)
	hero.EscapeHTML(Config["description"], buffer)
	buffer.WriteString(`'>
<meta name="author" content='`)
	hero.EscapeHTML(Config["seo_title"], buffer)
	buffer.WriteString(`'>
`)

	buffer.WriteString(`
    <!-- Bootstrap Core CSS -->
    <link href="/css/home/bootstrap.min.css" rel="stylesheet">

    <!-- Owl Carousel Assets -->
    <link href="/css/home/owl-carousel/owl.carousel.css" rel="stylesheet">
    <link href="/css/home/owl-carousel/owl.theme.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link rel="stylesheet" href="/css/home/style.css">
    <link href="/css/home/bootstrap-datetimepicker.min.css" rel="stylesheet" media="screen">

    <!-- Custom Fonts -->
    <link rel="stylesheet" href="/css/home/font-awesome-4.4.0/css/font-awesome.min.css" type="text/css">

    <!-- jQuery and Modernizr-->
    <script src="/js/jquery-3.3.1.min.js"></script>

    <!-- Core JavaScript Files -->
    <script src="/js/bootstrap.min.js"></script>

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <![endif]-->
    <link rel="stylesheet" href="/css/home/animate.css">
    <link rel="stylesheet" href="/css/home/index.css">
    <link rel="stylesheet" href="/css/sweetalert2.min.css">
    `)
	buffer.WriteString(Config["baidu_tongji"])
	buffer.WriteString(`

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
                        <li><a title="首页" href=""><i class="fa fa-home"></i> 首页</a></li>
                        `)
	if len(AuthInfo) <= 0 {
		buffer.WriteString(`
                        <li style="cursor: pointer"><a title="登录" class="oauth_login"><i class="fa fa-user"></i> 登录</a>
                        </li>
                        `)
	} else {
		buffer.WriteString(`
                        <li style="cursor: pointer"><a title="退出登录" class="oauth_quit"><i class="fa fa-user"></i>
                                退出登录</a></li>
                        `)
	}
	buffer.WriteString(`
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
                                        src="/image/sina-login.png" alt="微博登录" title="微博登录"></a>
                        </li>
                        <li class="col-xs-6 col-md-4 col-lg-4 b-login-img">
                            <a href="/oauth/redirectToProvider/qq"><img
                                        src="/image/qq-login.png" alt="QQ登录" title="QQ登录"></a>
                        </li>
                        <li class="col-xs-6 col-md-4 col-lg-4 b-login-img">
                            <img src="/image/wechat.png" alt="微信登录" title="微信登录" style="cursor: pointer"
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
            `)
	if len(Category) > 0 {
		buffer.WriteString(`
            <ul class="nav navbar-nav">

                `)
		for _, ca := range Category {
			if len(ca.Children) > 0 {
				buffer.WriteString(`
                <li class="dropdown">
                    <a href="`)
				hero.EscapeHTML(ca.Url, buffer)
				buffer.WriteString(`" class="dropdown-toggle"
                       title="`)
				hero.EscapeHTML(*ca.Title, buffer)
				buffer.WriteString(`">`)
				hero.EscapeHTML(*ca.Title, buffer)
				buffer.WriteString(`</a>
                    <div class="dropdown-menu">
                        <div class="dropdown-inner">
                            <ul class="list-unstyled">
                                `)
				for _, child := range ca.Children {
					buffer.WriteString(`
                                <li><a href="`)
					hero.EscapeHTML(child.Url, buffer)
					buffer.WriteString(`">`)
					hero.EscapeHTML(*child.Title, buffer)
					buffer.WriteString(`</a></li>
                                `)
				}
				buffer.WriteString(`
                            </ul>
                        </div>
                    </div>
                </li>
                `)
			} else {
				buffer.WriteString(`
                <li>
                    <a href="`)
				hero.EscapeHTML(ca.Url, buffer)
				buffer.WriteString(`" title="`)
				hero.EscapeHTML(*ca.Title, buffer)
				buffer.WriteString(`">`)
				hero.EscapeHTML(*ca.Title, buffer)
				buffer.WriteString(`</a>
                </li>
                `)
			}
		}
		buffer.WriteString(`
            </ul>

            `)
	}
	if len(SocialData) > 0 {
		buffer.WriteString(`
            <ul class="list-inline navbar-right top-social">
                `)
		for _, so := range SocialData {
			buffer.WriteString(`
                <li><a title="博主`)
			hero.EscapeHTML(*so.Title, buffer)
			buffer.WriteString(`" href="`)
			hero.EscapeHTML(*so.Val, buffer)
			buffer.WriteString(`"
                       target="_blank"><i class="fa fa-`)
			hero.EscapeHTML(*so.Name, buffer)
			buffer.WriteString(`"></i></a></li>
                `)
		}
		buffer.WriteString(`
            </ul>
            `)
	}
	buffer.WriteString(`
        </div>
    </nav>
</header>
<body>
`)
	buffer.WriteString(`
<div id="page-content" class="archive-page container">
    <div class="">
        <div class="row">
            `)
	buffer.WriteString(`
<div id="main-content" class="col-md-8">
    <div class="box">
        <center>
            <div class="box-header">
                <h1 class="center">留言</h1>
            </div>
        </center>
        <div class="box-content">
            <div id="contact_form" class="my_contact">
                <div id="ff">
                    <label>
                        <span>请输入你的名字:</span>
                        <input type="text" name="name" id="name" required placeholder="任意名字都可以哦！">
                    </label>
                    <label>
                        <span>请输入你的邮箱:</span>
                        <input type="email" name="email" id="email" required placeholder="请输入正确邮箱,错误邮箱收不到回复哦！">
                    </label>
                    <label>
                        <span>在这里写下你想说的话:</span>
                        <textarea name="content" id="message"
                                  placeholder="说点你对这个家伙想说的话！,如果是交换友链，请写下你的友链地址和友链名字。"></textarea>
                    </label>
                    <center><input class="sendButton" type="submit" name="Submit" value="确定" onclick="contact()">
                    </center>
                </div>
            </div>
        </div>
    </div>
</div>
`)

	buffer.WriteString(`
            <div id="sidebar" class="col-md-4">
                <!---- Start Widget ---->
                <div class="widget wid-tags">
                    <div class="heading"><h4>搜索</h4></div>
                    <br/>
                    <div class="content">
                        <form role="form" class="form-horizontal" method="get" action="/search">
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
                        `)
	if len(Tag) > 0 {
		for _, t := range Tag {
			buffer.WriteString(`
                        <a class="label" href="/tag/`)
			hero.FormatInt(int64(*t.ID), buffer)
			buffer.WriteString(`">`)
			hero.EscapeHTML(t.Name, buffer)
			buffer.WriteString(`</a>
                        `)
		}
	} else {
		buffer.WriteString(`
                        <span>暂无标签</span>
                        `)
	}
	buffer.WriteString(`
                    </div>
                </div>
                <div class="widget ">
                    <div class="heading"><h4>最热文章</h4></div>
                    <br/>
                    <div class="content">
                        `)
	if len(HotArticle) > 0 {
		for _, article := range HotArticle {
			buffer.WriteString(`
                        <div class="wrap-vid">
                            <h3 class="vid-name"><a href='/article/`)
			buffer.WriteString(article["id"])
			buffer.WriteString(`'
                                                    title='`)
			hero.EscapeHTML(article["title"], buffer)
			buffer.WriteString(`'>
                                    `)

			title := []rune(article["title"])
			var article_title string
			if len(title) > 20 {
				article_title = string(title[0:18]) + "..."
			} else {
				article_title = string(title[:])
			}
			buffer.WriteString(article_title)

			buffer.WriteString(`
                                </a>
                            </h3>
                            <div class="info">
                                <span><i class="fa fa-calendar"></i>`)
			buffer.WriteString(article["created_at"])
			buffer.WriteString(`</span>
                                <span><i class="fa fa-eye"></i>`)
			buffer.WriteString(article["click"])
			buffer.WriteString(`</span>
                                <span><i class="fa fa-comments-o"></i>`)
			buffer.WriteString(article["comment_count"])
			buffer.WriteString(`&nbsp;Comments</span>
                            </div>
                        </div>
                        `)
		}
	} else {
		buffer.WriteString(`
                        <span>暂无文章</span>
                        `)
	}
	buffer.WriteString(`
                    </div>
                </div>
                <!---- Start Widget ---->
                <div class="widget wid-comment">
                    <div class="heading"><h4>最新评论</h4></div>
                    <br/>
                    <div class="content">
                        `)
	if len(RecentComment) <= 0 {
		buffer.WriteString(`
                        <h6 class="text-center">暂时没有评论</h6>
                        `)
	} else {
		for _, comment := range RecentComment {
			buffer.WriteString(`
                        <div class="post">
                            <a> <img src="`)
			buffer.WriteString(comment.Avatar)
			buffer.WriteString(`" class="img-circle img-responsive"
                                     title="`)
			buffer.WriteString(comment.Name)
			buffer.WriteString(`"/></a>
                            <div class="wrapper">
                                <a href="/article/`)
			hero.FormatInt(int64(comment.ArticleID), buffer)
			buffer.WriteString(`"
                                   title="`)
			hero.EscapeHTML(comment.Content, buffer)
			buffer.WriteString(`"><span>
                                         `)

			content_rune := []rune(comment.Content)
			var content string
			if len(content_rune) > 20 {
				content = string(content_rune[0:18]) + "..."
			} else {
				content = string(content_rune[:])
			}
			hero.EscapeHTML(content, buffer)

			buffer.WriteString(`
                                    </span></a>
                                <ul class="list-inline">
                                    <li><i class="fa fa-calendar"></i>&nbsp;`)
			buffer.WriteString(helper.GetDateTime(comment.CreatedAt, helper.YMD))
			buffer.WriteString(`</li>
                                </ul>
                            </div>
                        </div>
                        `)
		}
	}
	buffer.WriteString(`
                    </div>
                </div>
                <div class="widget wid-tags">
                    <div class="heading"><h4>友情链接</h4></div>
                    <p><a style="color: black;cursor: pointer" class="link_modal">申请友链</a></p>
                    <div class="content">
                        `)
	if len(Link) > 0 {
		for _, l := range Link {
			buffer.WriteString(`
                        <a class="label" href="`)
			hero.EscapeHTML(l.Url, buffer)
			buffer.WriteString(`" style="color: black"
                           title="`)
			hero.EscapeHTML(l.Description, buffer)
			buffer.WriteString(`" target="_blank">`)
			hero.EscapeHTML(l.Name, buffer)
			buffer.WriteString(`</a>
                        `)
		}
	}
	buffer.WriteString(`
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
                        `)
	if FooterData["famous_remark"] != nil {
		for _, fr := range FooterData["famous_remark"] {
			buffer.WriteString(`
                        <p>`)
			hero.EscapeHTML(fr.Val, buffer)
			buffer.WriteString(`</p>
                        `)
		}
	}
	buffer.WriteString(`
                    </div>
                </div>
                <div class="col-md-4 col-footer footer-3">
                    <div class="footer-heading"><h4>博客相关</h4></div>
                    <br/>
                    <div class="content">
                        <ul>
                            <li><i class="fa fa-cogs"></i>&nbsp;框架：`)
	hero.EscapeHTML(FooterData["blog_related"][8].Val, buffer)
	buffer.WriteString(`</li>
                            <li><i class="fa fa-list-ol"></i>&nbsp;文章总计：&nbsp;`)
	buffer.WriteString(FooterData["blog_article_count"][25].Val)
	buffer.WriteString(`&nbsp;篇</li>
                            <li>
                                <i class="fa fa-group"></i>&nbsp;访问总计：&nbsp;`)
	buffer.WriteString(FooterData["blog_view_count"][24].Val)
	buffer.WriteString(`
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
                            `)
	for _, bc := range FooterData["blogger_collection"] {
		buffer.WriteString(`
                            <li><a href="`)
		hero.EscapeHTML(bc.Val, buffer)
		buffer.WriteString(`" target="_blank" title="`)
		hero.EscapeHTML(bc.Title, buffer)
		buffer.WriteString(`"
                                   style="color: whitesmoke">`)
		hero.EscapeHTML(bc.Title, buffer)
		buffer.WriteString(`</a>
                            </li>
                            `)
	}
	buffer.WriteString(`

                            <li><a></a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    </div>

    <div class="copy-right">
        <p>`)
	buffer.WriteString(FooterData["copyright"][17].Val)
	buffer.WriteString(`</p>
    </div>
</footer>
<script src="/css/home/owl-carousel/owl.carousel.js"></script>
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
    let APP_URL = "`)
	hero.EscapeHTML(AppUrl, buffer)
	buffer.WriteString(`"
</script>
`)
	buffer.WriteString(Config["baidu_tuisong"])
	buffer.WriteString(`
<script src="/js/layer/layer.js"></script>
<script src="/js/sweetalert2.all.min.js"></script>
<script src="/js/common.js"></script>
<script src="/js/home/jquery.goup.min.js"></script>
<script src="/js/home/a_index.js"></script>
`)
	buffer.WriteString(`
<script>
    function contact() {
        var btn = $(this);
        if (btn.hasClass('disabled')) {
            Swal({
                position: 'center',
                type: 'error',
                title: '正在处理,请勿重复点击!',
                showConfirmButton: false,
                timer: 1500
            })
            return false;
        }
        let name = $("input[name=name]").val(),
            email = $("input[name=email]").val(),
            content = $('textarea[name=content]').val();
        if (name.length < 1) {
            Swal({
                position: 'center',
                type: 'error',
                title: '请输入正确的名字！',
                showConfirmButton: false,
                timer: 1500
            })
            return false;
        }
        if (email.length < 1) {
            Swal({
                position: 'center',
                type: 'error',
                title: '请输入正确的邮箱！',
                showConfirmButton: false,
                timer: 1500
            })
            return false;
        }
        var reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$/;
        if (!reg.test(email)) {
            Swal({
                position: 'center',
                type: 'error',
                title: '请输入正确的邮箱！',
                showConfirmButton: false,
                timer: 1500
            })
            return false;
        }
        if (content.length < 1) {
            Swal({
                position: 'center',
                type: 'error',
                title: '随便写点什么内容吧！',
                showConfirmButton: false,
                timer: 1500
            })
            return false;
        }
        btn.addClass('disabled');
        $.ajax({
            url: '/contact',
            type: 'POST',
            headers: {
                'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
            },
            data: JSON.stringify({
                name: name,
                email: email,
                content: content
            }),
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            success: function (res) {
                btn.removeClass('disabled');
                if (res.code == 200) {
                    Swal({
                        position: 'center',
                        type: 'success',
                        title: res.msg,
                        showConfirmButton: false,
                        timer: 1500
                    })
                    setTimeout(function () {
                        window.location.href = window.location.href;
                    }, 1600);
                } else {
                    Swal({
                        position: 'center',
                        type: 'error',
                        title: res.msg,
                        showConfirmButton: false,
                        timer: 1500
                    })
                }
            },
            error: function (res) {
                console.log(res.responseText);
            }
        });
    }
</script>
`)

	buffer.WriteString(`
</body>
</html>
`)

}
