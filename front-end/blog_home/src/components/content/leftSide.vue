<template>
  <div>
    <div id="b-content">
      <div class="col-xs-12 col-md-12 col-lg-12">
        <div class="row b-one-article" v-for="item in article">
          <h3 class="col-xs-12 col-md-12 col-lg-12">
            <router-link class="b-oa-title" :to="{name:'article',params: { id:item.id }}" style="color: black;" :title="item.title">
              {{item.title.length>30?item.title.substring(0, 30) +"...":item.title}}
            </router-link>
            <label class="btn btn-danger pull-right inline" style="font-size: 10px;" v-if="item.is_top==1">置顶</label>
          </h3>
          <div class="col-xs-12 col-md-12 col-lg-12 b-date">
            <ul class="row">
              <li class="col-xs-5 col-md-2 col-lg-2">
                <i class="fa fa-user"></i>&nbsp;{{item.author}}
              </li>
              <li class="col-xs-7 col-md-3 col-lg-3">
                <i class="fa fa-calendar"></i>&nbsp;{{item.created_at.substring(0,item.created_at.indexOf(' '))}}
              </li>
              <li class="col-xs-5 col-md-2 col-lg-2">
                <i class="fa fa-eye"></i>&nbsp;{{item.click}}

              </li>
              <li class="col-xs-7 col-md-5 col-lg-4 "><i class="fa fa-tags"></i>
                <a class="b-tag-name" style="color: black" v-for="tag in item.tags">{{tag.name}}</a>
              </li>
            </ul>
          </div>
          <div class="col-xs-12 col-md-12 col-lg-12">
            <div class="row">
              <!-- 文章封面图片开始 -->
              <div class="col-sm-6 col-md-6 col-lg-4 hidden-xs">
                <figure class="b-oa-pic b-style1">
                  <router-link :to="{name:'article',params: { id:item.id }}" :title="item.title">
                    <img :src="item.thumb" :alt="item.title" :title="item.title">
                  </router-link>
                  <figcaption>
                    <router-link :to="{name:'article',params: { id:item.id }}" :title="item.title"></router-link>
                  </figcaption>
                </figure>
              </div>
              <!-- 文章封面图片结束 -->

              <!-- 文章描述开始 -->
              <div class="col-xs-12 col-sm-6  col-md-6 col-lg-8 b-des-read">
                {{item.description}}
              </div>
              <!-- 文章描述结束 -->
            </div>
          </div>
          <router-link class=" b-readall" :to="{name:'article',params: { id:item.id }}">阅读全文</router-link>
        </div>
        <div class="a_page" style="text-align: center">
          <q-btn :loading="loading" color="primary" @click="simulateProgress">
            加载更多
            <span slot="loading">
          <q-spinner-hourglass class="on-left"/>
          加载中...
        </span>
          </q-btn>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import {getArticleList} from '../../api/index'
  export default {
    name: 'leftSide',
    data() {
      return {
        loading:false,
        visible:false,
        article:[],
        listData:{
          pageSize:8,
          pageNum:1
        }
      }
    },
    methods:{
      getArticleList(){
        getArticleList(this.listData).then(res=>{
          this.loading = false;
          let data = res.data;
          if(data.code==200){
            this.article = data.data;
          }else {
            this.$q.notify({
              color: 'negative',
              message: data.msg,
              icon: 'warning',
              position:'top'
            })
          }
        }).catch(err=>{
          this.$q.notify({
            color: 'negative',
            message: '服务器错误！',
            icon: 'warning',
            position:'top'
          })
        })
      },
      simulateProgress () {
        // we set loading state
        this.loading = true;
        this.listData.pageSize += 3;
        this.getArticleList();
      },
    },
    created() {
      this.getArticleList();
    }
  }
</script>

<style scoped>
  @import "../../statics/css/index.css";
  .bg-primary {
    color: #fff;
    background-color: #428bca;
    width: 100px;
  }
</style>
