<template>
    <div>
        <nav id="menu" class="navbar container">
            <div class="navbar-header">
                <button type="button" class="btn btn-navbar navbar-toggle" data-toggle="collapse"
                        data-target=".navbar-ex1-collapse"><i class="fa fa-bars"></i></button>
                <router-link to="/" class="navbar-brand">
                    <div class="logo" :title="title"><span class="ha_title">{{title}}</span></div>
                </router-link>
            </div>
            <div class="collapse navbar-collapse navbar-ex1-collapse">
                <ul class="nav navbar-nav">
                    <li class="dropdown" v-for="item of category">
                        <router-link class="dropdown-toggle" :to="item.url" :title="item.title">{{item.title}}</router-link>
                        <div class="dropdown-menu" v-if="item.children.length>0">
                            <div class="dropdown-inner">
                                <ul class="list-unstyled">
                                    <li v-for="child of item.children">
                                        <router-link :to="child.url" :title="child.title">{{child.title}}</router-link>
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </li>
                </ul>
                <ul class="list-inline navbar-right top-social">
                    <li><a title="博主QQ" :href="socialAddress.qq"
                           target="_blank"><i class="fa fa-qq"></i></a></li>
                    <li><a title="博主twitter" :href="socialAddress.twitter" target="_blank"><i
                            class="fa fa-twitter"></i></a></li>
                    <li><a title="博主git" :href="socialAddress.git" target="_blank"><i class="fa fa-git"></i></a>
                    </li>
                    <li><a title="博主github" :href="socialAddress.github" target="_blank"><i
                            class="fa fa-github"></i></a>
                    </li>
                    <li><a title="博主新浪微博"
                           :href="socialAddress.weibo"
                           target="_blank"><i class="fa fa-weibo"></i></a></li>
                </ul>
            </div>
        </nav>

    </div>
</template>

<script>
    import {getCategory} from '../../api/header'
    import config from '../../config/index'
    const {title} = config;
    export default {
        name: 'middle',
        data() {
            return {
                category:[],
                title:title,
                socialAddress:{
                    qq:'',
                    weibo:'',
                    twitter:'',
                    git:'',
                    github:''
                }
            }
        },
        methods:{
            getCategory(){
                getCategory().then(res=>{
                    let re = res.data;
                    if(re.code===200){
                        this.category = re.data.list;
                        this.socialAddress = re.data.social;
                    }else {
                        this.$q.notify({
                            color: 'negative',
                            message: re.msg,
                            icon: 'warning',
                            position:'top'
                        })
                        this.category = []
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
            this.getCategory()
        }
    }
</script>

<style scoped>
    @import "../../statics/css/bootstrap.min.css";
    @import "../../statics/css/font-awesome-4.4.0/css/font-awesome.min.css";

    #menu {
        background-color: #252628;
        border-radius: 0;
    }

    #menu a.navbar-brand:hover {
        color: #fff;
    }

    #menu .navbar-brand {
        margin: 0;
        padding: 14px 18px;
        background-color: #000000;
        height: auto;
        text-transform: uppercase;
    }

    #menu .logo {
    }

    #menu .navbar-collapse {
        padding: 0;
    }

    #menu .dropdown-menu {
        background-color: #fff;
        border: none;
    }

    #menu ul.nav .dropdown-menu li a {
        color: #000;
        font-weight: bold;
        border-radius: 5px !important;
    }

    #menu ul.nav .dropdown-menu li a:hover {
        color: #fff;
        background-color: #000;
    }

    #menu .dropdown-inner {
        display: table;
    }

    #menu .dropdown-inner ul {
        display: table-cell;
    }

    #menu .dropdown-inner a {
        min-width: 160px;
        display: block;
        padding: 3px 20px;
        clear: both;
        line-height: 20px;
        color: #000;
        font-size: 14px;
        margin: 0 5px;
    }

    #menu li.dropdown:hover > a, #menu li.dropdown:focus > a, #menu li.dropdown:active > a {
        background: #fff;
        color: #000;
    }

    #menu ul.nav li a {
        font-weight: bold;
        color: #fff;
        border-radius: 5px;
        padding: 14px 19px;
    }

    #menu ul.nav li.dropdown a {
        border-top-left-radius: 5px;
        border-top-right-radius: 5px;
        border-bottom-right-radius: 0px;
        border-bottom-left-radius: 0px;
    }

    #menu .nav > li > a:hover, #menu .nav > li > a:focus {
        background-color: #fff;
        color: #000;
    }

    #menu .top-social {
        background-color: #505052;
        padding: 5px;
        font-size: 20px;
        text-align: center;
    }

    #menu ul.top-social {
        margin: 0;
    }

    #menu ul.top-social li {
        height: 38px;
        width: 38px;
    }

    #menu ul.top-social a i {
        color: #fff;
        line-height: 1.9;
        height: 38px;
        width: 38px;
        border-radius: 50%;
        background-color: #252628;
    }

    #menu ul.top-social a i:hover {
        background-color: #fff;
        color: #000;
    }

    #menu .btn-navbar {
        font-size: 20px;
        color: #FFF;
        padding: 5px 15px;
        float: right;
        border: 3px solid #fff;
    }

    @media (max-width: 768px) {
        #menu {
            background-color: #000;
            border-radius: 0;
        }
    }

    @media (min-width: 768px) {
        #menu.navbar {
            padding: 0;
            height: auto;
            margin-top: 20px;
        }

        #menu .dropdown:hover .dropdown-menu {
            display: block;
        }

    }

    @media (max-width: 767px) {

        #menu .navbar-brand {
            display: block;
            color: #fff;
        }

        #menu div.dropdown-inner > ul.list-unstyled {
            display: block;
        }

        #menu .dropdown-inner a {
            width: 100%;
            color: #fff;
        }

        #menu div.dropdown-menu {
            margin-left: 0 !important;
            padding-bottom: 10px;
            background-color: rgba(0, 0, 0, 0.1);
        }
    }

    a {
        color: #7DC314;
        -webkit-transition: all .2s ease-in-out;
        -moz-transition: all .2s ease-in-out;
        transition: all .2s ease-in-out;
    }

    a:hover,
    a:focus {
        text-decoration: none;
        color: #000;
    }

    .list-inline > li {
        display: inline-block !important;
        padding-right: unset !important;
        padding-left: unset !important;
        margin-right: 3px;
    }

    .ha_title {
        color: #7DC314 !important;
    }
</style>
