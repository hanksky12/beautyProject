import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/main/home/HomeView.vue'
import store from '@/store/index.js'

Vue.use(VueRouter)


const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/login',
    name: 'login',
    component: () =>
        import('@/views/user/authentication/login/LoginView.vue')
  },
  {
    path: '/personal-register',
    name: 'register',
    component: () =>
        import('@/views/user/authentication/register/RegisterView.vue')
  },
  {
    path: '/server-average-record-query',
    name: 'server-average-record-query',
    component: () =>
        import('@/views/server/record-average/query/QueryView.vue')
  },
  {
    path: '/server-raw-record-query',
    name: 'server-raw-record-query',
    component: () =>
        import('@/views/server/record-raw/query/QueryView.vue')
  },
  {
    path: '/server-hardware-query-and-operate',
    name: 'server-hardware-query-and-operate',
    component: () =>
        import('@/views/server/hardware/query-and-operate/QueryAndOperateView.vue')
  },
  {
    path: '/user-raw-record-query',
    name: 'user-raw-record-query',
    component: () =>
        import('@/views/user/record-raw/query/QueryView.vue')
  },

  {
    path: '/user-mouse-action-query-and-operate',
    name: 'user-mouse-action-query-and-operate',
    component: () =>
        import('@/views/user/mouse-action/query-and-operate/QueryAndOperateView.vue')
  },
]



const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

console.log(process.env.NODE_ENV)
if (process.env.NODE_ENV === 'production' || process.env.NODE_ENV === 'test'){
  router.beforeEach((to, from, next) => {
    if (!store.state.authenticate.isAuthenticated && to.name !== 'login' && to.name !== 'register'){
    //沒有登入且不是登入頁面且不是註冊頁面
     next({ name: 'login' })
    }
    else {
      next()
    }
  })
}


router.afterEach(() => {
  document.title = '美麗後台';
  store.commit("form/resetFormData");
  store.commit("crudOperate/resetOperate");
});


export default router
