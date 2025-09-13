<script setup>
import { store } from '@/store/store';
import { computed, onMounted, ref } from 'vue';

onMounted(async () => {
	await store.fetchApprovals();
})

const selectableIds = computed(() =>
	store.approvals.filter(r => r.status === "pending")
		.map(r => r.id)
)

const isAllChecked = ref(false)

const checkAll = () => {
	if (isAllChecked.value) store.selectedIds = selectableIds.value
	else store.selectedIds = []
}
const checkSingle = () => {
	if (store.selectedIds.length === selectableIds.value.length) isAllChecked.value = true
	else isAllChecked.value = false
}
const translateStatus = (status) => {
	switch (status) {
		case "pending": return "รออนุมัติ"
		case "approved": return "อนุมัติ"
		case "rejected": return "ไม่อนุมัติ"
		default: return "Invalid"
	}
}
</script>

<template>
	<table>
		<thead>
			<tr>
				<th style="width: 10%;">
					<input type="checkbox" @change="checkAll" v-model="isAllChecked">
				</th>
				<th style="width: 30%;">รายการ</th>
				<th style="width: 30%;">เหตุผล</th>
				<th style="width: 30%;">สถานะเอกสาร</th>
			</tr>
		</thead>
		<tbody>
			<tr v-for="item of store.approvals">
				<td>
					<input type="checkbox" :disabled='item.status !== "pending"' :value="item.id" @change="checkSingle"
						v-model="store.selectedIds">
				</td>
				<td>{{ item.title }}</td>
				<td>{{ item.description }}</td>
				<td>{{ translateStatus(item.status) }}</td>
			</tr>
		</tbody>
	</table>
</template>

<style scoped>
table {
	width: 100%;
	text-align: center;
	border-collapse: collapse;
}

tr {
	height: 40px;
}

th {
	width: 50px;
	background-color: rgb(0, 112, 192);
	border: 1px solid black;
}

td {
	width: 50px;
	border: 1px solid black;
}

input[type="checkbox"] {
	width: 20px;
	height: 20px;
}
</style>