<script setup>
import { store } from '@/store/store';
import { ref } from 'vue';

defineProps({
	isApprove: Boolean
})

const description = ref("")

const updateApprovals = async (selectedIds, description, isApprove) => {
	await store.updateApprovals(selectedIds, description, isApprove)
}

</script>

<template>
	<div class="modal-overlay">
		<div class="modal-content">
			<div class="modal-header">
				<h2 style="margin: 0;">ยืนยันการ{{ isApprove ? "" : "ไม่" }}อนุมัติ</h2>
			</div>
			<div class="modal-body row">
				<label for="description">เหตุผล : </label>
				<textarea type="text" v-model="description"></textarea>
			</div>
			<div class="modal-footer row">
				<button class="btn btn-blue" type="submit"
					@click="updateApprovals(store.selectedIds, description, isApprove), $emit('close')">
					{{ isApprove ? "" : "ไม่" }}อนุมัติ
				</button>
				<button class="btn btn-red" type="button" @click="$emit('close')">
					ยกเลิก
				</button>
			</div>
		</div>
	</div>
</template>

<style scoped>
.modal-overlay {
	z-index: 1000;
	position: absolute;
	display: flex;
	justify-content: center;
	align-items: center;
	top: 0;
	left: 0;
	right: 0;
	bottom: 0;
	width: 100%;
	height: 100%;
	background-color: rgba(75, 75, 75, 0.5);
}

.modal-content {
	position: relative;
	background-color: #fff;
	min-width: 50%;
	max-width: 95%;
	min-height: 30%;
	max-height: 95%;
}

.modal-header {
	background-color: rgb(0, 112, 192);
	color: white;
	font-size: 12px;
	padding: 10px;
	margin: 0px;
}

.modal-body {
	gap: 10px;
	padding: 15px;
}

.modal-body label {
	margin-top: 5px;
}

.modal-body textarea {
	flex: 1;
	width: 100%;
	height: 150px;
	resize: none;
	border: 1px solid #ccc;
	padding: 5px;
	font-size: 16px;
}

.modal-footer {
	justify-content: flex-end;
	gap: 10px;
	padding: 10px 15px;
}
</style>