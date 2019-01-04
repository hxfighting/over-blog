<template>
    <div>
        <div id="sidebar">
            <!---- Start Widget ---->
            <div class="widget wid-tags">
                <div class="heading"><h4>搜索</h4></div>
                <br>
                <div class="content">
                    <form role="form" class="form-horizontal" method="get" action="http://192.168.1.161/search">
                        <input type="text" placeholder="回车键搜索文章" value="" name="searcharticle" id="v_search"
                               class="form-control">
                    </form>
                </div>
            </div>
            <!---- Start Widget ---->
            <div class="widget wid-tags">
                <div class="heading"><h4>标签云</h4></div>
                <br>
                <div class="content">
                    <a class="label article_tag_f" href="http://192.168.1.161/tag/1.html" v-for="ta in tag">{{ta.name}}</a>
                </div>
            </div>
            <div class="widget ">
                <div class="heading"><h4>最热文章</h4></div>
                <br>
                <div class="content">
                    <div class="wrap-vid" v-for="ar in hot_article">
                        <h3 class="vid-name"><a href="http://192.168.1.161/article/96.html"
                                                :title="ar.title">{{ar.title}}</a>
                        </h3>
                        <div class="info">
                            <span><i class="fa fa-calendar"></i>{{ar.created_at.substring(0,ar.created_at.indexOf(' '))}}</span>
                            <span><i class="fa fa-eye"></i>{{ar.click}}</span>
                            <span><i class="fa fa-comments-o"></i>{{ar.comment_count}}&nbsp;Comments</span>
                        </div>
                    </div>
                </div>
            </div>
            <!---- Start Widget ---->
            <div class="widget wid-comment">
                <div class="heading"><h4>最新评论</h4></div>
                <br>
                <div class="content">
                    <div class="post" v-for="com in comment">
                        <a> <img :src="com.avatar"
                                 class="img-circle img-responsive" :title="com."></a>
                        <div class="wrapper">
                            <a href="http://192.168.1.161/article/100.html"><span>抱歉，daocloud目前不...</span></a>
                            <ul class="list-inline">
                                <li><i class="fa fa-calendar"></i>&nbsp;2018-09-25 14:44:29</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
            <div class="widget wid-tags">
                <div class="heading"><h4>友情链接</h4></div>
                <p><a style="color: black;cursor: pointer" class="link_modal">申请友链</a></p>
                <div class="content">
                    <a class="label article_tag_f" href="http://123.207.241.214" style="color: black" title="王大大的博客"
                       target="_blank">王大大的博客</a>
                    <a class="label article_tag_f" href="http://9898192.cn" style="color: black" title="一个吊儿郎当的骚年！"
                       target="_blank">王洋洋的文笔博客</a>
                    <a class="label article_tag_f" href="http://tunanshan.com" style="color: black" title="图南山"
                       target="_blank">图南山</a>
                    <a class="label article_tag_f" href="http://www.whiteeeen.cn/" style="color: black" title="陈杨的博客"
                       target="_blank">陈杨的博客</a>
                    <a class="label article_tag_f" href="http://www.hzj233.cn" style="color: black" title="dota2交友论坛"
                       target="_blank">dota2交友论坛</a>
                    <a class="label article_tag_f" href="http://www.bestmx.top/" style="color: black" title="疯言疯语"
                       target="_blank">疯言疯语</a>
                    <a class="label article_tag_f" href="https://ituring.me" style="color: black" title="iTuring"
                       target="_blank">iTuring</a>
                    <a class="label article_tag_f" href="http://www.zsk6.top" style="color: black" title="吃辣椒的小蜜蜂"
                       target="_blank">吃辣椒的小蜜蜂</a>
                    <a class="label article_tag_f" href="http://www.ericnothing.cn/" style="color: black" title="个人博客"
                       target="_blank">Eric-Nothing</a>
                </div>
            </div>

            <!-- link --->
            <div class="modal fade" id="hx_link_modal" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel">
                <div class="modal-dialog" role="document">
                    <div class="modal-content">
                        <div class="modal-header">
                            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
                                    aria-hidden="true">×</span></button>
                            <h4 class="modal-title" id="exampleModalLabel">友链申请</h4>
                        </div>
                        <div class="modal-body">
                            <div>
                                <div class="form-group">
                                    <label for="ur_name" class="control-label">友链名字:</label>
                                    <input type="text" class="form-control" id="ur_name" name="ur_name">
                                </div>
                                <div class="form-group">
                                    <label for="ur_url" class="control-label">友链URL:</label>
                                    <input type="text" class="form-control" id="ur_url" name="ur_url" value="http://">
                                </div>
                                <div class="form-group">
                                    <label for="ur_description" class="control-label">友链描述:</label>
                                    <textarea class="form-control" id="ur_description" name="ur_description"></textarea>
                                </div>
                            </div>
                        </div>
                        <div class="modal-footer">
                            <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                            <button type="button" class="btn btn-primary save">提交</button>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
    import {getSidebarList} from "../../api";

    export default {
        name: 'rightSide',
        data() {
            return {
                tag:[],
                hot_article:[],
                comment:[],
                link:[]
            }
        },
        methods:{
            getSidebarList(){
                getSidebarList().then(res=>{
                    let data = res.data;
                    if(data.code==200){
                        this.tag = data.data.tag;
                        this.hot_article = data.data.hot_article;
                        this.comment = data.data.comment;
                        this.link = data.data.link;
                    }else {
                        this.$q.notify({
                            color: 'negative',
                            message: data.msg,
                            icon: 'warning',
                            position:'top'
                        })
                    }
                })
            }
        },
        created() {
            this.getSidebarList();
        }
    }
</script>

<style scoped>
    @import "../../statics/css/style.css";

    .article_tag_f {
        float: left;
        margin: 2px 2px;
    }
</style>
