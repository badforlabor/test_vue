import Vue from 'vue'
import Trend from 'vuetrend';
Vue.use(Trend)
import App from './App.vue'

Vue.config.productionTip = false

new Vue({
  render: h => h(App)
}).$mount('#app')
