<template>
    <div>
        <footer>
            <div class="wrap-footer">
                <div class="container">
                    <div class="row">
                        <div class="col-md-4 col-footer footer-1">
                            <div class="footer-heading"><h1><span style="color: #fff;">拖油瓶</span></h1></div>
                            <br>
                            <div class="content">
                                <p v-for="item in famous_remark">{{item.val}}</p>
                            </div>
                        </div>
                        <div class="col-md-4 col-footer footer-3">
                            <div class="footer-heading"><h4>博客相关</h4></div>
                            <br>
                            <div class="content">
                                <ul>
                                    <li v-for="it in blog_related">{{it.title}}: {{it.val}}</li>
                                    <li></li>
                                </ul>
                            </div>
                        </div>
                        <div class="col-md-4 col-footer footer-3">
                            <div class="footer-heading"><h4>博主长逛</h4></div>
                            <br>
                            <div class="content">
                                <ul>
                                    <li v-for="cg in blogger_collection"><a :href="cg.val" target="_blank"
                                                                            :title="cg.title"
                                                                            style="color: whitesmoke">{{cg.title}}</a>
                                    </li>
                                    <li><a></a>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <div class="copy-right">
                <p v-html="copyright.val"></p>
            </div>
        </footer>
    </div>
</template>

<script>
    import {getFooterList} from "../../api/footer";

    export default {
        name: 'index',
        data() {
            return {
                famous_remark: [],
                blog_related: [],
                blogger_collection: [],
                copyright: ''
            }
        },
        methods: {
            getFooterList() {
                getFooterList().then(res => {
                    let data = res.data;
                    if (data.code === 200) {
                        this.famous_remark = data.data.famous_remark;
                        this.blog_related = data.data.blog_related;
                        this.blogger_collection = data.data.blogger_collection;
                        this.copyright = data.data.copyright[0];
                    } else {
                        this.$q.notify({
                            color: 'negative',
                            message: data.msg,
                            icon: 'warning',
                            position: 'top'
                        })
                    }
                })
            }
        },
        created() {
            this.getFooterList()
        }
    }
</script>

<style scoped>

</style>
