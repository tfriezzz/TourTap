// api.ts
import axios from 'axios';

// Optional: Configure axios with base URL if needed (e.g., for proxy setup in vite.config.ts)
// Here we assume the frontend proxies /api/ to the backend, or it's same-origin
// const api = axios.create({ baseURL: '/api' }); // Uncomment if using proxy

// Define the shape of a tour (adjust based on your actual backend response)
export interface Tour {
  id: number;
  name: string;
  base_price: string | number; // often string with currency, but number if raw
}

// GET /api/tours - Fetch all available tours
export const getAllTours = async (): Promise<Tour[]> => {
  try {
    const response = await axios.get('/api/tours');
    // Assuming backend returns the array directly (response.data = Tour[])
    // If it wraps in { data: Tour[] }, change to: return response.data.data;
    return response.data;
  } catch (error) {
    // Let the caller handle the error (as in your onMounted try/catch)
    throw error;
  }
};

// POST /api/groups/create - Submit a group request
// Payload matches the backend struct exactly (numbers as int32, date as RFC3339 string)
export const createGroupRequest = async (payload: {
  email: string;
  name: string;
  pax: number;
  requested_tour_id: number;
  requested_date: string; // e.g., "2026-02-04T00:00:00Z"
}): Promise<void> => {
  try {
    await axios.post('/api/groups/create', payload, {
      headers: {
        'Content-Type': 'application/json',
      },
    });
    // Success: no return value needed (matches your original await createGroupRequest(...))
  } catch (error: any) {
    // Throw to let the caller catch and handle (e.g., err.response?.data?.error)
    throw error;
  }
};
