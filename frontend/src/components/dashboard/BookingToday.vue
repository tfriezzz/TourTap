<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import { getAllBookingsByDate } from '@/api'

export interface Booking {
  booking_id: number
  tour_name: string
  date: string
  group_count: number
  total_pax: number
  attending_groups: string
}

const bookings = ref<Booking[]>([])
const bookingsLoading = ref(false)
const bookingsError = ref('')
const selectedDate = ref<Date | null>(new Date())

const formatDateForApi = (date: Date | null): string | null => {
  if (!date) return null
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}T00:00:00Z`
}

const fetchBookings = async () => {
  const dateStr = formatDateForApi(selectedDate.value)
  if (!dateStr) return

  bookingsLoading.value = true
  bookingsError.value = ''
  bookings.value = []

  try {
    const response = await getAllBookingsByDate(dateStr)
    bookings.value = response
  } catch (err) {
    bookingsError.value = 'Could not load bookings for this date. Make sure backend is running!'
  } finally {
    bookingsLoading.value = false
  }
}

watch(selectedDate, () => {
  fetchBookings()
})

onMounted(() => {
  fetchBookings()
})
</script>

<template>
  <div class="card">
    <h2 class="flex justify-content-around flex-wrap">Bookings</h2>
    <div class="flex justify-content-between flex-wrap">
      <!-- <p>Bookings for:</p> -->
      <Button label="Refresh" icon="pi pi-refresh" @click="fetchBookings" :loading="bookingsLoading" />
      <DatePicker v-model="selectedDate" dateFormat="dd-mm-yy" :showButtonBar="true" :showIcon="true"
        iconDisplay="input" class="w-48" />
    </div>

    <div v-if="bookingsLoading" class="flex justify-center items-center h-32">
      <p>Loading bookings...</p>
    </div>

    <Message v-else-if="bookingsError" severity="error" :closable="false">
      {{ bookingsError }}
    </Message>

    <DataTable v-else :value="bookings" tableStyle="min-width: 50rem" :paginator="bookings.length > 10" :rows="10"
      responsiveLayout="scroll" stripedRows>
      <Column field="booking_id" header="Booking ID" sortable />
      <Column field="tour_name" header="Tour Name" sortable />
      <Column field="date" header="Date" sortable>
        <template #body="slotProps">
          {{ new Date(slotProps.data.date).toLocaleDateString('en-US', {
            year: 'numeric', month: 'short', day: 'numeric'
          }) }}
        </template>
      </Column>
      <Column field="group_count" header="Group Count" sortable />
      <Column field="total_pax" header="Total Pax" sortable />
      <Column field="attending_groups" header="Attending Groups" sortable>
        <template #body="slotProps">
          {{ slotProps.data.attending_groups || 'None' }}
        </template>
      </Column>
    </DataTable>

    <Message v-if="!bookingsLoading && !bookingsError && bookings.length === 0" severity="info" :closable="false">
      No bookings found for {{ selectedDate?.toLocaleDateString() }}.
    </Message>
  </div>
</template>
