<template>
  <div>
      <v-top></v-top>
      <v-middle></v-middle>
  </div>

</template>

<script>
import top from './top'
import middle from './middle'
export default {
  name: 'index',
  components: {
    'v-top': top,
    'v-middle': middle
  }
}
</script>

<style scoped>

</style>
