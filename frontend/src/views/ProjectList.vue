<template>
  <div class="project-list">
    <div class="header">
      <h1>我的作品集</h1>
      <el-button type="primary" size="large" @click="handleCreateProject">
        新建作品
      </el-button>
    </div>
    <div class="content">
      <el-empty description="暂无作品，点击上方按钮开始创作吧" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const handleCreateProject = () => {
  ElMessage.info('功能开发中...')
}
</script>

<style scoped>
.project-list {
  min-height: 100vh;
  background: linear-gradient(135deg, #f5f5f0 0%, #faf8f5 100%);
  padding: 40px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;
}

.header h1 {
  font-size: 32px;
  color: #2c3e50;
  margin: 0;
  letter-spacing: 2px;
}

.content {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}
</style>
