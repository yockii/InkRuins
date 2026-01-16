import request from '@/utils/request'
import type { Project, CreateProjectRequest, ListProjectRequest, ApiResponse, PaginatedResponse } from '@/types'

export const projectApi = {
  getMyProjects(params: ListProjectRequest): Promise<PaginatedResponse<Project>> {
    return request.get('/projects/list', { params })
  },

  getProject(id: string): Promise<ApiResponse<Project>> {
    return request.get(`/projects/instance/${id}`)
  },

  createProject(data: CreateProjectRequest): Promise<ApiResponse<Project>> {
    return request.post('/projects/create', data)
  },

  updateProject(id: string, data: Project): Promise<ApiResponse<null>> {
    return request.put(`/projects/update/${id}`, data)
  },

  deleteProject(id: string): Promise<ApiResponse<null>> {
    return request.delete(`/projects/delete/${id}`)
  },
}
