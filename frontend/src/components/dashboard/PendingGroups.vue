<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import { api, getPendingGroups } from '@/api'
import axios from 'axios';
import { useToast } from 'primevue/usetoast';

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


const toast = useToast()

const groups = ref<Group[]>([])
const groupsLoading = ref(false)
const groupsError = ref('')


onMounted(async () => {
  groupsLoading.value = true
  groupsError.value = ''
  try {
    const response = await getPendingGroups()
    groups.value = response;
  } catch (err) {
    groupsError.value = 'Could not load tours. Make sure backend is running! Refresh to retry.'
  } finally {
    groupsLoading.value = false
  }
})

const acceptGroup = async (groupId: string, email: string) => {
  try {
    await api.put(`/api/groups/${groupId}/accept`);

    // toast.add({
    //   severity: 'success',
    //   summary: `${email} accepted`,
    //   detail: 'Payment info sent',
    //   life: 3000,
    // });

    groups.value = groups.value.filter(group => group.id !== groupId);

  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to accept group',
      life: 3000,
    });
  }
};

const declineGroup = async (groupId: string, email: string) => {
  try {
    await axios.put(`/api/groups/${groupId}/decline`);

    // toast.add({
    //   severity: 'warn',
    //   summary: `${email} declined`,
    //   detail: 'a rejection email has been sent',
    //   life: 3000,
    // });
    //
    groups.value = groups.value.filter(group => group.id !== groupId);

  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to reject group',
      life: 3000,
    });
  }
};

</script>

<template>
  <div class="card">
    <div>
      <h2 class="flex justify-content-around flex-wrap">Pending tour requests</h2>
    </div>

    <div v-if="groupsLoading" class="flex justify-center items-center h-32">
      <p>Loading groups...</p>
    </div>

    <Message v-else-if="groupsError" severity="error" :closable="false">
      {{ groupsError }}
    </Message>

    <DataTable v-else :value="groups" tableStyle="min-width: 50rem" :paginator="groups.length > 10" :rows="10"
      responsiveLayout="scroll" stripedRows>
      <Column field="requested_date" header="Requested date" sortable>
        <template #body="slotProps">
          {{ new Date(slotProps.data.requested_date).toLocaleDateString('en-GB', {
            year: 'numeric', month: 'short', day: 'numeric'
          }) }}
        </template>
      </Column>
      <Column field="name" header="Name" sortable />
      <Column field="email" header="Email">
        <template #body="{ data }">
          <a :href="`mailto:${data.email}`" class="text-primary hover:underline">
            {{ data.email }}
          </a>
        </template>
      </Column>
      <Column field="pax" header="Pax" />
      <Column field="requested_tour_id" header="Tour ID" sortable />
      <Column field="booking_id" header="Booking ID" sortable />
      <!-- <Column field="id" header="Group ID" /> -->
      <Column header="Actions">
        <template #body="slotProps">
          <div class="flex gap-2">
            <Button label="Accept" icon="pi pi-check" class="p-button-sm p-button-success"
              @click="acceptGroup(slotProps.data.id, slotProps.data.email)" />
            <Button label="Decline" icon="pi pi-times" class="p-button-sm p-button-danger"
              @click="declineGroup(slotProps.data.id, slotProps.data.email)" />
          </div>
        </template>
      </Column>

    </DataTable>
  </div>
</template>
