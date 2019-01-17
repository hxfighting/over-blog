@extends('home.layouts.home')
@section('content')
@section('main-content')


    <div id="main-content" class="col-md-8">
        <div class="btn-success" style="background-color: white">
            @if(count($errors)>0)
                @if(is_object($errors))
                    @foreach($errors->all() as $error)
                        <div class="row b-tag-title">
                            <div class="col-xs-12 col-md-12 col-lg-12">
                                <h3 style="color: red">{{ $error }}</h3>
                            </div>
                        </div>
                    @endforeach
                @else
                    <div class="row b-tag-title">
                        <div class="col-xs-12 col-md-12 col-lg-12">
                            <h3 style="color: red">{{ $errors }}</h3>
                        </div>
                    </div>
                @endif
            @endif
        </div>
        <div class="box">
            <center>
                <div class="box-header">
                    <h1 class="center">Contact</h1>
                </div>
            </center>
            <div class="box-content">
                <div id="contact_form">
                    <form name="contact1" id="ff" method="post" action="{{url('contact')}}" onSubmit="return submitOnce(this)">
                        <input type="hidden" name="token2" value="{{session('token2')}}"/>
                        {{csrf_field()}}
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
                        <center><input class="sendButton" type="submit" name="Submit" value="确定"></center>
                    </form>
                </div>
            </div>
        </div>
    </div>
@endsection

<script language="javascript">
    var submitcount=0;
    function submitOnce (form){
        if (submitcount == 0){
            submitcount++;
            return true;
        } else{
            alert("正在操作，请不要重复提交，谢谢！");
            return false;
        }
    }
</script>
@endsection
@push('scripts')
    @include('flashy::message')
@endpush