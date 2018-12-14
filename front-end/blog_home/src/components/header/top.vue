<template>
    <div>
        <nav id="top">
            <div class="container">
                <div class="row">
                    <div class="col-6" title="Welcome">
                        <strong>Welcome !</strong>
                    </div>
                    <div class="col-6">
                        <ul class="list-inline top-link link">
                            <li><a title="首页" href="https://www.ohdata.top"><i class="fa fa-home"></i> 首页</a></li>
                            <li style="cursor: pointer"><a title="登录" class="oauth_login"><i class="fa fa-user"></i> 登录</a>
                            </li>
                            <li><a title="联系我" href="https://www.ohdata.top/category/8"><i class="fa fa-comments"></i>↓联系我↓</a>
                            </li>
                        </ul>
                    </div>
                </div>
            </div>
        </nav>
    </div>
</template>

<script>
export default {
  name: 'top'
}
</script>

<style scoped>
    .container {
        padding-right: 15px;
        padding-left: 15px;
        margin-right: auto;
        margin-left: auto;
        width: 1170px;
    }

    .list-inline {
        padding-left: 0;
        margin-left: -5px;
        list-style: none;
    }

    .list-inline > li {
        display: inline-block;
        padding-right: 5px;
        padding-left: 5px;
    }

    #top {
        background: #fff;
        padding: 10px 0 0;
        color: #000;
        border-bottom: 1px solid #ddd;
    }
    #top a {
        color: #000;
    }

    #top a:hover {
        color: #69BF03
    }

    .link a {
        display: inline-block;
        text-decoration: none;
        font-size: 12px;
        font-weight: bold;
        letter-spacing: 1px
    }

    .link li:after {
        content: '/';
        margin: 0 0 0 12px;
        display: inline-block;
        font-weight: bold;
    }

    .link li:last-child:after {
        content: '';
        margin-right: 0px;
        margin-left: 0;
        display: none
    }

    .link li ul {
        display: none
    }
    .link {
        text-align: right;
    }
</style>
