import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { User } from '@/types'
import { userApi } from '@/api'

export const useUserStore = defineStore('user', () => {
  const user = ref<User | null>(null)
  const token = ref<string | null>(localStorage.getItem('token'))

  const setToken = (newToken: string) => {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  const setUser = (userData: User) => {
    user.value = userData
  }

  const logout = () => {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
  }

  const fetchUserInfo = async () => {
    try {
      const response = await userApi.getMyInfo()
      if (response.code === 200) {
        user.value = response.data
      }
    } catch (error) {
      console.error('获取用户信息失败', error)
    }
  }

  return {
    user,
    token,
    setToken,
    setUser,
    logout,
    fetchUserInfo,
  }
})
