import Vue from 'vue'
// import Argon from '@/plugins/argon-kit'
import BootstrapVue from "bootstrap-vue"
import App from './App.vue'
import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap-vue/dist/bootstrap-vue.css"

Vue.config.productionTip = false
// Vue.use(Argon);
Vue.use(BootstrapVue)


new Vue({
  render: h => h(App),
}).$mount('#app')
