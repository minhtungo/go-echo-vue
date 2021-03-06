import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Post from './views/Post.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/post',
      name: 'post',
      component: Post
    }
  ]
})