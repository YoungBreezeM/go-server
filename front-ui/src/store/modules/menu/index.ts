import { defineStore } from 'pinia'
import type { MenuState } from './helper'
import { getLocalState, } from './helper'

export const useMenuStore = defineStore('menu-store', {
  state: (): MenuState => getLocalState(),
  actions: {

  },
})
