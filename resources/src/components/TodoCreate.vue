<template>
	<div class="submit-form">
		<div v-if="!submitted">
			<div class="form-group">
				<label for="content">content</label>
				<input type="text" class="form-control" id="content" required v-model="todoItem.Content" name="content">
			</div>
			<button @click="submitTodoItem" class="btn btn-success">Submit</button>
		</div>
		<div v-else>
			<h4>submitted successfully!</h4>
			<button class="btn btn-success" @click="addTodoItem">Add</button>
		</div>
	</div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import ApiService from '@/services/ApiService'
import ResponseData from '@/types/ResponseData'
import {TodoItem} from '@/types/TodoItem'

@Component({
})
export default class TodoCreate extends Vue {
	public todoItem: TodoItem = {
		id: 0,
		Content: "",
		status: "todo"
	}
	public submitted: boolean = false

	public submitTodoItem() {
		let data = {
			content: this.todoItem.Content,
			status: this.todoItem.status,
		};
		console.log(data);
		ApiService.create(data)
		.then((response: ResponseData) => {
			this.todoItem.id = response.data.id;
			this.submitted = true;
		})
		.catch((e: Error) => {
			console.log(e);
		})
	}
	public addTodoItem() {
		this.submitted = false;
		this.todoItem = {
			id: 0,
			Content: '',
			status: 'todo'
		};
	}
}
</script>

<style>
.submit-form {
	max-width: 300px;
	margin: auto;
}
</style>