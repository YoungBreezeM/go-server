<script setup lang="ts">


import { ref } from 'vue';
import { HelloRequest } from './api/v1/hello_pb';
import { HelloWorldServiceClient } from './api/v1/HelloServiceClientPb';
const v = ref("")
const client = new HelloWorldServiceClient('http://localhost:8088');

// 发起gRPC请求
const request = new HelloRequest();

// 设置请求参数
// ...

const setValue = (e: any) => {
  request.setName(e.target.value)
  client.sayHello(request, {}, (error, response) => {
    if (error) {
      console.error(error);
    } else {
      v.value = response.getMessage()
    }
  });
}
</script>

<template>
  <div>
    <input @change="setValue" />
    <h1>{{ v }}</h1>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
