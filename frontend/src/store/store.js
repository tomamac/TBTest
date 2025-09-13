import axios from "axios";
import { reactive } from "vue";

const BASE_URL = "http://localhost:8001/api";
const MOCK_URL = "http://localhost:8001/api/mock";

export const store = reactive({
  approvals: [],
  selectedIds: [],
  async fetchApprovals() {
    try {
      const res = await axios.get(`${MOCK_URL}/approval`);
      this.approvals = res.data;
    } catch (error) {
      console.log("error", error);
    }
  },
  async updateApprovals(selectedIds, description, isApprove) {
    try {
      if (selectedIds.length === 0) return;
      const status = isApprove ? "approved" : "rejected";
      const body = {
        ids: selectedIds,
        description,
        status,
      };

      const res = await axios.put(`${MOCK_URL}/approval`, body);
      this.approvals = this.approvals.map((approval) =>
        selectedIds.includes(approval.id)
          ? { ...approval, description, status }
          : approval
      );
      this.selectedIds = [];

      console.log(res);
    } catch (error) {
      console.log("error", error);
    }
  },
});
