<template>
    <div>
        <div class="featured container">
            <div class="col-sm-8">
                <q-carousel
                        color="white"
                        arrows
                        height="400px"
                        autoplay
                        infinite
                >
                    <q-carousel-slide
                            :img-src="item.image_url" v-for="item in rotation">
                        <div class="absolute-bottom custom-caption">
                            <div class="q-display-1">{{item.image.words}}</div>
                            <!--<div class="q-headline">Mountains</div>-->
                        </div>
                    </q-carousel-slide>
                </q-carousel>
            </div>
            <div class="col-sm-4">
                <div class="col-sm-12">
                    <q-carousel
                            color="white"
                            quick-nav
                            height="300px"
                            autoplay
                            infinite
                    >
                        <q-carousel-slide
                                :img-src="img.image_url" v-for="img in photo"/>
                    </q-carousel>
                </div>
            </div>
            <div class="col-sm-4" style="height: 99px;">
                <p class="h_word">
                    123
                </p>
            </div>
        </div>
    </div>
</template>

<script>
    import {getRotation} from '../../api/header'
    export default {
        name: 'bottom',
        data() {
            return {
                rotation:[],
                photo:[]
            }
        },
        methods:{
            getRotation(){
                getRotation().then(res=>{
                    let re = res.data;
                    if(re.code===200){
                        this.rotation = re.data.rotation;
                        this.photo = re.data.photo;
                    }else {
                        this.rotation = [];
                        this.photo = [];
                        this.$q.notify({
                            color: 'negative',
                            message: re.msg,
                            icon: 'warning',
                            position:'top'
                        })
                    }
                }).catch(err=>{
                    this.$q.notify({
                        color: 'negative',
                        message: '服务器错误!',
                        icon: 'warning',
                        position:'top'
                    })
                })
            }
        },
        created() {
            this.getRotation()
        }
    }
</script>

<style scoped>
    .custom-caption {
        text-align: center;
        padding: 12px;
        color: #e0e0e0;
        background: rgba(0, 0, 0, .5);
    }

    .featured {
        background-color: #fff;
        padding: 15px;
        border: 1px solid #ddd;
    }

    .h_word {
        text-align: center;
        vertical-align: middle;
    }
</style>
