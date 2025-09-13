<script setup>
import { ref } from 'vue';
import ApprovalTable from './components/ApprovalTable.vue'
import Modal from './components/Modal.vue'

const showModal = ref(false)
const isApprove = ref(true)

const openModal = (approve) => {
  showModal.value = true;
  isApprove.value = approve
}
</script>

<template>
  <div class="row" style="gap: 10px; margin-bottom: 5px;">
    <button class="btn btn-green" @click="openModal(true)">อนุมัติ</button>
    <button class="btn btn-red" @click="openModal(false)">ไม่อนุมัติ</button>
  </div>
  <div>
  </div>
  <ApprovalTable />
  <Modal v-if="showModal" :isApprove="isApprove" @close="showModal = false" />
</template>
