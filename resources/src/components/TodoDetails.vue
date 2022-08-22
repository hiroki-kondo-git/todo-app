<template>
	<div>
		<b-form @submit="updateTodoStatus" v-if="show">
			<b-form-group
				id="input-group-1"
				label="content:"
				label-for="input-1"
			>
				<b-form-input
					id="input-1"
					v-model="todoItem.Content"
					type="content"
					placeholder="Enter content"
					required
				></b-form-input>
			</b-form-group>
			<b-form-group id="input-group-2" label="status:" label-for="input-2">
				<b-form-select
					id="input-2"
					v-model="todoItem.status"
					:options="statusOptions"
					required
					size="sm"
				></b-form-select>
			</b-form-group>
			<b-button type="submit" variant="primary">Submit</b-button>
		</b-form>
	</div>
</template>

<script lang="ts">
import ApiService from '@/services/ApiService';
import { TodoItem } from '@/types/TodoItem';
import { Component, Vue } from 'vue-property-decorator';
import ResponseData from '@/types/ResponseData'
@Component

export default class TodoDetails extends Vue {
	public todoItem: TodoItem = {
		id: 0,
		Content: "",
		status: ""
	}
	public statusOptions: string[] = [
		'todo',
		'doing',
		'done',
	]
	public show: boolean = true
	public getTodo(id: string) {
		ApiService.get(id)
		.then((response: ResponseData) => {
			this.todoItem = response.data;
		})
		.catch((e: Error) => {
			console.log(e);
		})
	}
	public updateTodoStatus() {
		console.log("hello");
		let data = {
			content: this.todoItem.Content,
			status: this.todoItem.status,
		}
		ApiService.update(this.todoItem.id, data)
		.then((response: ResponseData) => {
			this.todoItem.status = response.data.status;
			this.todoItem.Content = response.data.Content
			this.getTodo(String(this.todoItem.id))
		})
		.catch((e: Error) => {
			console.log(e);
		})
	}
	public created() {
		this.getTodo(this.$route.params.id);
	}
}
</script>

<style>
	.card.mb-2 {
		margin: 0 auto;
	}
</style>