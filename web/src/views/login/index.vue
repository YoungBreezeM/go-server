<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { GetUserInfo, LoginWeiXin } from '@/api/index'
import { Status } from '@/constant/Code'
import { useUserStore } from '@/store'
const route = useRoute()
const router = useRouter()
const userStore = useUserStore()
const { code } = route.query

interface AuthInfo {
  access_token: string
  openid: string
}

interface UserInfo {
  city: string
  country: string
  headimgurl: string
  language: string
  nickname: string
  privilege: string []
  province: string
  sex: number
}

if (code) {
  LoginWeiXin<AuthInfo>(code).then((res) => {
    if (res.status === Status.SUCCESS) {
      GetUserInfo<UserInfo>(res.data.access_token, res.data.openid).then((res) => {
        if (res.status === Status.SUCCESS) {
          userStore.updateUserInfo({
            avatar: res.data.headimgurl,
            name: res.data.nickname,
          })
          //
          router.push({
            path: '/home',
          })
        }
      })
    }
  })
}
const login = () => {
  // 获取当前页面地址作为回调地址，并且对地址进行urlEncode处理
  const local = 'http://192.168.31.202:3000'
  // 获取公众号appid
  const appid = 'wxd1db01c7241edc28'
  // 跳转到授权页面
  window.location.href
            = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=${
             appid
             }&redirect_uri=${
             local
             }&response_type=code&scope=snsapi_userinfo&state=1#wechat_redirect`
}
</script>

<template>
  <button @click="login">
    登录
  </button>
</template>
