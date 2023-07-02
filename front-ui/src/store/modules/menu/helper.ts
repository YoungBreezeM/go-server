import { ss } from '@/utils/storage'

const LOCAL_NAME = 'menuStorage'

export interface MenuInfo {
  Active: string
}

export interface MenuState {
  menuInfo: MenuInfo
}

export function defaultSetting(): MenuState {
  return {
    menuInfo: {
      Active: "/home"
    },
  }
}

export function getLocalState(): MenuState {
  const localSetting: MenuState | undefined = ss.get(LOCAL_NAME)
  return { ...defaultSetting(), ...localSetting }
}

export function setLocalState(setting: MenuState): void {
  ss.set(LOCAL_NAME, setting)
}
