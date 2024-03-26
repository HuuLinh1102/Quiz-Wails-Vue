<template>
  <div>
    <input type="file" @change="onFileChange">
  </div>
</template>

<script>
import * as XLSX from 'xlsx';

export default {
  methods: {
    onFileChange(e) {
    const file = e.target.files[0];
    const reader = new FileReader();
    reader.onload = (e) => {
        const data = new Uint8Array(e.target.result);
        const workbook = XLSX.read(data, {type: 'array'});
        const worksheet = workbook.Sheets[workbook.SheetNames[0]];
        const jsonData = XLSX.utils.sheet_to_json(worksheet);
        // Convert jsonData to your exam format
        const examData = {
        metadata: {
            ten_de_thi: jsonData[0].ten_de_thi,
            thoi_gian_lam_bai: jsonData[0].thoi_gian_lam_bai,
            so_luong_cau_hoi: jsonData.length
        },
        danh_sach_cau_hoi: jsonData.map(row => ({
            noi_dung_cau_hoi: row.noi_dung_cau_hoi,
            lua_chon: [row.lua_chon_1, row.lua_chon_2, row.lua_chon_3, row.lua_chon_4],
            dap_an_dung: row.dap_an_dung
        }))
        };
        // Emit event with examData
        this.$emit('exam-data-imported', examData);
    };
    reader.readAsArrayBuffer(file);
    }
  }
};
</script>