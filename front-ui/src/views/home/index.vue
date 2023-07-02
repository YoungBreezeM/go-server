<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { Login } from '@/api/index'
// import { computed } from 'vue'
import { useAuthStore, } from '@/store'
import { NCard, NAvatar } from 'naive-ui'
const route = useRoute()
const router = useRouter()
const { key } = route.query
const authSt = useAuthStore()
// const userStore = useUserStore()
// //
// const user = computed(() => userStore.userInfo)
// const toChat = () => {
//   router.push({
//     path: `/chat/${user.value.openId}`,
//   })
// }
// console.log(route.query)
if (key) {
  authSt.removeToken()
  Login<string>(key).then(res => {
    if (res.status == 0) {
      authSt.setToken(res.data)
      router.push({
        path: '/home'
      })
    }
  })
}
const items = [
  {
    type: 1,
    path: "/chat"
  },
  {
    type: 2,
  },
  {
    type: 3,
    path: "https://plus.web.xxyx.ltd/#/?uid=335131"
  }
]
const toTarget = (index) => {
  if (items[index].type == 1) {
    router.push({
      path: items[index].path
    })
  }
  if (items[index].type == 2) {
    window.$message.warning("当前功能尚未开放")
  }
  if (items[index].type == 3) {
    window.open(items[index].path, "_blank")
  }
}
</script>

