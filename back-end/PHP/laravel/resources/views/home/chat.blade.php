@extends('home.layouts.home')

@section('content')
@section('main-content')
            <div id="main-content" class="col-md-8">
                <div id="b-content" class="container">
                    <div class="row">
                        <div class="col-xs-12 col-md-12 col-lg-8 b-chat">
                            <div class="b-chat-left">
                                @foreach($data as $k => $v)
                                    @if($k%2 == 0)
                                        <ul class="b-chat-one animated bounceInLeft">
                                            <li class="b-chat-title ">{{ date('Y-m-d H:i:s',$v->ch_time)}}</li>
                                            <li class="b-chat-content">{{ $v->ch_content }}</li>
                                            <div class="b-arrows-right1">
                                                <div class="b-arrows-round"></div>
                                            </div>
                                            <div class="b-arrows-right2"></div>
                                        </ul>
                                    @endif
                                @endforeach
                            </div>
                            <div class="b-chat-middle"></div>
                            <div class="b-chat-right">
                                @foreach($data as $k => $v)
                                    @if($k%2 == 1)
                                        <ul class="b-chat-one animated bounceInRight">
                                            <li class="b-chat-title ">{{ date('Y-m-d H:i:s',$v->ch_time) }}</li>
                                            <li class="b-chat-content">{{ $v->ch_content }}</li>
                                            <div class="b-arrows-right1">
                                                <div class="b-arrows-round"></div>
                                            </div>
                                            <div class="b-arrows-right2"></div>
                                        </ul>
                                    @endif
                                @endforeach
                            </div>
                        </div>
                    </div>
                </div>
            </div>
@endsection
    <script src="{{asset('js/jquery/index.js')}}" type="text/javascript"></script>

@endsection
