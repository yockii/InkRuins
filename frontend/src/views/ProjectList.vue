<template>
  <div class="project-list">
    <div class="header">
      <div class="header-left">
        <h1 class="title">我的作品集</h1>
        <p class="subtitle">让灵感在墨色中流淌</p>
      </div>
      <div class="header-right">
        <el-input
          v-model="searchQuery"
          placeholder="搜索作品..."
          prefix-icon="Search"
          clearable
          size="large"
          @input="handleSearch"
        />
        <el-button
          type="primary"
          size="large"
          class="create-button"
          @click="showCreateDialog = true"
        >
          新建作品
        </el-button>
      </div>
    </div>



    <div class="content">
      <div v-if="loading" class="loading-container">
        <el-skeleton :rows="8" animated />
      </div>
      <div v-else-if="projects.length === 0" class="empty-container">
        <el-empty description="暂无作品，点击上方按钮开始创作吧" />
      </div>
      <div v-else class="project-grid">
        <div
          v-for="project in projects"
          :key="project.id"
          class="project-card"
          @click="handleProjectClick(project)"
        >
          <div class="project-cover">
            <div class="project-title">{{ project.title }}</div>
            <div class="project-genre">{{ project.genre }}</div>
          </div>
          <div class="project-info">
            <div class="project-description">{{ project.description || '暂无简介' }}</div>
            <div class="project-stats">
              <div class="stat-item">
                <span class="stat-label">进度</span>
                <el-progress
                  :percentage="calculateProgress(project)"
                  :stroke-width="4"
                  :color="getProgressColor(project)"
                  :show-text="false"
                />
              </div>
              <div class="stat-item">
                <span class="stat-label">状态</span>
                <el-tag :type="getStatusType(project.status)">
                  {{ getStatusText(project.status) }}
                </el-tag>
              </div>
              <div class="stat-item">
                <span class="stat-label">字数</span>
                <span class="stat-value">{{ formatWordCount(project.current_words) }}</span>
              </div>
            </div>
            <div class="project-actions">
              <el-button
                type="text"
                size="small"
                @click.stop="handleEditProject(project)"
              >
                编辑
              </el-button>
              <el-button
          type="text"
          size="small"
          text-color="#c0392b"
          @click.stop="handleDeleteClick(project)"
        >
          删除
        </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div v-if="!loading && projects.length > 0" class="pagination">
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :page-sizes="[8, 16, 24, 32]"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalProjects"
        @size-change="handlePageSizeChange"
        @current-change="handleCurrentPageChange"
      />
    </div>

    <el-dialog
      v-model="showCreateDialog"
      title="新建作品"
      width="500px"
      center
    >
      <el-form
        ref="createFormRef"
        :model="createForm"
        :rules="createRules"
        label-position="top"
        size="large"
      >
        <el-form-item prop="title">
          <el-input
            v-model="createForm.title"
            placeholder="作品标题"
            prefix-icon="EditPen"
          />
        </el-form-item>
        <el-form-item prop="description">
          <el-input
            v-model="createForm.description"
            placeholder="作品简介"
            type="textarea"
            :rows="3"
          />
        </el-form-item>
        <el-form-item prop="genre">
          <el-select v-model="createForm.genre" placeholder="作品类型">
            <el-option label="玄幻" value="玄幻" />
            <el-option label="都市" value="都市" />
            <el-option label="科幻" value="科幻" />
            <el-option label="武侠" value="武侠" />
            <el-option label="历史" value="历史" />
            <el-option label="言情" value="言情" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showCreateDialog = false">取消</el-button>
          <el-button type="primary" @click="handleCreateProject">创建</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog
      v-model="showEditDialog"
      title="编辑作品"
      width="500px"
      center
    >
      <el-form
        ref="editFormRef"
        :model="editForm"
        :rules="editRules"
        label-position="top"
        size="large"
      >
        <el-form-item prop="title">
          <el-input
            v-model="editForm.title"
            placeholder="作品标题"
            prefix-icon="EditPen"
          />
        </el-form-item>
        <el-form-item prop="description">
          <el-input
            v-model="editForm.description"
            placeholder="作品简介"
            type="textarea"
            :rows="3"
          />
        </el-form-item>
        <el-form-item prop="genre">
          <el-select v-model="editForm.genre" placeholder="作品类型">
            <el-option label="玄幻" value="玄幻" />
            <el-option label="都市" value="都市" />
            <el-option label="科幻" value="科幻" />
            <el-option label="武侠" value="武侠" />
            <el-option label="历史" value="历史" />
            <el-option label="言情" value="言情" />
          </el-select>
        </el-form-item>
        <el-form-item prop="status">
          <el-select v-model="editForm.status" placeholder="作品状态">
            <el-option label="策划中" value="planning" />
            <el-option label="写作中" value="writing" />
            <el-option label="已完成" value="completed" />
            <el-option label="暂停" value="paused" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showEditDialog = false">取消</el-button>
          <el-button type="primary" @click="handleUpdateProject">保存</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog
      v-model="showDeleteDialog"
      title="确认删除"
      width="400px"
      center
    >
      <div class="delete-dialog-content">
        <p>您确定要删除作品 <strong>{{ projectToDelete?.title }}</strong> 吗？</p>
        <p class="delete-warning">此操作不可恢复，请谨慎操作。</p>
      </div>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="showDeleteDialog = false">取消</el-button>
          <el-button type="danger" @click="handleDeleteProject">确认删除</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, type FormInstance, type FormRules } from 'element-plus'
