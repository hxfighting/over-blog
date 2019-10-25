var index_ops = {
    init: function () {
        this.eventBind();
    },
    eventBind: function () {

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
                url: "/wechat/scene",
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
            $.get("/logout", function (res) {
                if (res.code == 200) {
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
                Swal({
                    position: 'center',
                    type: 'error',
                    title: '正在处理,请勿重复点击!',
                    showConfirmButton: false,
                    timer: 1500
                })
                return false;
            }
            var ur_name = $("#hx_link_modal input[name=ur_name]").val();
            var ur_url = $("#hx_link_modal input[name=ur_url]").val();
            var ur_description = $("#hx_link_modal textarea[name=ur_description]").val();
            if (ur_name.length < 1) {
                Swal({
                    position: 'center',
                    type: 'error',
                    title: '请输入正确的友链名字！',
                    showConfirmButton: false,
                    timer: 1500
                })
                return false;
            }
            var Expression = /http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?/;
            var objExp = new RegExp(Expression);
            if (objExp.test(ur_url) != true) {
                Swal({
                    position: 'center',
                    type: 'error',
                    title: '请输入正确的友链地址！',
                    showConfirmButton: false,
                    timer: 1500
                })
                return false;
            }
            if (ur_description.length < 1) {
                Swal({
                    position: 'center',
                    type: 'error',
                    title: '请对你的友链进行描述！',
                    showConfirmButton: false,
                    timer: 1500
                })
                return false;
            }
            btn.addClass('disabled');
            $.ajax({
                url: '/link',
                type: 'POST',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    url: ur_url,
                    name: ur_name,
                    description: ur_description
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
        });
        let timeout = false;
        let dd = null;
        let dsq = function () {
            if (timeout) return;
            $.ajax({
                url: '/wechat/status',
                type: 'get',
                headers: {
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                data: {
                    scene: window.scene
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
            dd = setTimeout(dsq, 3000);
        };

        $('#b-modal-login').on('hidden.bs.modal', function (event) {
            if(dd!=null){
                clearTimeout(dd)
            }
        })

        $('.hx-wechat').click(function () {
            let i = $(this);
            if (dd == null) {
                dsq();
            }
            let image = `<img src="${APP_URL}/wechat/qrcode/${window.scene}">`;
            i.attr('data-content', image);
            i.popover('show');
            setTimeout(function () {
                clearTimeout(dd);
            }, 60000);
        });
    },
};

$(document).ready(function () {
    index_ops.init();
});
