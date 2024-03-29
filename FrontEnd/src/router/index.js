import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import UsersView from '../views/UsersView.vue'
import PostView from '../views/PostView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/users',
      name: 'users',
      component: UsersView
    },
    {
      path: '/posts',
      name: 'posts',
      component: PostView
    },
    {
      path: '/login',
      name: 'login',
      component : () => import ('../views/Login.vue')
    },
    {
      path: '/register',
      name: 'register',
      component : () => import ('../views/Register.vue')
    },
    {
      path: '/profile',
      name: 'profile',
      component : () => import ('../views/Profile.vue')
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

// router.beforeEach(() => {
//   if (this.$store)
// })

export default router
