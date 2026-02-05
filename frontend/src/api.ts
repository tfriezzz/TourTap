import axios from 'axios';
import authStore from './store';

// axios.defaults.baseURL = 'http://localhost:8080';
export const api = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 10000,
});

api.interceptors.request.use(
  (config) => {
    const token = authStore.state?.accessToken;
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

export interface Group {
  id: string;
  email: string;
  name: string;
  pax: number;
  customer_status: 'pending' | 'confirmed' | 'cancelled' | string;
  requested_tour_id: number;
  requested_date: string;
  booking_id: number;
}

export interface Tour {
  id: number;
  name: string;
  base_price: string | number;
}

export interface Booking {
  booking_id: number;
  tour_name: string;
  date: string;
  group_count: number;
  total_pax: number;
  attending_groups: string;
}

export const getAllTours = async (): Promise<Tour[]> => {
  try {
    const response = await axios.get('/api/tours');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const getPendingGroups = async (): Promise<Group[]> => {
  try {
    const response = await axios.get('/api/groups/get-pending');
    return response.data;
  } catch (error) {
    throw error;
  }
};


export const getAllBookingsByDate = async (date: string): Promise<Booking[]> => {
  try {
    const response = await axios.get('/api/bookings/all-date', {
      params: { date }
    });
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const createGroupRequest = async (payload: {
  email: string;
  name: string;
  pax: number;
  requested_tour_id: number;
  requested_date: string;
}): Promise<void> => {
  try {
    await axios.post('/api/groups/create', payload, {
      headers: {
        'Content-Type': 'application/json',
      },
    });
  } catch (error: any) {
    throw error;
  }
};
