<template>
  <div>
    <v-header></v-header>
    <router-view></router-view>
    <div id="page-content" class="container">
      <div class="row">
        <div class="col-md-8">
          <router-view name="h_content"></router-view>
        </div>
        <div class="col-md-4">
          <v-right></v-right>
        </div>
      </div>
    </div>
    <v-footer></v-footer>
  </div>
</template>

<script>
  import header from '../components/header/index'
  import footer from '../components/footer/index'
  import rightSide from '../components/content/rightSide'

  export default {
    name: "mainPage",
    components: {
      'v-header': header,
      'v-footer': footer,
      'v-right': rightSide
    }
  }
</script>

<style scoped>
  @import "../statics/css/bootstrap.min.css";

  #page-content {
    margin-top: 20px;
  }

  .container {
    padding-left: unset !important;
  }
</style>
