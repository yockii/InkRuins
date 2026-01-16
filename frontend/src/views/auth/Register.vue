<template>
  <div class="register-container">
    <InkBackground />
    <div class="register-content">
      <div class="register-header">
        <h1 class="title">墨墟</h1>
        <p class="subtitle">InkRuins</p>
        <p class="welcome-text">开始你的创作之旅</p>
      </div>

      <div class="register-form">
        <el-form
          ref="formRef"
          :model="registerForm"
          :rules="rules"
          label-position="top"
          size="large"
        >
          <el-form-item prop="username">
            <el-input
              v-model="registerForm.username"
              placeholder="用户名"
              prefix-icon="User"
              clearable
            />
          </el-form-item>

          <el-form-item prop="email">
            <el-input
              v-model="registerForm.email"
              placeholder="邮箱"
              prefix-icon="Message"
              clearable
            />
          </el-form-item>

          <el-form-item prop="password">
            <el-input
              v-model="registerForm.password"
              type="password"
              placeholder="密码"
              prefix-icon="Lock"
              show-password
              clearable
            />
          </el-form-item>

          <el-form-item prop="confirmPassword">
            <el-input
              v-model="registerForm.confirmPassword"
              type="password"
              placeholder="确认密码"
              prefix-icon="Lock"
              show-password
              clearable
            />
          </el-form-item>

          <el-form-item>
            <el-button
              type="primary"
              class="register-button"
              :loading="loading"
              @click="handleRegister"
            >
              注册
            </el-button>
          </el-form-item>
        </el-form>

        <div class="register-footer">
          <span class="footer-text">已有账号？</span>
          <router-link to="/login" class="login-link">
            立即登录
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import InkBackground from '@/components/common/InkBackground.vue'
import { userApi } from '@/api'

const router = useRouter()

const formRef = ref<FormInstance>()
const loading = ref(false)

const registerForm = reactive({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
})

const validateConfirmPassword = (_rule: any, value: any, callback: any) => {
  if (value === '') {
    callback(new Error('请再次输入密码'))
  } else if (value !== registerForm.password) {
    callback(new Error('两次输入密码不一致'))
  } else {
    callback()
  }
}

const rules: FormRules = {
  username: [
    { required: true, message: '请输入用户名', trigger: 'blur' },
    { min: 3, max: 20, message: '用户名长度在 3 到 20 个字符', trigger: 'blur' },
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入正确的邮箱格式', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度至少为 6 个字符', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, validator: validateConfirmPassword, trigger: 'blur' },
  ],
}

const handleRegister = async () => {
  if (!formRef.value) return

  await formRef.value.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    try {
      const response = await userApi.register({
        username: registerForm.username,
        email: registerForm.email,
        password: registerForm.password,
      })
      if (response.code === 200) {
        ElMessage.success('注册成功，请登录')
        router.push('/login')
      } else {
        ElMessage.error(response.message || '注册失败')
      }
    } catch (error: any) {
      ElMessage.error(error.response?.data?.message || '注册失败，请稍后重试')
    } finally {
      loading.value = false
    }
  })
}
</script>

<style scoped>
.register-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  overflow-y: auto;
}

.register-content {
  position: relative;
  z-index: 1;
  width: 100%;
  max-width: 420px;
  padding: 40px;
  animation: fadeIn 0.8s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.register-header {
  text-align: center;
  margin-bottom: 40px;
}

.title {
  font-size: 48px;
  font-weight: bold;
  color: #f5f5f0;
  margin: 0 0 10px 0;
  letter-spacing: 8px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
}

.subtitle {
  font-size: 18px;
  color: #d4af37;
  margin: 0 0 20px 0;
  letter-spacing: 4px;
  font-weight: 300;
}

.welcome-text {
  font-size: 16px;
  color: #f5f5f0;
  margin: 0;
  font-weight: 300;
  letter-spacing: 2px;
}

.register-form {
  background: rgba(245, 245, 240, 0.95);
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3);
  backdrop-filter: blur(10px);
}

:deep(.el-form-item) {
  margin-bottom: 24px;
}

:deep(.el-input__wrapper) {
  background-color: rgba(245, 245, 240, 0.8);
  border: none !important;
  border-bottom: 2px solid #2c3e50 !important;
  border-radius: 0;
  padding: 12px 0;
  box-shadow: none !important;
  transition: all 0.3s ease;
}

:deep(.el-input__wrapper:hover) {
  border-bottom-color: #d4af37 !important;
  background-color: rgba(245, 245, 240, 1);
}

:deep(.el-input__wrapper.is-focus) {
  border-bottom-color: #d4af37 !important;
  background-color: rgba(245, 245, 240, 1);
}

:deep(.el-form-item.is-error .el-input__wrapper) {
  border-bottom-color: #c0392b !important;
  background-color: rgba(245, 245, 240, 1);
  box-shadow: none !important;
}

:deep(.el-form-item__error) {
  color: #c0392b;
  font-size: 12px;
  margin-top: 4px;
  padding-left: 4px;
  position: relative;
}

:deep(.el-form-item__error::before) {
  content: '•';
  position: absolute;
  left: -8px;
  color: #c0392b;
  font-size: 14px;
}

:deep(.el-input__inner) {
  font-size: 16px;
  color: #2c3e50;
}

:deep(.el-input__prefix) {
  color: #2c3e50;
}

.register-button {
  width: 100%;
  height: 48px;
  font-size: 18px;
  font-weight: 500;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  border: none;
  border-radius: 6px;
  color: #f5f5f0;
  letter-spacing: 2px;
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.register-button::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 0;
  height: 0;
  border-radius: 50%;
  background: rgba(212, 175, 55, 0.3);
  transform: translate(-50%, -50%);
  transition: width 0.6s, height 0.6s;
}

.register-button:hover::before {
  width: 300px;
  height: 300px;
}

.register-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

.register-button:active {
  transform: translateY(0);
}

.register-footer {
  margin-top: 24px;
  text-align: center;
}

.footer-text {
  color: #2c3e50;
  font-size: 14px;
  margin-right: 8px;
}

.login-link {
  color: #c0392b;
  text-decoration: none;
  font-weight: 500;
  transition: color 0.3s ease;
}

.login-link:hover {
  color: #d4af37;
  text-decoration: underline;
}
</style>