<template>
  <div class="flex flex-col w-full h-full">
    <main class="flex-1 overflow-hidden">
      <n-card title="AI 聊天" class="dark:bg-[#24272e]" :bordered="false">
        <div id="scrollRef" ref="scrollRef" class="h-full overflow-hidden overflow-y-auto flex justify-left  flex-wrap">
          <div @click="toTarget(0)"
            class="cursor-pointer w-32 h-32 shadow-lg rounded-lg bg-green-400 shadow-lg shadow-green-500/50 flex justify-center items-center flex-wrap  ml-5">
            <div class="w-1/2 h-1/2">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32">
                <path
                  d="M29.71,13.09A8.09,8.09,0,0,0,20.34,2.68a8.08,8.08,0,0,0-13.7,2.9A8.08,8.08,0,0,0,2.3,18.9,8,8,0,0,0,3,25.45a8.08,8.08,0,0,0,8.69,3.87,8,8,0,0,0,6,2.68,8.09,8.09,0,0,0,7.7-5.61,8,8,0,0,0,5.33-3.86A8.09,8.09,0,0,0,29.71,13.09Zm-12,16.82a6,6,0,0,1-3.84-1.39l.19-.11,6.37-3.68a1,1,0,0,0,.53-.91v-9l2.69,1.56a.08.08,0,0,1,.05.07v7.44A6,6,0,0,1,17.68,29.91ZM4.8,24.41a6,6,0,0,1-.71-4l.19.11,6.37,3.68a1,1,0,0,0,1,0l7.79-4.49V22.8a.09.09,0,0,1,0,.08L13,26.6A6,6,0,0,1,4.8,24.41ZM3.12,10.53A6,6,0,0,1,6.28,7.9v7.57a1,1,0,0,0,.51.9l7.75,4.47L11.85,22.4a.14.14,0,0,1-.09,0L5.32,18.68a6,6,0,0,1-2.2-8.18Zm22.13,5.14-7.78-4.52L20.16,9.6a.08.08,0,0,1,.09,0l6.44,3.72a6,6,0,0,1-.9,10.81V16.56A1.06,1.06,0,0,0,25.25,15.67Zm2.68-4-.19-.12-6.36-3.7a1,1,0,0,0-1.05,0l-7.78,4.49V9.2a.09.09,0,0,1,0-.09L19,5.4a6,6,0,0,1,8.91,6.21ZM11.08,17.15,8.38,15.6a.14.14,0,0,1-.05-.08V8.1a6,6,0,0,1,9.84-4.61L18,3.6,11.61,7.28a1,1,0,0,0-.53.91ZM12.54,14,16,12l3.47,2v4L16,20l-3.47-2Z" />
              </svg>
            </div>
            <span class="dark:text-black">ChatGtp-3.5</span>
          </div>
          <div @click="toTarget(1)"
            class="cursor-pointer w-32 h-32 shadow-lg rounded-lg bg-green-400 shadow-lg shadow-green-500/50 flex justify-center items-center flex-wrap  ml-5">
            <div class="w-1/2 h-1/2">
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32">
                <path
                  d="M29.71,13.09A8.09,8.09,0,0,0,20.34,2.68a8.08,8.08,0,0,0-13.7,2.9A8.08,8.08,0,0,0,2.3,18.9,8,8,0,0,0,3,25.45a8.08,8.08,0,0,0,8.69,3.87,8,8,0,0,0,6,2.68,8.09,8.09,0,0,0,7.7-5.61,8,8,0,0,0,5.33-3.86A8.09,8.09,0,0,0,29.71,13.09Zm-12,16.82a6,6,0,0,1-3.84-1.39l.19-.11,6.37-3.68a1,1,0,0,0,.53-.91v-9l2.69,1.56a.08.08,0,0,1,.05.07v7.44A6,6,0,0,1,17.68,29.91ZM4.8,24.41a6,6,0,0,1-.71-4l.19.11,6.37,3.68a1,1,0,0,0,1,0l7.79-4.49V22.8a.09.09,0,0,1,0,.08L13,26.6A6,6,0,0,1,4.8,24.41ZM3.12,10.53A6,6,0,0,1,6.28,7.9v7.57a1,1,0,0,0,.51.9l7.75,4.47L11.85,22.4a.14.14,0,0,1-.09,0L5.32,18.68a6,6,0,0,1-2.2-8.18Zm22.13,5.14-7.78-4.52L20.16,9.6a.08.08,0,0,1,.09,0l6.44,3.72a6,6,0,0,1-.9,10.81V16.56A1.06,1.06,0,0,0,25.25,15.67Zm2.68-4-.19-.12-6.36-3.7a1,1,0,0,0-1.05,0l-7.78,4.49V9.2a.09.09,0,0,1,0-.09L19,5.4a6,6,0,0,1,8.91,6.21ZM11.08,17.15,8.38,15.6a.14.14,0,0,1-.05-.08V8.1a6,6,0,0,1,9.84-4.61L18,3.6,11.61,7.28a1,1,0,0,0-.53.91ZM12.54,14,16,12l3.47,2v4L16,20l-3.47-2Z" />
              </svg>
            </div>
            <span class="dark:text-black">ChatGtp-4</span>
          </div>
        </div>
      </n-card>
      <n-card title="科学上网" class="dark:bg-[#24272e] " :bordered="false">
        <div id="scrollRef" ref="scrollRef" class="h-full overflow-hidden overflow-y-auto flex justify-left  flex-wrap">
          <div @click="toTarget(1)"
            class="cursor-pointer w-32 h-32 shadow-lg rounded-lg bg-green-400 shadow-lg shadow-green-500/50 flex justify-center items-center flex-wrap  ml-5">
            <div class="w-1/2 h-1/2">
              <img src='@/assets/clash.png'>
            </div>
            <span class="dark:text-black">Clash-v2.1.1</span>
          </div>
        </div>
      </n-card>
      <n-card title="薅羊毛" class="dark:bg-[#24272e]" :bordered="false">
        <div id="scrollRef" ref="scrollRef" class="h-full overflow-hidden overflow-y-auto flex justify-left  flex-wrap">
          <div @click="toTarget(2)"
            class="cursor-pointer w-32 h-32 shadow-lg rounded-lg bg-green-400 shadow-lg shadow-green-500/50 flex justify-center items-center flex-wrap  ml-5">
            <n-avatar round size="large" :style="{
              color: 'white',
              backgroundColor: '#24272e',
            }">
              霸王
            </n-avatar>
            <span class="dark:text-black">晓晓优选-霸王餐</span>
          </div>
        </div>
      </n-card>

    </main>
  </div>
</template>
