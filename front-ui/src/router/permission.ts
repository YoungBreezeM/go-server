import { getToken } from '@/store/modules/auth/helper'
import type { Router } from 'vue-router'

export function setupPageGuard(router: Router) {
  router.beforeEach(async (to, from, next) => {
    if (to.path == '/login' || to.path == "/404" || to.path == "/500") {
      next()
      return
    }
    //
    const token = getToken()
    if (token) {
      next()
      return
    } else {
      next({ name: '404' })
    }
    // const authStore = useAuthStoreWithout()
    // if (!authStore.session) {
    //   try {
    //     const data = await authStore.getSession()
    //     if (String(data.auth) === 'false' && authStore.token)
    //       authStore.removeToken()
    //     if (to.path === '/500')
    //       next({ name: 'Root' })
    //     else
    //       next()
    //   }
    //   catch (error) {
    //     if (to.path !== '/500')
    //       next({ name: '500' })
    //     else
    //       next()
    //   }
    // }
    // else {
    //   next()
    // }
  })
}
