<template>
	<div class="home">
		<img alt="Vue logo" src="../assets/logo.png">
		<table>
			<thead>
				<tr>
					<th class="id">ID</th>
					<th class="content">Content</th>
					<th class="status">Status</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="todo in todoList" :key="todo.id">
					<th><router-link :to="{name: 'todoEdit', params: {id: todo.id}}">{{ todo.id }}</router-link></th>
					<td><router-link :to="{name: 'todoEdit', params: {id: todo.id}}">{{ todo.Content }}</router-link></td>
					<td><router-link :to="{name: 'todoEdit', params: {id: todo.id}}">{{todo.status}}</router-link></td>
				</tr>
			</tbody>
		</table>
		<router-link to="/new">Add Todo</router-link>
	</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import ApiService from '@/services/ApiService'
import ResponseData from '@/types/ResponseData'
import {TodoItem} from '@/types/TodoItem'

@Component({
})
export default class TodoList extends Vue {
	// todo一覧
	public todoList: TodoItem[] = []

	public retrieveTodoList() {
		ApiService.getAll()
		.then((response: ResponseData) => {
			this.todoList = response.data;
			console.log(response.data);
			console.log(this.todoList);
			
		})
		.catch((e: Error) => {
			console.log(e);
		});
	}
	public created() {
		this.retrieveTodoList();
	}
}
</script>

<style>
* {
	box-sizing: border-box;
}

#app {
	max-width: 640px;
	margin: 0 auto;
}

table {
	width: 100%;
	border-collapse: collapse;
}

thead th {
	border-bottom: 2px solid #0099e4; /*#d31c4a */
	color: #0099e4;
}

th,
th {
	padding: 0 8px;
	line-height: 40px;
}

thead th.id {
	width: 50px;
}

thead th.status {
	width: 100px;
}

thead th.button {
	width: 60px;
}

tbody td.button, tbody td.status {
	text-align: center;
}

tbody tr td,
tbody tr th {
	border-bottom: 1px solid #ccc;
	transition: All 0.4s;
}

tbody tr.Done td,
tbody tr.Done th {
	background: #f8f8f8;
	color: #bbb;
}

tbody tr:hover td,
tbody tr:hover th {
	background: #f4fbff;
}

button {
	border: none;
	border-radius: 20px;
	line-height: 24px;
	padding: 0 8px;
	background: #0099e4;
	color: #fff;
	cursor: pointer;
}
</style>