@php
$agent = new \Jenssegers\Agent\Agent();
@endphp
@extends('home.layouts.home')
@push('other-css')
    <link rel="stylesheet" href="{{asset('css/home/hx.css')}}">
@endpush
@section('content')
<body>
<div class="featured container">
    <div class="row">
        @if($agent->is('Windows') || $agent->isDesktop())
        <div class="col-sm-8">
            <!-- Carousel -->
            <div id="carousel-example-generic" class="carousel slide" data-ride="carousel">
                <!-- Indicators -->
                <ol class="carousel-indicators">
                    <li data-target="#carousel-example-generic" data-slide-to="0" class="active"></li>
                    <li data-target="#carousel-example-generic" data-slide-to="1"></li>
                    <li data-target="#carousel-example-generic" data-slide-to="2"></li>
                </ol>
                <!-- Wrapper for slides -->

                <div class="carousel-inner">
                @foreach($rotation as $v)
                    <div class="item @if($v->id==1) active @endif">
                        <img src="{{$v->image_url}}" alt="拖油瓶">
                        <!-- Static Header -->
                        <div class="header-text hidden-xs">
                            <div class="col-md-12 text-center">
                                <h2 style="opacity: 0">&emsp;&emsp;&emsp;</h2>
                                <h2 style="opacity: 0">&emsp;&emsp;&emsp;</h2>
                                <br/>
                                <br>
                                <br/>
                                <br>
                                <h3>{{$v->image->words ?? ''}}</h3>
                                <br>
                            </div>
                        </div><!-- /header-text -->
                    </div>
                    @endforeach
                </div>

                <!-- Controls -->
                <a class="left carousel-control" href="#carousel-example-generic" data-slide="prev">
                    <span class="glyphicon glyphicon-chevron-left"></span>
                </a>
                <a class="right carousel-control" href="#carousel-example-generic" data-slide="next">
                    <span class="glyphicon glyphicon-chevron-right"></span>
                </a>
            </div><!-- /carousel -->
        </div>
        @endif
        <div class="col-sm-4" style="text-align: center">
            <div id="owl-demo-1" class="owl-carousel">
                @foreach($photo as $v)
                <img src="{{$v->image_url}}" alt="博主照片"/>
                    @endforeach
            </div>
            <br/>
            {{--<img src="{{asset('resources/views/home/images/banner.jpg')}}" />--}}
            {{--js控制显示时间--}}
            <h5 title="距离毕业">
                <div >
                    <p class="center">距离结婚还有：</p>
                    <span id="clock" style="font-size: 18px"></span>
                </div>
            </h5>
        </div>
    </div>
</div>

<!-- /////////////////////////////////////////Content -->
@section('main-content')
            <div id="main-content" class="col-md-8">
                <div id="b-content" class="container">
                    <div class="row">
                        <div class="col-xs-12 col-md-12 col-lg-8">
                            @foreach($newArticle as $new)
                                <div class="row b-one-article">
                                    <h3 class="col-xs-12 col-md-12 col-lg-12">
                                        <a class="b-oa-title" href="{{url('article/'.$new->id)}}"
                                           style="color: black;" title="{{$new->title}}">{{mb_strimwidth($new->title,0,30,'...','utf8')}}</a>
                                        @if($new->is_top==1)
                                        <label class="btn btn-danger pull-right inline" style="font-size: 10px;">置顶</label>
                                            @endif
                                    </h3>
                                    <div class="col-xs-12 col-md-12 col-lg-12 b-date">
                                        <ul class="row">
                                            <li class="col-xs-5 col-md-2 col-lg-2">
                                                <i class="fa fa-user"></i>&nbsp;{{$new->author}}
                                            </li>
                                            <li class="col-xs-7 col-md-3 col-lg-3">
                                                <i class="fa fa-calendar"></i>&nbsp;{{$new->created_at}}
                                            </li>
                                            <li class="col-xs-5 col-md-2 col-lg-2">
                                                <i class="fa fa-eye"></i>&nbsp;{{$new->click}}

                                            </li>
                                            <li class="col-xs-7 col-md-5 col-lg-4 "><i class="fa fa-tags"></i>
                                                @foreach($new->tags as $tags)
                                                    <a class="b-tag-name"
                                                       style="color: black">{{$tags->name}}</a>
                                                @endforeach
                                            </li>
                                        </ul>
                                    </div>
                                    <div class="col-xs-12 col-md-12 col-lg-12">
                                        <div class="row">
                                            <!-- 文章封面图片开始 -->
                                            <div class="col-sm-6 col-md-6 col-lg-4 hidden-xs">
                                                <figure class="b-oa-pic b-style1">
                                                    <a href="{{url('article/'.$new->id)}}">
                                                        <img src="{{$new->thumb}}" alt="{{$new->title}}" title="{{$new->title}}">
                                                    </a>
                                                    <figcaption>
                                                        <a href="{{url('article/'.$new->id)}}"></a>
                                                    </figcaption>
                                                </figure>
                                            </div>
                                            <!-- 文章封面图片结束 -->

                                            <!-- 文章描述开始 -->
                                            <div class="col-xs-12 col-sm-6  col-md-6 col-lg-8 b-des-read">
                                                {{$new->description}}
                                            </div>
                                            <!-- 文章描述结束 -->
                                        </div>
                                    </div>
                                    <a class=" b-readall" href="{{url('article/'.$new->id)}}">阅读全文</a>
                                </div>
                            @endforeach
                            <div class="a_page" style="text-align: center">
                                {{$newArticle->links()}}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
@endsection
<script type="text/javascript" src="{{asset('js/jquery.countdown.min.js')}}"></script>
<script type="text/javascript">
    $(function () {
        $('#clock').countdown('2024/5/20', function (event) {
            var $this = $(this).html(event.strftime(''
                + '<span>%w</span> 星期 '
                + '<span>%d</span> 天 '
                + '<span>%H</span> 小时 '
                + '<span>%M</span> 分钟 '
                + '<span>%S</span> 秒'));
        });

    })
</script>
@endsection
