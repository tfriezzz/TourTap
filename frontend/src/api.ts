import axios from 'axios';

axios.defaults.baseURL = 'http://localhost:8080';

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
