<template>
    <div>
        <Row :gutter="20">
            <i-col :xs="12" :md="8" :lg="4" v-for="(infor, i) in inforCardData" :key="`infor-${i}`"
                   style="height: 120px;padding-bottom: 10px;">
                <infor-card shadow :color="infor.color" :icon="infor.icon" :icon-size="36">
                    <count-to :end="infor.count" count-class="count-style"/>
                    <p>{{ infor.title }}</p>
                </infor-card>
            </i-col>
        </Row>
    </div>
</template>

<script>
    import InforCard from '_c/info-card'
    import CountTo from '_c/count-to'
    import {getCountData} from "../../../api/dashboard";

    export default {
        name: 'home',
        components: {
            InforCard,
            CountTo,
        },
        data() {
            return {
                inforCardData: []
            }
        },
        methods: {
            getCountData() {
                this.$Spin.show();
                getCountData().then(res => {
                    this.$Spin.hide();
                    let data = res.data;
                    if (data.code == 200) {
                        this.inforCardData = data.data;
                    } else {
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created(){
            this.getCountData()
        },
        mounted() {
            //
        }
    }
</script>

<style lang="less">
    .count-style {
        font-size: 50px;
    }
</style>
