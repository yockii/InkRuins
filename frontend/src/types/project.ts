export interface Project {
  id: string
  user_id: string
  title: string
  description: string
  genre: string
  world_time_period: string
  world_location: string
  world_atmosphere: string
  world_rules: string
  target_words: number
  current_words: number
  status: string
  narrative_perspective: string
  writing_style_id?: string
  created_at: number
  updated_at: number
}

export interface CreateProjectRequest {
  title: string
  description: string
  genre: string
}

export interface ListProjectRequest {
  page: number
  size: number
  title?: string
}