import type { Project, CreateProjectRequest } from '@/types'
import { projectApi } from '@/api'

const router = useRouter()

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref(8)
const loading = ref(false)
const projects = ref<Project[]>([])
const totalProjects = ref(0)

const showCreateDialog = ref(false)
const showEditDialog = ref(false)
const showDeleteDialog = ref(false)
const projectToDelete = ref<Project | null>(null)

const createFormRef = ref<FormInstance>()
const editFormRef = ref<FormInstance>()

const createForm = ref<CreateProjectRequest>({
  title: '',
  description: '',
  genre: '',
})

const editForm = ref<Project>({
  id: '',
  user_id: '',
  title: '',
  description: '',
  genre: '',
  world_time_period: '',
  world_location: '',
  world_atmosphere: '',
  world_rules: '',
  target_words: 0,
  current_words: 0,
  status: 'planning',
  narrative_perspective: '',
  created_at: 0,
  updated_at: 0,
})

const createRules: FormRules = {
  title: [
    { required: true, message: '请输入作品标题', trigger: 'blur' },
    { min: 2, max: 100, message: '标题长度在 2 到 100 个字符', trigger: 'blur' },
  ],
  description: [
    { max: 500, message: '简介长度不能超过 500 个字符', trigger: 'blur' },
  ],
  genre: [
    { required: true, message: '请选择作品类型', trigger: 'change' },
  ],
}

const editRules: FormRules = createRules



const calculateProgress = (project: Project) => {
  if (!project.target_words) return 0
  return Math.min(100, Math.round((project.current_words / project.target_words) * 100))
}

const getProgressColor = (project: Project) => {
  const progress = calculateProgress(project)
  if (progress < 33) return '#e67e22'
  if (progress < 66) return '#f39c12'
  return '#27ae60'
}

const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    planning: '策划中',
    writing: '写作中',
    completed: '已完成',
    paused: '暂停',
  }
  return statusMap[status] || status
}

const getStatusType = (status: string) => {
  const typeMap: Record<string, string> = {
    planning: 'info',
    writing: 'success',
    completed: 'primary',
    paused: 'warning',
  }
  return typeMap[status] || 'info'
}

const formatWordCount = (count: number) => {
  if (count < 1000) return `${count}字`
  return `${(count / 1000).toFixed(1)}万字`
}

const fetchProjects = async () => {
  loading.value = true
  try {
    const response = await projectApi.getMyProjects({
      page: currentPage.value,
      size: pageSize.value,
      title: searchQuery.value,
    })
    if (response.code === 0) {
      projects.value = response.data ?? []
      totalProjects.value = response.total ?? 0
    } else {
      ElMessage.error(response.message || '获取项目列表失败')
      projects.value = []
      totalProjects.value = 0
    }
  } catch (error: any) {
    ElMessage.error(error.response?.data?.message || '获取项目列表失败')
    projects.value = []
    totalProjects.value = 0
  } finally {
    loading.value = false
  }
}

