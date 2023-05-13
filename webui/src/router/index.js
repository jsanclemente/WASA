import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NoPageFound from '../views/NoPageFound.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			component: () => import('../views/Login.vue')
		},

		{
			path: '/myAccount',
			component: () => import('../views/MyAccountLayout.vue'),
			children: [
				{
					path: 'home',
					component: () => import('../views/HomeView.vue'),
					name: 'home'
				},
				{
					path: 'profile',
					component: () => import('../views/ProfileView.vue')
				}	
			]
		},
		
		{
			path: '/:pathMatch(.*)*',
			component: () => import('../views/NoPageFound.vue')
		},
	]
})

export default router
