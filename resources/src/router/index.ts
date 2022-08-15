import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import TodoList from '../components/TodoList.vue'
import TodoEdit from '../components/TodoEdit.vue'
import TodoCreate from '../components/TodoCreate.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
	{
		path: '/',
		name: 'todoList',
		component: TodoList
	},
	{
		path: '/about',
		name: 'about',
		// route level code-splitting
		// this generates a separate chunk (about.[hash].js) for this route
		// which is lazy-loaded when the route is visited.
		component: () => import(/* webpackChunkName: "about" */ '../components/AboutView.vue')
	},
	{
		path: '/new',
		name: 'todoCreate',
		component: TodoCreate,
	},
	{
		path: `/edit/:id`,
		name: 'todoEdit',
		component: TodoEdit,
	}
]

const router = new VueRouter({
	routes
})

export default router
