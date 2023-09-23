import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Acceuil',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../pages/Acceuil.vue')
    },
    {
      path: '/profile',
      name: 'profile',
      component: () => import('../pages/Profile.vue'),
      meta: {
        requiresAuth: true // Marquez cette route comme nécessitant une authentification
      }
    },
    {
      path: '/uploadImage',
      name: 'PageUploadImage',
      component: () => import('../pages/Dashbord/PageUploadImage.vue')
    },
    {
      path: '/connection',
      name: 'Login',
      component: () => import('../pages/connections/Login.vue')
    },
    {
      path: '/inscription',
      name: 'Signup',
      component: () => import('../pages/connections/Signup.vue')
    },
    {
      path: '/boutique',
      name: 'Boutique',
      component: () => import('../pages/Boutique.vue')
    },
    {
      path: '/collections',
      name: 'Collections',
      component: () => import('../pages/Collection.vue')
    },          
  ]
})

export default router
