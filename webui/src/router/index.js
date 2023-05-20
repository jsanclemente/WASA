import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import NoPageFound from '../views/NoPageFound.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'login',
			component: () => import('../views/Login.vue')
		},

		{
			path: '/myAccount',
			component: () => import('../views/MyAccountLayout.vue'),
			children: [
				{
					path: '/myAccount/home',
					component: () => import('../views/HomeView.vue'),
					name: 'home'
				},
				{
					path: '/myAccount/profile/:id',
					component: () => import('../views/ProfileView.vue'),
					name: 'profile',
					props: true
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
