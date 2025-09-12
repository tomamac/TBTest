<script setup>
import { ref } from 'vue';

const rows = ref([])

const fetchData = () => {
  fetch("http://localhost:8001/api/approval")
    .then(res => res.json())
    .then((result) => {
      rows.value = result
      console.log(rows.value)
    })
}
fetchData()
</script>

<template>
  <div>
    <button></button>
    <button></button>
  </div>
  <div>
    <table>
      <thead>
        <tr>
          <th>
            <input type="checkbox" id="vehicle1" name="vehicle1" value="Bike"></input>
          </th>
          <th>รายการ</th>
          <th>เหตุผล</th>
          <th>สถานะเอกสาร</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item of rows">
          <td>
            <input type="checkbox" id="vehicle1" name="vehicle1" value="Bike"></input>
          </td>
          <td>{{ item.title }}</td>
          <td>{{ item.description }}</td>
          <td>{{ item.status }}</td>
        </tr>
      </tbody>
    </table>
  </div>
  <div class="modal-overlay">
    <div class="modal-content">
      <div class="modal-header">
        <h2>ยืนยันการอนุมัติ</h2>
      </div>
      <form>
        <div class="modal-body row">
          <label for="description">เหตุผล : </label>
          <textarea id="description" type="text" name="description"></textarea>
        </div>
        <div class="modal-footer row">
          <button class="btn submit-btn" type="submit">
            อนุมัติ
          </button>
          <button class="btn cancel-btn" type="button">
            ยกเลิก
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
button {
  height: 100px;
  width: 100px;
}

table {
  height: 200px;
  width: 200px;
  border: 1px solid red;
}

th {
  height: 50px;
  width: 50px;
  border: 1px solid green;
}

td {
  height: 50px;
  width: 50px;
  border: 1px solid green;
}

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
  background-color: rgba(0, 0, 0, 0.5);
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
  background-color: rgb(71, 71, 199);
  color: white;
  font-size: 12px;
  padding: 10px;
  margin: 0px;
}

h2 {
  margin: 0;
}

.modal-body {
  gap: 10px;
  padding: 15px;
}

.modal-body label{
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

.btn {
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
  font-size: 14px;
  font-weight: bold;
  color: white;
  height: 40px;
}

.submit-btn {
  background-color: rgb(71, 71, 199);
}

.cancel-btn {
  background-color: rgb(199, 71, 71);
}

.row {
  display: flex;
  flex-wrap: wrap;
}

* {
  box-sizing: border-box;
}


@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
