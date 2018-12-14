<template>
    <div>
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
    </div>
</template>

<script>
import left from './leftSide'
import right from './rightSide'
export default {
  name: 'index',
  components: {
    'v-left': left,
    'v-right': right
  }
}
</script>

<style scoped>
    #page-content {
        margin-top: 20px;
    }
    .container{
        padding-left: unset !important;
    }
</style>
