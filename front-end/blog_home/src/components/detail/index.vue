<template>
    <div>
        <header-top></header-top>
        <header-middle></header-middle>
        <header-bottom></header-bottom>
        <div id="page-content" class="container">
            <div class="row">
                <div class="col-md-8">
                    <v-left></v-left>
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
import headerTop from '../header/top'
import headerMiddle from '../header/middle'
import headerBottom from './detailHeaderBottom'
import left from './detail'
import right from '../content/rightSide'
import footer from '../footer/index'

export default {
  name: 'index',
  components: {
    headerTop,
    headerMiddle,
    headerBottom,
    'v-left': left,
    'v-right': right,
    'v-footer': footer

  }
}
</script>

<style scoped>
    #page-content {
        margin-top: 20px;
    }
</style>
