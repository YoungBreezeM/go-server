import type { AxiosProgressEvent, GenericAbortSignal } from 'axios'
import type { LocationQueryValue } from 'vue-router'
import { get, post } from '@/utils/request'
import { useAuthStore, useSettingStore } from '@/store'
import axios from 'axios'
export function fetchChatAPI<T = any>(
  prompt: string,
  options?: { conversationId?: string; parentMessageId?: string },
  signal?: GenericAbortSignal,
) {
  return post<T>({
    url: '/chat',
    data: { prompt, options },
    signal,
  })
}

export function fetchChatConfig<T = any>() {
  return post<T>({
    url: '/config',
  })
}

export function fetchChatAPIProcess<T = any>(
  params: {
    prompt: string
    options?: { conversationId?: string; parentMessageId?: string }
    signal?: GenericAbortSignal
    onDownloadProgress?: (progressEvent: AxiosProgressEvent) => void
  },
) {
  const settingStore = useSettingStore()
  const authStore = useAuthStore()

  let data: Record<string, any> = {
    prompt: params.prompt,
    options: params.options,
  }

  if (authStore.isChatGPTAPI) {
    data = {
      ...data,
      systemMessage: settingStore.systemMessage,
      temperature: settingStore.temperature,
      top_p: settingStore.top_p,
    }
  }

  return post<T>({
    url: '/chat-process',
    data,
    signal: params.signal,
    onDownloadProgress: params.onDownloadProgress,
  })
}

export function fetchSession<T>() {
  return post<T>({
    url: '/session',
  })
}

export function fetchVerify<T>(token: string) {
  return post<T>({
    url: '/verify',
    data: { token },
  })
}

export function LoginWeiXin<T>(code: LocationQueryValue[] | string) {
  return get<T>({
    url: `wx/auth/${code}`,
  })
}

export function GetWechatUserInfo<T>(access_token: string, openid: string) {
  return get<T>({
    url: `wx/user/${access_token}/${openid}`,
  })
}

export function Chat<T>(op: any) {
  return post<T>({
    url: '/v1/chat/completions',
    data: op,
  })
}

export function ChatGTP4<T>(op: any) {
  return post<T>({
    url: '/v2/chat/completions',
    data: op,
  })
}

export function GetQrCode(sceneId: string) {
  return axios({
    url: `/api/wx/qrCode/${sceneId}`,
    responseType: 'blob',
  })
}


export function CheckQrCodeIsScan<T>(sceneId: string) {
  return get<T>({
    url: `/watchQrCodeScan/${sceneId}`,
  })
}

export function Login<T>(key: string | undefined) {
  return get<T>({
    url: `/loginByKey?key=${key}`,
  })
}
export function GetUserInfo<T>() {
  return get<T>({
    url: `/user`,
  })
}