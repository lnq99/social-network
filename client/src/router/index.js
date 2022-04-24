import { createRouter, createWebHistory } from 'vue-router'
import store from '../store/index.js'
import Main from '../views/Main.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Main,
    children: [
      {
        path: '',
        name: 'NewsFeed',
        component: () => import('../views/NewsFeed.vue'),
      },
      {
        path: 'notif',
        name: 'Notification',
        component: () => import('../views/Notification.vue'),
      },
      {
        path: 'search',
        name: 'Search',
        component: () => import('../views/Search.vue'),
      },
      {
        path: 'profile/:id',
        name: 'Profile',
        component: () => import('../views/Profile.vue'),
      },
      {
        path: 'photo/:id',
        name: 'Photo',
        component: () => import('../views/Photo.vue'),
      },
    ],
    meta: {
      requiredAuth: true,
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/Login.vue'),
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: () => import('../views/SignUp.vue'),
  },
  {
    path: '/logout',
    beforeEnter(to, from, next) {
      store.dispatch('logout').then((res) => {
        next('/login')
      })
    },
  },
  { path: '/:pathMatch(.*)*', name: 'not-found', redirect: '/' },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

// router.beforeEach((to, from, next) => {
//   if (to.meta.requireAuth) {
//     if (store.getters.isAuthenticated) next()
//     else {
//       store.dispatch('login').then((r) => {
//         if (r) next()
//         else next('/login')
//       })
//     }
//   } else {
//     next()
//   }
// })

router.beforeEach((to, from, next) => {
  if (to.meta.requiredAuth && !store.getters.isAuthenticated) {
    store.dispatch('login').then((r) => {
      if (r) next()
      else next('/login')
    })
  } else {
    next()
  }
})

export default router
