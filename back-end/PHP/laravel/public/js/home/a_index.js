var index_ops = {
    init:function () {
        this.eventBind();
    },
    eventBind:function () {

        $.goup({
            trigger: 100,
            bottomOffset: 150,
            locationOffset: 100,
            title: '回到顶部',
            titleAsText: true
        });
        $(".link .oauth_login").click(function () {
            let btn = $(this);
            if (btn.hasClass('disabled')) {
                common_ops.alert('正在处理，请勿重复点击！');
                return false;
            }
            btn.addClass('disabled');
            $.ajax({
                url: "/home/wechat/getScene",
                type: 'GET',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                dataType: 'json',
                success: function (res) {
                    btn.removeClass('disabled');
                    if (res.code == 200) {
                        window.scene = res.data;
                    }
                },
                error: function (res) {
                    window.location.href = window.location.href;
                    console.log(res.responseText);
                }
            });
            $("#b-modal-login").modal("show");
        });

        $(".link .oauth_quit").click(function () {
           $.get("/user/quit",function (res) {
               if(res.code==200){
                   window.location.href = window.location.href;
               }
           });
        });

        $(".link_modal").click(function () {
            $("#hx_link_modal").modal("show");
        });

        $("#hx_link_modal .save").click(function () {
            var btn = $(this);
            if (btn.hasClass('disabled')) {
                common_ops.alert('正在处理，请勿重复提交！');
                return false;
            }
            var ur_name = $("#hx_link_modal input[name=ur_name]").val();
            var ur_url = $("#hx_link_modal input[name=ur_url]").val();
            var ur_description = $("#hx_link_modal textarea[name=ur_description]").val();
            if(ur_name.length<1){
                common_ops.tip('请输入正确的友链名字！',$("#hx_link_modal input[name=ur_name]"));
                return false;
            }
            var Expression=/http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?/;
            var objExp=new RegExp(Expression);
            if(objExp.test(ur_url)!=true)
            {
                common_ops.tip('请输入正确的友链地址！',$("#hx_link_modal input[name=ur_url]"));
                return false;
            }
            if(ur_description.length<1){
                common_ops.tip('请对你的友链进行描述！',$("#hx_link_modal textarea[name=ur_description]"));
                return false;
            }
            btn.addClass('disabled');
            $.ajax({
                url:'/link/store',
                type: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    ur_url:ur_url,
                    ur_name:ur_name,
                    ur_description:ur_description
                },
                dataType: 'json',
                success: function (res) {
                    btn.removeClass('disabled');
                    if (res.code == 200) {
                            window.location.href = window.location.href;
                    }
                },
                error: function (res) {
                    console.log(res.responseText);
                }
            });
        });
        let timeout = false;
        let dd = null;
        let dsq = function () {
            if (timeout) return;
            $.ajax({
                url:'/home/wechat/getStatus',
                type: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    scene:window.scene
                },
                dataType: 'json',
                success: function (res) {
                    if (res.code == 200) {
                        timeout = true;
                        window.scene = null;
                        window.location.href = window.location.href;
                    }
                },
                error: function (res) {
                    console.log(res.responseText);
                }
            });
            dd = setTimeout(dsq,3000);
        };

        $('.hx-wechat').click(function () {
            let i = $(this);
           if(timeout){
               timeout = false;
               sleep(1000);
               dsq();
           }else {
               dsq();
           }
           let image = `<img src="https://www.ohdata.top/home/wechat/qrcode/${window.scene}">`;
           i.attr('data-content',image);
           i.popover('show');
            setTimeout(function () {
                clearTimeout(dd);
            },60000);
        });
    },
};

$(document).ready(function () {
    index_ops.init();
});