const handleSearch = () => {
  currentPage.value = 1
  fetchProjects()
}



const handlePageSizeChange = (size: number) => {
  pageSize.value = size
  currentPage.value = 1
}

const handleCurrentPageChange = (page: number) => {
  currentPage.value = page
}

const handleProjectClick = (project: Project) => {
  router.push(`/workspace/${project.id}`)
}

const handleCreateProject = async () => {
  if (!createFormRef.value) return

  await createFormRef.value.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    try {
      const response = await projectApi.createProject(createForm.value)
      if (response.code === 0) {
        ElMessage.success('项目创建成功')
        showCreateDialog.value = false
        fetchProjects()
        createForm.value = {
          title: '',
          description: '',
          genre: '',
        }
      } else {
        ElMessage.error(response.message || '项目创建失败')
      }
    } catch (error: any) {
      ElMessage.error(error.response?.data?.message || '项目创建失败')
    } finally {
      loading.value = false
    }
  })
}

const handleEditProject = (project: Project) => {
  editForm.value = { ...project }
  showEditDialog.value = true
}

const handleUpdateProject = async () => {
  if (!editFormRef.value) return

  await editFormRef.value.validate(async (valid) => {
    if (!valid) return

    loading.value = true
    try {
      const response = await projectApi.updateProject(editForm.value.id, editForm.value)
      if (response.code === 0) {
        ElMessage.success('项目更新成功')
        showEditDialog.value = false
        fetchProjects()
      } else {
        ElMessage.error(response.message || '项目更新失败')
      }
    } catch (error: any) {
      ElMessage.error(error.response?.data?.message || '项目更新失败')
    } finally {
      loading.value = false
    }
  })
}

const handleDeleteClick = (project: Project) => {
  projectToDelete.value = project
  showDeleteDialog.value = true
}

const handleDeleteProject = async () => {
  if (!projectToDelete.value) return
  
  try {
    const response = await projectApi.deleteProject(projectToDelete.value.id)
    if (response.code === 0) {
      ElMessage.success('项目删除成功')
      fetchProjects()
    } else {
      ElMessage.error(response.message || '项目删除失败')
    }
  } catch (error: any) {
    ElMessage.error(error.response?.data?.message || '项目删除失败')
  } finally {
    showDeleteDialog.value = false
    projectToDelete.value = null
  }
}

onMounted(() => {
  fetchProjects()
})
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
  align-items: flex-start;
  margin-bottom: 40px;
  flex-wrap: wrap;
  gap: 20px;
}

.header-left {
  flex: 1;
}

.title {
  font-size: 36px;
  color: #2c3e50;
  margin: 0 0 8px 0;
  letter-spacing: 2px;
}

.subtitle {
  font-size: 16px;
  color: #2c3e50;
  margin: 0;
  opacity: 0.7;
  letter-spacing: 1px;
}

.header-right {
  display: flex;
  gap: 20px;
  align-items: flex-start;
  flex-wrap: wrap;
}

.header-right :deep(.el-input) {
  width: 300px;
  margin-bottom: 10px;
}



.content {
  min-height: 400px;
}

.loading-container {
  width: 100%;
}

.empty-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 400px;
}

.project-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 30px;
}

.project-card {
  background: rgba(255, 255, 255, 0.85);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 6px 25px rgba(0, 0, 0, 0.12);
  transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
  cursor: pointer;
  backdrop-filter: blur(15px);
  border: 1px solid rgba(245, 245, 240, 0.7);
  position: relative;
}

.project-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(135deg, rgba(212, 175, 55, 0.05) 0%, transparent 100%);
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}

.project-card:hover {
  transform: translateY(-8px) rotateX(5deg);
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.2);
  border-color: #d4af37;
}

.project-card:hover::before {
  opacity: 1;
}

