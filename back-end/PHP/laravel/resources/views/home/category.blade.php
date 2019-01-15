@extends('home.layouts.home')
@section('content')

@section('main-content')
    <div id="main-content" class="col-md-8">
        <div id="b-content" class="container">
            <div class="row">
                <div class="col-xs-12 col-md-12 col-lg-8">
                    @if(!empty($tagName))
                        <div class="row b-tag-title">
                            <div class="col-xs-12 col-md-12 col-lg-12">
                                <h3>拥有<span class="b-highlight">{{ $tagName }}</span>标签的文章</h3>
                            </div>
                        </div>
                    @endif
                    @if(request()->has('searcharticle'))
                        <div class="row b-tag-title">
                            <div class="col-xs-12 col-md-12 col-lg-12">
                                <h3>搜索到的与<span
                                            class="b-highlight">{{ mb_strimwidth(request()->input('searcharticle'),0,20,'...','utf8') }}</span>相关的文章
                                </h3>
                            </div>
                        </div>
                    @endif
                    {{--@if(count($php)==0)--}}
                        {{--<div class="row b-one-article">--}}
                            {{--<div class="col-xs-12 col-md-12 col-lg-12 b-date">--}}
                                {{--<div class="box">--}}
                                    {{--<center>--}}
                                        {{--<div class="box-header">--}}
                                            {{--<div class="zoom-container">--}}
                                                {{--<div class="zoom-caption">--}}
                                                {{--</div>--}}
                                                {{--<img src="{{asset('image/sorry.jpg')}}"/>--}}
                                            {{--</div>--}}
                                        {{--</div>--}}
                                    {{--</center>--}}
                                    {{--<div class="box-content">--}}
                                        {{--<div id="contact_form">--}}
                                            {{--@if(request()->input('searcharticle'))--}}
                                                {{--<p style="text-align: center;font-size: large">--}}
                                                    {{--博主很懒,暂时没有写关于"{{mb_strimwidth(request()->input('searcharticle'),0,13,'...','utf-8')}}--}}
                                                    {{--"的文章,看看其它文章吧！</p>--}}
                                            {{--@else--}}
                                                {{--<p style="text-align: center;font-size: large">博主很懒,暂时没有写关于"{{$tagName}}--}}
                                                    {{--"的文章,看看其它文章吧！</p>--}}
                                            {{--@endif--}}
                                        {{--</div>--}}
                                    {{--</div>--}}
                                {{--</div>--}}
                            {{--</div>--}}
                        {{--</div>--}}
                    {{--@endif--}}
                    @foreach($list as $ph)
                        <div class="row b-one-article">
                            <h3 class="col-xs-12 col-md-12 col-lg-12">
                                <a class="b-oa-title" href="{{url('article/'.$ph->id.'.html')}}"
                                   style="color: black;">{{mb_strimwidth($ph->title,0,30,'...','utf8')}}</a>
                            </h3>
                            <div class="col-xs-12 col-md-12 col-lg-12 b-date">
                                <ul class="row">
                                    <li class="col-xs-5 col-md-2 col-lg-2">
                                        <i class="fa fa-user"></i>&nbsp;{{$ph->author}}
                                    </li>
                                    <li class="col-xs-7 col-md-3 col-lg-3">
                                        <i class="fa fa-calendar"></i>{{$ph->created_at}}
                                    </li>
                                    <li class="col-xs-5 col-md-2 col-lg-2">
                                        <i class="fa fa-eye">&nbsp;{{$ph->click}}</i>
                                    </li>
                                    <li class="col-xs-7 col-md-5 col-lg-4 "><i class="fa fa-tags"></i>
                                        @foreach($ph->tags as $tag)
                                            <a class="b-tag-name"
                                               style="color: black">{{$tag->name}}</a>
                                        @endforeach
                                    </li>
                                </ul>
                            </div>
                            <div class="col-xs-12 col-md-12 col-lg-12">
                                <div class="row">
                                    <!-- 文章封面图片开始 -->
                                    <div class="col-sm-6 col-md-6 col-lg-4 hidden-xs">
                                        <figure class="b-oa-pic b-style1">
                                            <a href="{{url('article/'.$ph->id.'.html')}}">
                                                <img src="{{url($ph->thumb)}}" alt="胡鑫" title="胡鑫">
                                            </a>
                                            <figcaption>
                                                <a href="" target="_blank"></a>
                                            </figcaption>
                                        </figure>
                                    </div>
                                    <!-- 文章封面图片结束 -->

                                    <!-- 文章描述开始 -->
                                    <div class="col-xs-12 col-sm-6  col-md-6 col-lg-8 b-des-read">
                                        {{$ph->description}}
                                    </div>
                                    <!-- 文章描述结束 -->
                                </div>
                            </div>
                            <a class=" b-readall" href="{{url('article/'.$ph->id.'.html')}}">阅读全文</a>
                        </div>
                    @endforeach
                    <div class="ca_page" style="text-align: center">
                        {{$list->links()}}
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection
<script>
    $(function () {

        /**
         * 处理分页
         */
        $(".ca_page ul li").click(function (e) {
            e.preventDefault();
            var value = $(this).children().text();
            var search = window.location.search;
            var se_re = new RegExp('search');
            var href = '';
            if(se_re.test(search) && search.indexOf('&')===-1){
                href = search + '&page=' + value;
            }else if(se_re.test(search) && search.indexOf('&')!==-1){
                var index = search.indexOf('&');
                var right_search = search.substring(0,index);
                href = right_search +  '&page=' + value;
            }else {
                href = '?page=' + value;
            }
            window.location.href = href;
        });
    })
</script>
<!-- Footer -->
@endsection
