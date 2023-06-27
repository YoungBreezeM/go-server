<script setup lang="ts">
import { NConfigProvider, NMessageProvider } from 'naive-ui'

import { NaiveProvider, NaiveMessage } from '@/components/common'
import { useTheme } from '@/hooks/useTheme'
import { useLanguage } from '@/hooks/useLanguage'

const { theme, themeOverrides } = useTheme()
const { language } = useLanguage()
</script>

<template>
  <NConfigProvider class="h-full" :theme="theme" :theme-overrides="themeOverrides" :locale="language">
    <NMessageProvider>
      <NaiveMessage></NaiveMessage>
    </NMessageProvider>
    <NaiveProvider>
      <RouterView />
    </NaiveProvider>
  </NConfigProvider>
</template>
