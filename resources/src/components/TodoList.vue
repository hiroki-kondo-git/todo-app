<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Todo List"/>
    <p>{{greetText}}</p>
    <p>{{todoList}}</p>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import ApiService from '@/services/ApiService'
import ResponseData from '@/types/ResponseData'
import HelloWorld from '@/components/HelloWorld.vue'; // @ is an alias to /src
import TodoItem from '@/types/TodoItem'

@Component({
  components: {
    HelloWorld,
  },
})
export default class TodoList extends Vue {
  public greetText: string = "Hello";
  public todoList: TodoItem[] = []
  public title = ''

  public retrieveTodoList() {
    ApiService.getAll()
    .then((response: ResponseData) => {
      this.todoList = response.data;
      console.log(response.data);
      
    })
    .catch((e: Error) => {
      console.log(e);
    });
  }
  mounted() {
    this.retrieveTodoList();
  }
}
</script>
