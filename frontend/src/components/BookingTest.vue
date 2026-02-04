<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Form } from '@primevue/forms'
import type { FormSubmitEvent } from '@primevue/forms'
import { zodResolver } from '@primevue/forms/resolvers/zod'
import * as z from 'zod'
import { getAllTours, createGroupRequest } from '../api.ts'


interface Tour {
  id: number
  name: string
  base_price: string | number
}

interface FormValues {
  requested_tour_id: number
  requested_date: string
  name: string
  email: string
  pax: number
}

const schema = z.object({
  requested_tour_id: z.coerce.number().min(1, { message: 'Please select a tour' }),
  requested_date: z.string()
    .min(1, { message: 'Please select a preferred date' })
    .transform((val) => `${val}T00:00:00Z`),
  name: z.string().min(1, { message: 'Your name is required' }).trim(),
  email: z.string().email({ message: 'Please enter a valid email address' }),
  pax: z.coerce.number().min(1, { message: 'At least 1 person is required' })
})

const tours = ref<Tour[]>([])
const toursLoading = ref(true)
const toursError = ref('')
const loading = ref(false)
const messageText = ref('')
const messageSeverity = ref<'success' | 'error' | 'info' | 'warn'>('success')

onMounted(async () => {
  toursLoading.value = true
  toursError.value = ''
  try {
    const response = await getAllTours()
    tours.value = response;
  } catch (err) {
    toursError.value = 'Could not load tours. Make sure backend is running! Refresh to retry.'
  } finally {
    toursLoading.value = false
  }
})

const onSubmit = async (event: FormSubmitEvent) => {
  if (!event.valid) return

  const values = event.values as FormValues

  loading.value = true
  messageText.value = ''
  messageSeverity.value = 'success'

  try {
    await createGroupRequest(values);
    messageText.value = "Request received!\nWe'll review it ASAP\nand email you soon."
    messageSeverity.value = 'success'

  } catch (err: any) {
    messageText.value = err.response?.data?.error || 'Something went wrong – please try again.'
    messageSeverity.value = 'error'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="form-container">
    <h3>Request Your Tour Spot!</h3>

    <div v-if="toursLoading">Loading available tours...</div>
    <Message v-else-if="toursError" severity="error">
      {{ toursError }}
    </Message>
    <Message v-else-if="tours.length === 0" severity="warn">
      No tours available yet. Check back soon!
    </Message>

    <Form v-else v-slot="$form" :resolver="zodResolver(schema)" @submit="onSubmit">

      <div class="form-group">
        <label>Which tour?</label>
        <Dropdown name="requested_tour_id" :options="tours" optionValue="id"
          :optionLabel="(tour) => `${tour.name} – JPY ${tour.base_price}`" placeholder="Select a tour" fluid />
        <Message v-if="$form.requested_tour_id?.invalid" severity="error" size="small">
          {{ $form.requested_tour_id?.error?.message }}
        </Message>
      </div>

      <div class="form-group">
        <label>Preferred date</label>
        <InputText name="requested_date" type="date" fluid />
        <Message v-if="$form.requested_date?.invalid" severity="error" size="small">
          {{ $form.requested_date?.error?.message }}
        </Message>
      </div>

      <div class="form-group">
        <label>Your name</label>
        <InputText name="name" placeholder="Name McNameson" fluid />
        <Message v-if="$form.name?.invalid" severity="error" size="small">
          {{ $form.name?.error?.message }}
        </Message>
      </div>

      <div class="form-group">
        <label>Email address</label>
        <InputText name="email" type="email" placeholder="you@example.com" fluid />
        <Message v-if="$form.email?.invalid" severity="error" size="small">
          {{ $form.email?.error?.message }}
        </Message>
      </div>

      <div class="form-group">
        <label>How many people?</label>
        <InputNumber name="pax" :min="1" showButtons fluid />
        <Message v-if="$form.pax?.invalid" severity="error" size="small">
          {{ $form.pax?.error?.message }}
        </Message>
      </div>

      <Message v-if="messageText" :severity="messageSeverity" :closable="false">
        {{ messageText }}
      </Message>

      <Button type="submit" :label="loading ? 'Sending request...' : 'Submit Request'"
        :disabled="loading || toursLoading" :loading="loading" />

    </Form>
  </div>
</template>
