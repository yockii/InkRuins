export interface User {
  id: string
  username: string
  email: string
  password?: string
  created_at: number
  updated_at: number
}

export interface LoginRequest {
  username: string
  password: string
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}

export interface LoginResponse {
  token: string
  user: User
}

export interface ApiResponse<T = any> {
  code: number
  message: string
  data: T
}

export interface PaginatedResponse<T = any> {
  code: number
  message: string
  data: T[]
  total: number
  page: number
  page_size: number
}
