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
        import('../views/user/login/LoginView.vue')
  },
  {
    path: '/i-want-to-register',
    name: 'register',
    component: () =>
        import('../views/user/register/RegisterView.vue')
  },
  {
    path: '/record-query',
    name: 'record-query',
    component: () =>
        import('@/views/record/query/QueryView.vue')
  },
  {
    path: '/hardware-query-and-operate',
    name: 'hardware-query-and-operate',
    component: () =>
        import('@/views/hardware/query-and-operate/QueryAndOperateView.vue')
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
