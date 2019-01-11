import Vue from 'vue'
import Router from 'vue-router'
import Discovery from '@/components/Discovery'
import Search from '@/components/Search'
import Manage from '@/components/Manage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: '搜索',
      component: Search
    },
    {
      path: '/Search',
      name: '搜索',
      component: Search
    },
    {
      path: '/Discovery',
      name: '发现',
      component: Discovery
    },
    {
      path: '/Manage',
      name: '管理',
      component: Manage
    }
  ]
})
