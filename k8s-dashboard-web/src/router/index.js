import Vue from 'vue'
import Router from 'vue-router'
import Framework from '../components/Framework'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Framework',
      component: Framework
    }
  ]
})
