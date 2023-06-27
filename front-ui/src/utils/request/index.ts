import type { AxiosProgressEvent, AxiosResponse, GenericAbortSignal } from 'axios'
import request from './axios'


export interface HttpOption {
  url: string
  data?: any
  method?: string
  headers?: any
  onDownloadProgress?: (progressEvent: AxiosProgressEvent) => void
  signal?: GenericAbortSignal
  beforeRequest?: () => void
  afterRequest?: () => void
  responseType?: string
}

export interface Response<T = any> {
  data: T
  message: string | null
  status: number
}

function http<T = any>(
  { url, data, method, headers, onDownloadProgress, signal, beforeRequest, afterRequest, responseType }: HttpOption,
) {
  const successHandler = (res: AxiosResponse<Response<T>>) => {
    // const authStore = useAuthStore()

    // if (res.data.status === 'Success' || typeof res.data === 'string')
    //   return res.data

    // if (res.data.status === 'Unauthorized') {
    //   authStore.removeToken()
    //   window.location.reload()
    // }
    return res.data
  }

  const failHandler = (error: Response<Error>) => {
    afterRequest?.()
    console.log(error)
    if (error.response.status == 401) {

      window.$message?.warning(error.response.data.message)
      setTimeout(() => {
        window.location.href = "/login"
      }, 3000)

    }
    if (error.response.status == 402) {
      window.$message?.warning("当前积分不足请充值")
    }
    if (error.response.status == 501) {
      window.$message?.warning(error.response.data.message)
    }
    // throw new Error(error?.message || 'Error')
  }

  beforeRequest?.()

  method = method || 'GET'

  const params = Object.assign(typeof data === 'function' ? data() : data ?? {}, {})

  return method === 'GET'
    ? request.get(url, { params, signal, onDownloadProgress }).then(successHandler, failHandler)
    : request.post(url, params, { headers, signal, onDownloadProgress }).then(successHandler, failHandler)
}

export function get<T = any>(
  { url, data, method = 'GET', onDownloadProgress, signal, beforeRequest, afterRequest, responseType }: HttpOption,
): Promise<Response<T>> {
  return http<T>({
    url,
    method,
    data,
    onDownloadProgress,
    signal,
    beforeRequest,
    afterRequest,
    responseType,
  })
}

export function post<T = any>(
  { url, data, method = 'POST', headers, onDownloadProgress, signal, beforeRequest, afterRequest, responseType }: HttpOption,
): Promise<Response<T>> {
  return http<T>({
    url,
    method,
    data,
    headers,
    onDownloadProgress,
    signal,
    beforeRequest,
    afterRequest,
    responseType,
  })
}

export default post
