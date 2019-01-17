<template>
    <div>
        <div id="b-content" class="container">
            <div class="row">
                <div class="col-xs-12 col-md-12 col-lg-8">

                    <div class="row b-tag-title">
                        <div class="col-xs-12 col-md-12 col-lg-12">
                            <h3>拥有<span class="b-highlight">123321</span>标签的文章</h3>
                        </div>
                    </div>

                    <div class="row b-tag-title">
                        <div class="col-xs-12 col-md-12 col-lg-12">
                            <h3>搜索到的与<span
                                    class="b-highlight">asdasda</span>相关的文章
                            </h3>
                        </div>
                    </div>
                    <div class="row b-one-article">
                        <div class="col-xs-12 col-md-12 col-lg-12 b-date">
                            <div class="box">
                                <center>
                                    <div class="box-header">
                                        <div class="zoom-container">
                                            <div class="zoom-caption">
                                            </div>
                                            <img src=""/>
                                        </div>
                                    </div>
                                </center>
                                <div class="box-content">
                                    <div id="contact_form">
                                        <p style="text-align: center;font-size: large">
                                            博主很懒,暂时没有写关于"21312312
                                            "的文章,看看其它文章吧！</p>
                                        @else
                                        <p style="text-align: center;font-size: large">博主很懒,暂时没有写关于"1231
                                            "的文章,看看其它文章吧！</p>
                                        @endif
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="row b-one-article">
                        <h3 class="col-xs-12 col-md-12 col-lg-12">
                            <a class="b-oa-title" href="123213"
                               style="color: black;">1232312</a>
                        </h3>
                        <div class="col-xs-12 col-md-12 col-lg-12 b-date">
                            <ul class="row">
                                <li class="col-xs-5 col-md-2 col-lg-2">
                                    <i class="fa fa-user"></i>&nbsp;12312
                                </li>
                                <li class="col-xs-7 col-md-3 col-lg-3">
                                    <i class="fa fa-calendar"></i>2018-10-24
                                </li>
                                <li class="col-xs-5 col-md-2 col-lg-2">
                                    <i class="fa fa-eye">&nbsp;123</i>
                                </li>
                                <li class="col-xs-7 col-md-5 col-lg-4 "><i class="fa fa-tags"></i>

                                    <a class="b-tag-name"
                                       style="color: black">php</a>
                                </li>
                            </ul>
                        </div>
                        <div class="col-xs-12 col-md-12 col-lg-12">
                            <div class="row">
                                <!-- 文章封面图片开始 -->
                                <div class="col-sm-6 col-md-6 col-lg-4 hidden-xs">
                                    <figure class="b-oa-pic b-style1">
                                        <a href="">
                                            <img src="" alt="胡鑫" title="胡鑫">
                                        </a>
                                        <figcaption>
                                            <a href="" target="_blank"></a>
                                        </figcaption>
                                    </figure>
                                </div>
                                <!-- 文章封面图片结束 -->

                                <!-- 文章描述开始 -->
                                <div class="col-xs-12 col-sm-6  col-md-6 col-lg-8 b-des-read">
                                    21312312
                                </div>
                                <!-- 文章描述结束 -->
                            </div>
                        </div>
                        <a class=" b-readall" href="">阅读全文</a>
                    </div>

                    <div class="ca_page" style="text-align: center">

                    </div>
                </div>
            </div>
        </div>
    </div>

</template>

<script>
    export default {
        name: "index"
    }
</script>

<style scoped>
    @import "../../statics/css/index.css";
</style>
