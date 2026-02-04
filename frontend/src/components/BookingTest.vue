<script setup lang="ts">
import { Form } from '@primevue/forms'
import type { FormSubmitEvent } from '@primevue/forms';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import axios from 'axios';
import * as z from 'zod';

const schema = z.object({
  username: z.string().min(1, { message: 'Username is required' }),
  email: z.string().email({ message: 'Invalid email address' }),
});

type FormValues = z.infer<typeof schema>;

const onSubmit = async (event: FormSubmitEvent) => {
  if (!event.valid) {
    console.log('Form invalid, not sending');
    return;
  }

  const values = event.values as FormValues;

  console.log('Submitting values:', values);

  try {
    const response = await axios.post('/api/test', values);
    console.log('Success:', response.data);
    // Optionally: reset form or show success message
  } catch (error) {
    console.error('API error:', error);
  }
};
</script>

<template>
  <Form v-slot="$form" :resolver="zodResolver(schema, undefined, { raw: true })" @submit="onSubmit">
    <!-- Username field -->
    <!-- <div> -->
    <!--   <InputText name="username" placeholder="Username" /> -->
    <!--   <Message v-if="$form.username?.invalid" severity="error" size="small"> -->
    <!--     {{ $form.username?.error?.message }} -->
    <!--   </Message> -->
    <!-- </div> -->

    <!-- Email field -->
    <div>
      <InputText name="email" placeholder="Email" />
      <Message v-if="$form.email?.invalid" severity="error" size="small">
        {{ $form.email?.error?.message }}
      </Message>
    </div>

    <!-- Submit button -->
    <Button type="submit" label="Submit" />
  </Form>
</template>
