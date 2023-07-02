<script setup lang='ts'>
// import { computed } from 'vue'
import { NLayout, NLayoutContent, NModal, NCard, NStep, NSteps, NSpace } from 'naive-ui'
import HeaderComponent from './Header/index.vue'
import { useBasicLayout } from '@/hooks/useBasicLayout'
import type { MenuOption } from 'naive-ui'
import { h } from 'vue'
import { useMenuStore } from '@/store'
import { RouterLink } from "vue-router"
import { ref } from 'vue'
import { computed } from 'vue'


// import {
//     BookOutline as BookIcon,
//     PersonOutline as PersonIcon,
//     WineOutline as WineIcon,
//     HomeOutline as HomeIcon
// } from '@vicons/ionicons5'

// function renderIcon(icon: Component) {
//     return () => h(NIcon, null, { default: () => h(icon) })
// }
const menuStorage = useMenuStore()
// const router = useRouter()
// const router = useRouter()
// const appStore = useAppStore()
// const chatStore = useChatStore()
// const authStore = useAuthStore()

// router.replace({ name: 'Chat', params: { uuid: chatStore.active } })

const { isMobile } = useBasicLayout()

// const collapsed = computed(() => appStore.siderCollapsed)
const activeKey = computed(() => menuStorage.menuInfo.Active)
// const needPermission = computed(() => !!authStore.session?.auth && !authStore.token)

// const getMobileClass = computed(() => {
//     if (isMobile.value)
//         return ['rounded-none', 'shadow-none']
//     return ['border', 'rounded-md', 'shadow-md', 'dark:border-neutral-800']
// })

// const getContainerClass = computed(() => {
//     return [
//         'h-full',
//         { 'pl-0': !isMobile.value && !collapsed.value },
//     ]
// })

const menuOptions: MenuOption[] = [
    {
        label: () => h(
            RouterLink, {
            to: {
                name: 'home',
            }
        },
            { default: () => '首页' }
        ),
        key: '/home',


    },
    {
        label: () => h(
            RouterLink, {
            to: {
                name: 'chat',
            }
        },
            { default: () => 'AI 聊天' }
        ),
        key: '/chat',
    },
    // {
    //     label: 'AI 绘画',
    //     key: '/draw',

    // },
    {
        label: '科学上网',
        key: '/vpn',

    }
]
const isLogin = ref<boolean>(false)

const Login = () => {
    isLogin.value = true
}

const closeLogin = () => {
    isLogin.value = false
}
</script>

<template>
    <div class="h-full dark:bg-[#24272e] transition-all">
        <n-modal v-model:show="isLogin" @after-leave="closeLogin">
            <n-card style="width: 400px;" class="h-1/2" :bordered="true" size="huge" role="dialog" aria-modal="true">
                <p class="text-xl text-center mb-5 mt-1">登录</p>
                <div class="mx-auto  flex justify-center items-center" :class="[isMobile ? 'w-full' : 'lg:w-1/4']">
                    <div class="flex justify-center items-center flex-wrap flex-column">
                        <n-steps vertical>
                            <n-step title="微信扫码关注公众号" description="扫码关注公众号获得积分">
                                <img src="@/assets/qrcode.jpg" class="w-32 h-32 object-fit m-5">
                            </n-step>
                            <n-step title="发送: [链接],获取链接" />
                        </n-steps>
                    </div>
                </div>
            </n-card>
        </n-modal>
        <NSpace class="h-full" vertical size="large">
            <NLayout class="dark:bg-[#24272e]" position="absolute">
                <NLayoutHeader>
                    <HeaderComponent :is-mobile="isMobile" :active="activeKey" :menu-options="menuOptions"
                        @to-login="Login" />
                </NLayoutHeader>
                <NLayoutContent class="dark:bg-[#24272e]" :native-scrollbar="true">
                    <RouterView v-slot="{ Component, route }">
                        <component :is="Component" :key="route.fullPath" />
                    </RouterView>
                </NLayoutContent>
            </NLayout>

        </NSpace>
    </div>
</template>