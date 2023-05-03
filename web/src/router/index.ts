// Composables
import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Guest.vue'),
  },
  // {
  //   path: '/login',
  //   name: 'Login',
  //   component: () => import('@/views/Login.vue'),
  // },
  {
    path: "/:pathMatch(.*)*",
    name: "Error",
    component: () => import( "@/views/errors/404.vue"),
  },
]

// const dynamicRoutes: Array<RouteRecordRaw> = [
//   {
//     path: '/index',
//     name: 'Overview',
//     component: () => import('@/views/Home.vue'),
//     meta: {
//       role: ['teacher', 'student', 'admin'],
//     }
//   },
//   {
//     path: '/profile',
//     name: 'Profile',
//     component: () => import('@/views/profile/Profile.vue'),
//     meta: {
//       role: ['teacher', 'student'],
//     }
//   },
//   {
//     path: '/meetplan',
//     name: 'MeetPlan',
//     component: () => import('@/views/meetplan/Profile.vue'),
//     meta: {
//       role: ['teacher', 'student'],
//     }
//   },
//   {
//     path: '/meetplanorder',
//     name: 'MeetPlanOrder',
//     component: () => import('@/views/meetplan/Order.vue'),
//     meta: {
//       role: ['teacher', 'student'],
//     }
//   }
// ]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    }
    return {top: 0}
  }
})

router.beforeEach((to) => {
})

export default router
