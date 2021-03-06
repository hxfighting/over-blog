@extends('home.layouts.home')
@section('content')
@section('main-content')


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
                        <center><input class="sendButton" type="submit" name="Submit" value="确定" onclick="contact()"></center>
                    </div>
                </div>
            </div>
        </div>
    </div>
@endsection

<script>
    function contact(){
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
        if(email.length<1){
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
        if(content.length<1){
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
            data: {
                name: name,
                email: email,
                content: content
            },
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
                }else {
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
@endsection
