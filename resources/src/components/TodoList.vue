<template>
	<div class="home">
		<img alt="Vue logo" src="../assets/logo.png">
		<HelloWorld msg="Todo List"/>
		<p>{{greetText}}</p>
		<label v-for="[status, text] in Array.from(labels)" :key="status">
			<input type="radio" v-model="currentStatus" :value="status">
			{{text}}
		</label>
		{{ filteredTodoList }} 件
		<table>
			<thead>
				<tr>
					<th class="id">ID</th>
					<th class="content">Content</th>
					<th class="status">Status</th>
					<th class="button">-</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="todo in filteredTodoList" :key="todo.id">
					<th>{{ todo.id }}</th>
					<td>{{ todo.Content }}</td>
					<td>{{todo.status}}</td>
					<td class="status">
						<button @click="toggleStatus(todo)">
							{{ labels.get(todo.status )}}
						</button>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import ApiService from '@/services/ApiService'
import ResponseData from '@/types/ResponseData'
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src
import {Status, TodoItem} from '@/types/TodoItem'

@Component({
	components: {
		HelloWorld,
	},
})
export default class TodoList extends Vue {
	public greetText: string = "Hello";
	// todo一覧
	public todoList: TodoItem[] = []

	// status絞り込み
	public labels = new Map<Status, string>([
		[Status.all, '全件'],
		[Status.todo, '未着手'],
		[Status.doing, '対応中'],
		[Status.done, '完了']
	])
	// 表示したいstatus
	public currentStatus: Status = Status.all

	// currentStatusを元に絞り込み
	public get filteredTodoList() {
		return this.todoList.filter(t =>
		this.currentStatus === Status.all ? true : this.currentStatus === t.status)
	}

	// todoのstatus変更
	public toggleStatus(todo: TodoItem) {
		switch (todo.status) {
			case Status.todo:
				todo.status = Status.doing;
				break;
			case Status.doing:
				todo.status = Status.done;
				break;
			case Status.done:
				todo.status = Status.todo;
				break;
			default:
				break;
		}
	}

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