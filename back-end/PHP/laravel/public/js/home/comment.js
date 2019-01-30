var comment_ops = {
    init:function () {
        this.eventBind();
    },
    eventBind:function () {
        $(".a_comment .save").click(function () {
            alert(123)
            var btn = $(this);
            if (btn.hasClass('disabled')) {
                common_ops.alert('正在处理，请勿重复提交！');
                return false;
            }
            // if (user_id.length < 1 || isNaN(user_id)) {
            //     $("#b-modal-login").modal("show");
            //     return false;
            // }
            let article_id = $(this).attr('data');
            var co_email = $(".a_comment input[name=email]").val();
            var co_content = $(".a_comment textarea[name=message]").val();
            var pattern = /^([\.a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
            if (co_email.length < 1 || !pattern.test(co_email)) {
                common_ops.tip('请输入正确的邮箱！', $(".a_comment input[name=email]"));
                return false;
            }
            if (co_content.length < 1) {
                common_ops.tip('说点评论内容吧！', $(".a_comment textarea[name=message]"));
                return false;
            }
            if (co_aid.length < 1 || isNaN(co_aid)) {
                common_ops.alert('请选择要评论的文章！');
                return false;
            }
            btn.addClass('disabled');
            $.ajax({
                url: "/comment",
                type: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    user_id: user_id,
                    article_id: co_aid,
                    co_email: co_email,
                    co_content:co_content,
                    user_name:user_name
                },
                dataType: 'json',
                success: function (res) {
                    btn.removeClass('disabled');
                    if (res.code == 200) {
                        window.location.href = window.location.href;
                        layer.msg(res.msg,{icon:6});
                    }
                },
                error: function (res) {
                    window.location.href = window.location.href;
                    layer.msg('评论失败，请稍后再试',{icon:5});
                    console.log(res.responseText);
                }
            });
        });

        // $(".hx-comments .c_reply").click(function () {
        //     if (user_id.length < 1 || isNaN(user_id)) {
        //         $("#b-modal-login").modal("show");
        //         return false;
        //     }
        //     $(this).parents('.wrapper').next().css('display', 'block');
        // });
        // $(".hx-comments .co_cancel").click(function () {
        //     if (user_id.length < 1 || isNaN(user_id)) {
        //         $("#b-modal-login").modal("show");
        //         return false;
        //     }
        //    $(this).parent('.hx_group').css('display','none');
        // });
        // $(".hx-comments .hx_child_co_reply").click(function () {
        //     if (user_id.length < 1 || isNaN(user_id)) {
        //         $("#b-modal-login").modal("show");
        //         return false;
        //     }
        //     $(this).parents('.child_wrapper').next().css('display','block');
        // });
        // $(".hx-comments .child_co_cancel").click(function () {
        //     if (user_id.length < 1 || isNaN(user_id)) {
        //         $("#b-modal-login").modal("show");
        //         return false;
        //     }
        //     $(this).parent('.hx_child_group').css('display','none');
        // });
        $(".hx-comments .co_reply").click(function () {
            var btn = $(this);
            if (btn.hasClass('disabled')) {
                layer.msg('正在处理，请勿重复提交！',{icon:5});
                return false;
            }
            var co_pid = $(this).attr('data-pid');
            var co_email = $(this).parent('.hx_group').children('.co_email').val();
            var co_content = $(this).parent('.hx_group').children('.co_content').val();
            var co_aid =  $(this).attr('data-aid');
            var co_rid = $(this).attr('data-rid');
            var co_uid = user_id;
            if (co_uid.length < 1 || isNaN(co_uid)) {
                $("#b-modal-login").modal("show");
                return false;
            }
            if(co_pid.length<1 || isNaN(co_pid)){
                comment_ops.alert('请选择要回复的评论！');
                return false;
            }
            var pattern = /^([\.a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
            if (co_email.length < 1 || !pattern.test(co_email)) {
                common_ops.tip('请输入正确的邮箱！', $(this).parent('.hx_group').children('.co_email'));
                return false;
            }
            if(co_content.length<1){
                common_ops.tip('说点什么吧！', $(this).parent('.hx_group').children('.co_content'));
                return false;
            }
            btn.addClass('disabled');
            $.ajax({
                url:"/comment/store",
                type: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    co_uid:co_uid,
                    co_aid:co_aid,
                    co_rid:co_rid,
                    co_pid: co_pid,
                    co_email:co_email,
                    co_content: co_content,
                    user_name:user_name
                },
                dataType: 'json',
                success: function (res) {
                    btn.removeClass('disabled');
                    if (res.code == 200) {
                        window.location.href = window.location.href;
                        layer.msg(res.msg,{icon:6});
                    }
                },
                error: function (res) {
                    window.location.href = window.location.href;
                    layer.msg('回复失败，请稍后再试！',{icon:5});
                    console.log(res.responseText);
                }
            });
        });
        $(".hx-comments .hx_child_comments .child_co_reply").click(function () {
            var btn = $(this);
            if (btn.hasClass('disabled')) {
                layer.msg('正在处理，请勿重复提交！',{icon:5});
                return false;
            }
            var co_pid = $(this).attr('data-pid');
            var co_email = $(this).parent('.hx_child_group').children('.co_email').val();
            var co_content = $(this).parent('.hx_child_group').children('.co_content').val();
            var co_aid =  $(this).attr('data-aid');
            var co_rid = $(this).attr('data-rid');
            var co_uid = user_id;
            if (co_uid.length < 1 || isNaN(co_uid)) {
                $("#b-modal-login").modal("show");
                return false;
            }
            if(co_pid.length<1 || isNaN(co_pid)){
                comment_ops.alert('请选择顶级评论！');
                return false;
            }
            var pattern = /^([\.a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
            if (co_email.length < 1 || !pattern.test(co_email)) {
                common_ops.tip('请输入正确的邮箱！', $(this).parent('.hx_child_group').children('.co_email'));
                return false;
            }
            if(co_content.length<1){
                common_ops.tip('说点什么吧！', $(this).parent('.hx_child_group').children('.co_content'));
                return false;
            }
            if(co_rid.length<1 || isNaN(co_pid)){
                comment_ops.alert('请选择要回复的评论！');
                return false;
            }
            btn.addClass('disabled');
            $.ajax({
                url:"/comment/store",
                type: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    co_uid:co_uid,
                    co_aid:co_aid,
                    co_pid:co_pid,
                    co_rid:co_rid,
                    co_email: co_email,
                    co_content: co_content,
                    user_name:user_name
                },
                dataType: 'json',
                success: function (res) {
                    btn.removeClass('disabled');
                    if (res.code == 200) {
                        window.location.href = window.location.href;
                        layer.msg(res.msg,{icon:6});
                    }
                },
                error: function (res) {
                    window.location.href = window.location.href;
                    layer.msg('回复失败，请稍后再试！',{icon:5});
                    console.log(res.responseText);
                }
            });
        });
    }
};

$(document).ready(function () {
   comment_ops.init();
});
