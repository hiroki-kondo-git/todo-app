<template>
	<div>
		<p>hello</p>
		<p>{{todoItem}}</p>
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
		status: "todo"
	}
	public getTodo(id: string) {
		ApiService.get(id)
		.then((response: ResponseData) => {
			this.todoItem = response.data;
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