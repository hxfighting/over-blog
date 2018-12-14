import VueAwesomeSwiper from 'vue-awesome-swiper'

// require styles
import 'swiper/dist/css/swiper.css'// import something here

// leave the export, even if you don't use it
export default ({ app, router, Vue }) => {
  Vue.use(VueAwesomeSwiper)
}