.project-cover {
  height: 200px;
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 20px;
  position: relative;
  overflow: hidden;
  border-bottom: 3px solid #d4af37;
}

.project-cover::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: 
    repeating-linear-gradient(45deg, transparent, transparent 10px, rgba(255, 255, 255, 0.05) 10px, rgba(255, 255, 255, 0.05) 20px),
    radial-gradient(circle at 30% 40%, rgba(212, 175, 55, 0.15) 0%, transparent 50%),
    radial-gradient(circle at 70% 70%, rgba(245, 245, 240, 0.1) 0%, transparent 40%);
  mix-blend-mode: overlay;
  animation: subtleFlow 10s ease-in-out infinite;
}

@keyframes subtleFlow {
  0%, 100% { transform: translateY(0) translateX(0); }
  25% { transform: translateY(-5px) translateX(5px); }
  50% { transform: translateY(5px) translateX(-5px); }
  75% { transform: translateY(-3px) translateX(3px); }
}

.project-title {
  font-size: 26px;
  color: #f5f5f0;
  font-weight: bold;
  text-align: center;
  margin-bottom: 12px;
  position: relative;
  z-index: 1;
  text-shadow: 0 2px 8px rgba(0, 0, 0, 0.5);
  letter-spacing: 1px;
  line-height: 1.3;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.project-genre {
  font-size: 14px;
  color: #d4af37;
  position: relative;
  z-index: 1;
  background: rgba(0, 0, 0, 0.3);
  padding: 4px 12px;
  border-radius: 20px;
  backdrop-filter: blur(5px);
  letter-spacing: 1px;
  text-transform: uppercase;
  font-weight: 500;
}

.project-info {
  padding: 20px;
}

.project-description {
  font-size: 14px;
  color: #2c3e50;
  margin-bottom: 20px;
  line-height: 1.6;
  height: 60px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}

.project-stats {
  margin-bottom: 20px;
}

.stat-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
  font-size: 14px;
}

.stat-item:last-child {
  margin-bottom: 0;
}

.stat-label {
  color: #2c3e50;
  opacity: 0.7;
}

.stat-value {
  color: #2c3e50;
  font-weight: 500;
}

.project-actions {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
  padding-top: 20px;
  border-top: 1px solid rgba(26, 26, 46, 0.1);
}

.pagination {
  margin-top: 40px;
  display: flex;
  justify-content: center;
}

.delete-dialog-content {
  text-align: center;
  padding: 20px 0;
}

.delete-warning {
  color: #c0392b;
  margin-top: 15px;
  font-size: 14px;
  font-weight: 500;
}

.create-button {
  background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%);
  border: none;
  letter-spacing: 2px;
  font-weight: 500;
  position: relative;
  overflow: hidden;
  transition: all 0.3s ease;
}

.create-button::before {
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

.create-button:hover::before {
  width: 300px;
  height: 300px;
}

.create-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
}

:deep(.el-dialog__header) {
  border-bottom: 1px solid rgba(26, 26, 46, 0.1);
}

:deep(.el-dialog__title) {
  color: #2c3e50;
  font-size: 20px;
  font-weight: 500;
}



/* 响应式设计 */
@media (max-width: 1200px) {
  .project-grid {
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 25px;
  }
}

@media (max-width: 768px) {
  .project-list {
    padding: 20px;
  }
  
  .header {
    flex-direction: column;
    align-items: stretch;
    margin-bottom: 30px;
  }
  
  .header-right {
    flex-direction: column;
  }
  
  .header-right :deep(.el-input) {
    width: 100%;
  }
  
  .project-grid {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
    gap: 20px;
  }
  
  .project-cover {
    height: 180px;
  }
  
  .project-title {
    font-size: 22px;
  }
  
  .project-info {
    padding: 15px;
  }
}

@media (max-width: 480px) {
  .project-list {
    padding: 15px;
  }
  
  .title {
    font-size: 28px;
  }
  
  .project-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }
  
  .project-card {
    max-width: 100%;
  }
  
  .pagination {
    padding: 0 10px;
  }
  
  :deep(.el-pagination) {
    font-size: 13px;
  }
}
</style>
