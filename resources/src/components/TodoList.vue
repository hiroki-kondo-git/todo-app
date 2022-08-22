<template>
	<div class="home">
		<img alt="Vue logo" src="../assets/logo.png">
		<b-table striped hover :items="todoList">
			<template v-slot:cell(id)="data">
				<router-link :to="{name: 'todoDetails', params: {id: data.value}}">{{ data.value }}</router-link>
			</template>
		</b-table>
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