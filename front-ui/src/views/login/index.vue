<script setup lang="ts">
import { useRouter } from 'vue-router'

import { Login } from '@/api/index'
import { useAuthStore } from '@/store'
import { NInput, NSteps, NStep } from "naive-ui"
import { setToken } from '@/store/modules/auth/helper'
import { ref } from 'vue'
import { useBasicLayout } from "@/hooks/useBasicLayout"
const { isMobile } = useBasicLayout()
// console.log(isMobile)
// const route = useRoute()
const router = useRouter()
// const userStore = useUserStore()
const authSt = useAuthStore()
// const { code } = route.query
const key = ref<string>()
const ld = ref<boolean>(false)
// interface AuthInfo {
//   access_token: string
//   openid: string
// }

// interface UserInfo {
//   city: string
//   country: string
//   headimgurl: string
//   language: string
//   nickname: string
//   privilege: string[]
//   province: string
//   sex: number
//   Integral: number
// }
authSt.removeToken()
//
const submit = () => {
  ld.value = true
  Login(key.value).then((res: any) => {
    if (res.status == 0) {
      setToken(res.data)
      router.push({
        path: '/chat'
      })
      // GetUserInfo<UserInfo>().then(resp => {
      //   if (resp.status == 0) {
      //     useUserStore().updateUserInfo({
      //       integral: resp.data.Integral
      //     })
      //     ld.value = false
      //    

      //   }
      // })
    }
  })
}
</script>

<template>
  <div class="h-full flex  flex-wrap">
    <div class="lg:w-3/4  flex  justify-center items-center">
      <img src="@/assets/login.png" />
    </div>
    <div class="mx-auto bg-gray-50 flex justify-center items-center" :class="[isMobile ? 'w-full' : 'lg:w-1/4']">
      <div class="flex justify-center items-center flex-wrap flex-column">
        <n-steps vertical>
          <n-step title="微信扫码关注公众号" description="扫码关注公众号获得积分">
            <img src="@/assets/qrcode.jpg" class="w-48 h-48 m-5">
          </n-step>
          <n-step title="发送: [密钥],获取密钥登录" />
          <n-step title="登录">
            <n-input :loading="ld" class="mt-5" type="password" @keyup.enter.native="submit" v-model:value="key"
              show-password-on="mousedown" placeholder="密钥" :maxlength="32" />
          </n-step>
        </n-steps>
      </div>
    </div>

  </div>
</template>
