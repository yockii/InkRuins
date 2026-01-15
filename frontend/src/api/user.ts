import request from '@/utils/request'
import type { LoginRequest, RegisterRequest, LoginResponse, User, ApiResponse } from '@/types'

export const userApi = {
  login(data: LoginRequest): Promise<ApiResponse<LoginResponse>> {
    return request.post('/login', data)
  },

  register(data: RegisterRequest): Promise<ApiResponse<null>> {
    return request.post('/register', data)
  },

  getMyInfo(): Promise<ApiResponse<User>> {
    return request.get('/user/my-info')
  },
}
