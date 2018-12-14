<template>
    <div class="featured container">
        <span style="margin-left: 19px">随机文章推荐</span>
        <br>
       <div class="swiper-inner">
           <swiper :options="swiperOption">
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer" title="123"></swiper-slide>
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer" title="456"></swiper-slide>
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer"></swiper-slide>
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer"></swiper-slide>
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer"></swiper-slide>
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer"></swiper-slide>
               <swiper-slide style="background-image:url(http://192.168.1.161/upload/rotation/0da2db5b879a59aaed0903da3ae440d4.jpg);cursor: pointer"></swiper-slide>
               <div class="swiper-pagination" slot="pagination"></div>
           </swiper>
       </div>
    </div>
</template>

<script>
    import {swiper} from 'swiper'
export default {
  name: 'detailHeaderBottom',
    data(){
      return{
          swiperOption: {
              autoplay: {
                  disableOnInteraction: false,
                  delay: 3000
              },
              grabCursor: false,
              centeredSlides: true,
              slidesPerView: 'auto',
              slideToClickedSlide:true,
              preventClicksPropagation:true,
              spaceBetween: 30,
              slidesPerGroup: 5,
              loop: true,
              pagination: {
                  el: '.swiper-pagination'
              }
          }
      }
    }
}
</script>

<style scoped>
    .featured {
        background-color: #fff;
        padding: 15px;
        border: 1px solid #ddd;
    }
    .swiper-inner {
        width: 100%;
        height: 150px;
        padding-bottom: 50px;
    }
    .swiper-slide {
        background-position: center;
        background-size: cover;
        width: 200px;
        height: 140px;
    }
</style>
