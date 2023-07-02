<script lang="ts" setup>
import { computed, } from 'vue'
import { HoverButton, SvgIcon, } from '@/components/common'
import { useAppStore, useAuthStore } from '@/store'
import { NMenu, NInput, NButton, NBadge, NAvatar } from "naive-ui"
import type { MenuOption } from 'naive-ui'
import { ref, defineAsyncComponent } from 'vue'
interface Props {
  isMobile: boolean
  menuOptions: MenuOption[],
  active: string
}

interface Emit {
  (e: 'toLogin'): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emit>()

const appStore = useAppStore()
const authStore = useAuthStore()
const collapsed = computed(() => appStore.siderCollapsed)
const isLogin = computed(() => authStore.token)
const activeKey = ref<string>()
activeKey.value = props.active

//
function handleUpdateCollapsed() {
  appStore.setSiderCollapsed(!collapsed.value)
}



// function onScrollToTop() {
//   const scrollRef = document.querySelector('#scrollRef')
//   if (scrollRef)
//     nextTick(() => scrollRef.scrollTop = 0)
// }
const Setting = defineAsyncComponent(() => import('@/components/common/Setting/index.vue'))
const show = ref(false)
const toLogin = () => {
  emit("toLogin")
}

</script>

<template>
  <header
    class="sticky top-0 left-0 right-0 z-30 border-b dark:border-neutral-800 bg-white/80 dark:bg-black/80 backdrop-blur">
    <div class="relative flex items-center justify-between min-w-0 overflow-hidden h-16 ">
      <div class="flex items-center w-2/3">
        <button v-if="props.isMobile" class="flex items-center justify-center w-11 h-11" @click="handleUpdateCollapsed">
          <SvgIcon v-if="collapsed" class="text-2xl" icon="ri:align-justify" />
          <SvgIcon v-else class="text-2xl" icon="ri:align-right" />
        </button>
        <span class="text-xl pl-5 sm:pl-5">FreeAiBox</span>
        <span v-if="!props.isMobile">
          <n-menu class="text-base" v-model:value="activeKey" mode="horizontal" :options="props.menuOptions"
            default-value="1" />
        </span>
        <span v-if="!props.isMobile" class="w-1/2 pl-10">
          <n-input type="text" size="large" placeholder="搜索" />
        </span>

      </div>
      <div class="flex items-center space-x-2 mr-5">
        <HoverButton>
          <n-badge value="999">
            <SvgIcon class="text-2xl" icon="system-uicons:bell-ringing" />
          </n-badge>

        </HoverButton>
        <HoverButton @click="show = true">
          <span class="text-xl text-[#4f555e] dark:text-white">
            <SvgIcon icon="ri:settings-4-line" />
          </span>
        </HoverButton>
        <Setting v-if="show" v-model:visible="show" />
        <HoverButton v-if="isLogin">
          <span class="text-xl text-[#4f555e] dark:text-white">
            <n-avatar round size="large" :style="{
              color: 'white',
              backgroundColor: '#24272e',
            }">
              U
            </n-avatar>
            <!-- <UserAvatar /> -->
          </span>
        </HoverButton>
        <NButton v-else type="primary" ghost @click="toLogin">
          登录
        </NButton>
      </div>
    </div>
  </header>
</template>
