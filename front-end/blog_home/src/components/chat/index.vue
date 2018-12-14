<template>
  <div>
    <div id="b-content" class="container">
      <div class="row">
        <div class="col-xs-12 col-md-12 col-lg-8 b-chat" v-if="leftChat.length+rightChat.length>=2">
          <div class="b-chat-left">
            <ul class="b-chat-one animated bounceInLeft" v-for="val in leftChat">
              <li class="b-chat-title ">{{val.time}}</li>
              <li class="b-chat-content">{{val.content}}</li>
              <div class="b-arrows-right1">
                <div class="b-arrows-round"></div>
              </div>
              <div class="b-arrows-right2"></div>
            </ul>
          </div>
          <div class="b-chat-middle" :style="styleObject"></div>
          <div class="b-chat-right">
            <!--@foreach($data as $k => $v)-->
            <!--@if($k%2 == 1)-->
            <ul class="b-chat-one animated bounceInRight" v-for="val in rightChat">
              <li class="b-chat-title ">{{val.time}}</li>
              <li class="b-chat-content">{{val.content}}</li>
              <div class="b-arrows-right1">
                <div class="b-arrows-round"></div>
              </div>
              <div class="b-arrows-right2"></div>
            </ul>
            <!--@endif-->
            <!--@endforeach-->
          </div>
        </div>
        <div class="col-xs-12 col-md-12 col-lg-8 b-chat" v-else style="text-align: center">
          暂无说说记录
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: "index",
    data() {
      return {
        leftChat: [{time: '2018-10-14', content: '123321'}],
        rightChat: [{time: '2018-10-14', content: '123321'}],
        styleObject: {
          height: ''
        }
      }
    },
    created(){
      this.hh()
    },
    methods:{
      hh(){
        this.styleObject.height=(this.leftChat.length + this.rightChat.length) * 80 + 'px'
      }
    }
  }
</script>

<style scoped>
  @import "../../statics/css/animate.css";
  @import "../../statics/css/index.css";
</style>
