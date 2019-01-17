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
                    <router-link class="label article_tag_f" :to="{name:'tag',params: { id:ta.id }}" :title="ta.name" v-for="ta in tag">{{ta.name}}</router-link>
                </div>
            </div>
            <div class="widget ">
                <div class="heading"><h4>最热文章</h4></div>
                <br>
                <div class="content">
                    <div class="wrap-vid" v-for="ar in hot_article">
                        <h3 class="vid-name">
                            <router-link :to="{name:'article',params: { id:ar.id }}" :title="ar.title">{{ar.title.length>30?ar.title.substring(0, 30) +"...":ar.title}}</router-link>
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
                                 class="img-circle img-responsive" :title="com.name"></a>
                        <div class="wrapper">
                            <router-link :to="{name:'article',params: { id:com.article_id }}" :title="com.content">{{com.content.length>10?com.content.substring(0, 10) +"...":com.content}}</router-link>
                            <ul class="list-inline">
                                <li><i class="fa fa-calendar"></i>&nbsp;{{com.created_at}}</li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
            <div class="widget wid-tags">
                <div class="heading"><h4>友情链接</h4></div>
                <p><a style="color: black;cursor: pointer" class="link_modal" @click="addLinkModal = true">申请友链</a></p>
                <div class="content">
                    <a class="label article_tag_f" :href="li.url" style="color: black" :title="li.name"
                       target="_blank" v-for="li in link">{{li.name}}</a>
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
        <div>
            <q-dialog
                    v-model="addLinkModal"
                    stack-buttons
                    prevent-close
                    @cancel="onCancel"
                    @ok="onOk"
            >
                <!-- This or use "title" prop on <q-dialog> -->
                <span slot="title">Favorite Superhero</span>

                <!-- This or use "message" prop on <q-dialog> -->
                <span slot="message">What is your superhero of choice?</span>

                <div slot="body">
                    <q-field
                            icon="account_circle"
                            helper="We need your name so we can send you to the movies."
                            label="Your name"
                            :label-width="3"
                            :error="nameError"
                    >
                        <q-input v-model="formData.name" />
                    </q-field>
                </div>
            </q-dialog>

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
                link:[],
                addLinkModal:false,
                formData:{
                    url:'',
                    name:'',
                    description:''
                },
                nameError:false
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
            },
            onCancel(){

            },
            onOk(){

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
