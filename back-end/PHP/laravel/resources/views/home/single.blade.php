@php
    $agent = new \Jenssegers\Agent\Agent();
@endphp
@extends('home.layouts.home')
@section('meta')
    <title>{{config('webConfig.web_title')}}-{{$single->title ?? ''}}</title>
    <meta name="keywords" content="{{$single->keywords ?? ''}}">
    <meta name="description" content="{{$single->description ?? ''}}">
    @stop
@push('other-css')
    <style>
        code {
            padding: 2px 4px;
            font-size: 90%;
            color: #c7254e;
            background-color: #f9f2f4;
            border-radius: 4px;
            white-space:unset;
        }
    </style>
    @endpush
@section('content')
    <body>
    <div class="featured container">
        <h7>随机文章推荐</h7>
        <br/>
        <div id="owl-demo-2" class="owl-carousel">
            @foreach($randArticle as $ra)
                <div class="item">
                    <div class="zoom-container">
                        <div class="zoom-caption" title="{{$ra->title}}">
                            <span></span>
                            <a href="{{url('article/'.$ra->id.'.html')}}">
                                <i class="" style="color: #fff"></i>
                            </a>
                        </div>
                        <img src="{{$ra->thumb}}" style="width: 200px;height: 128px" title="{{$ra->title}}"/>
                    </div>
                </div>
            @endforeach
        </div>

    </div>
    <!-- Header -->
    @section('main-content')
        <div id="main-content" class="col-md-8">
            <div class="box single_article">
                <div class="line"></div>
                <h2 style="text-align: center;color: black;" title="{{$single->title}}">{{$single->title}}</h2>
                <div class="info">
                        <span class="center1"><i class="fa fa-user"></i>{{$single->author}}</span>
                        <span class="center1"><i class="fa fa-calendar"></i>{{$single->created_at}}</span>
                        <span class="center1"><i class="fa fa-eye"></i>{{$single->click}}</span>
                    @if($agent->isDesktop() || $agent->is('Windows'))
                    <span class="center1"><i class="fa fa-comments-o"></i>{{$single->comments_count}}&nbsp;comments</span>
                    @endif
                </div>
                <p style="margin-top: 20px;">{!! $single->content_html !!}</p>
                <br/>
                <ul class="b-prev-next">
                    <li class="b-prev">
                        上一篇：
                        @if(is_null($prev))
                            <span>没有了</span>
                        @else
                            <a href="{{ url('article/'.$prev->id.'.html') }}"
                               style="color: black">{{ $prev->title }}</a>
                        @endif

                    </li>
                    <li class="b-next">
                        下一篇：
                        @if(is_null($next))
                            <span>没有了</span>
                        @else
                            <a href="{{ url('article/'.$next->id.'.html') }}"
                               style="color: black">{{ $next->title }}</a>
                        @endif
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
                    <script>window._bd_share_config = {
                            "common": {
                                "bdSnsKey": {},
                                "bdText": "",
                                "bdMini": "2",
                                "bdMiniList": false,
                                "bdPic": "",
                                "bdStyle": "1",
                                "bdSize": "24"
                            }, "share": {}
                        };
                        with (document) 0[(getElementsByTagName('head')[0] || body).appendChild(createElement('script')).src = '{{asset("static/api/js/share.js")}}?v=89860593.js?'];</script>
                </div>
                <br/>
                <h6>标签</h6>
                <div class="vid-tags">
                    @foreach($single->tags as $tag)
                        <a href="{{url('tag/'.$tag->id.'.html')}}">{{$tag->name}}</a>
                    @endforeach
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
                                <button type="submit" class="btn btn-4 btn-block save" data="{{$single->id ?? ''}}">
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
                        @if($comment->isNotEmpty())
                            @foreach($comment as $com)
                                <div class="post">
                                    <img src="{{asset($com->user_avatar)}}" class="img-circle" title="{{$com->username}}"/>
                                    <div class="wrapper" style="max-width: 620px;display: inline-block;text-align: left;margin-left: 5px">
                                        <ul class="list-inline">
                                            <li>{{$com->username}}&emsp;评论:</li>
                                        </ul>
                                        <h5>{{$com->content ?? ''}}</h5>
                                        <ul class="list-inline">
                                            <li><i class="fa fa-calendar"></i>{{$com->created_at ?? ''}}</li>
                                            <li style="cursor: pointer" class="c_reply"><i class="fa fa-comment"></i>回复</li>
                                        </ul>
                                    </div>
                                    <div class="form-group hx_group" style="display: none">
                                        <textarea id="co_content" name="co_content" class="form-control co_content" required="required" placeholder="回复内容"></textarea>
                                        <input class="pull-left co_email" type="email" name="co_email" placeholder="请输入接收回复的邮箱" required="required"/>
                                        <button type="submit" class="btn btn-success co_reply" data-pid="{{$com->id ?? ''}}" data-aid="{{$single->id ?? ''}}" data-rid="{{$com->id ?? ''}}">回复</button>
                                        <button type="submit" class="btn btn-success co_cancel">取消</button>
                                    </div>
                                    @if($com->children->isNotEmpty())
                                        @foreach($com->children as $childComment)
                                            <div class="post hx_child_comments" style="margin-left: 55px;margin-bottom: 0; border-bottom: 1px solid #ddd;padding-top: 10px;">
                                                <img src="{{asset($childComment->user_avatar)}}" class="img-circle" title="{{$childComment->username}}"/>
                                                <div class="wrapper child_wrapper" style="max-width: 580px;display: inline-block;text-align: left;margin-left: 5px">
                                                    <ul class="list-inline">
                                                        <li>{{$childComment->username}}&nbsp;回复&nbsp;{{$childComment->reply_name}}</li>
                                                    </ul>
                                                    <h5>{{$childComment->content}}</h5>
                                                    <ul class="list-inline">
                                                        <li><i class="fa fa-calendar"></i>{{$childComment->created_at}}</li>
                                                        <li style="cursor: pointer" class="hx_child_co_reply"><i class="fa fa-comments"></i>回复</li>
                                                    </ul>
                                                </div>
                                                <div class="form-group hx_child_group" style="display: none">
                                                    <textarea id="co_content" name="co_content" class="form-control co_content" required="required" placeholder="回复内容"></textarea>
                                                    <input class="pull-left co_email" type="email" name="co_email" placeholder="请输入接收回复的邮箱" required="required"/>
                                                    <button type="submit" class="btn btn-success child_co_reply" data-pid="{{$com->id or ''}}" data-aid="{{$single->id or ''}}" data-rid="{{$childComment->id or ''}}">回复</button>
                                                    <button type="submit" class="btn btn-success child_co_cancel">取消</button>
                                                </div>
                                            </div>
                                        @endforeach
                                    @endif
                                </div>


                            @endforeach
                            @else
                        <h6 style="text-align: center">暂时没有评论,快来评论吧..</h6>
                        @endif
                    </div>
                </div>

            </div>
        </div>
    @endsection
    <!-- JS -->
    @endsection
    @push('scripts')
        <script>
            var user_id = "{{session('user.id')}}";
            var user_name = "{{session('user.name')}}";
        </script>
        <script src="{{asset('js/home/comment.js').'?ver='.RELEASE_VERSION}}"></script>
        <script src="{{asset('js/jqueryTpfd/js/postbird-img-glass.js').'?ver='.RELEASE_VERSION}}"></script>
        <script>
            PostbirdImgGlass.init({
                domSelector:".hx_article_images img",
                animation:true
            });
        </script>

    @endpush
