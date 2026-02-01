import axios from 'axios'
import type { AxiosResponse } from 'axios'

interface Tour {
  id: number
  name: string
  created_at: string
  updated_at: string
  base_price: number
}

interface GroupRequestPayload {
  requested_tour_id: number
  requested_date: string
  name: string
  email: string
  pax: number
}

const api = axios.create({
  baseURL: '/api',
})

export const getTours = (): Promise<AxiosResponse<Tour[]>> => {
  return api.get('/tours')
}

export const createGroupRequest = (
  data: GroupRequestPayload
): Promise<AxiosResponse> => {
  return api.post('/groups/create', data)
}
