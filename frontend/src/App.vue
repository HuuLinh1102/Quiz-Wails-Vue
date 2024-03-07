<!-- <template>
  <div id="app">
    <h1>Thi trắc nghiệm demo</h1>
    <h1 v-if="deThi.metadata">{{ deThi.metadata.ten_de_thi }}</h1>
    <div v-if="deThi.danh_sach_cau_hoi">
      <div v-for="(cauHoi, index) in deThi.danh_sach_cau_hoi" :key="index">
        <Question :noiDung="cauHoi.noi_dung_cau_hoi">
          <AnswerChoice v-for="(luaChon, idx) in cauHoi.lua_chon" :key="idx" :noiDung="luaChon" />
        </Question>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      deThi: {} // Lưu trữ dữ liệu đề thi
    };
  },
  async created() {
    // Lấy dữ liệu đề thi từ API
    await this.fetchData();
  },
  methods: {
    async fetchData() {
  try {
    const response = await fetch('http://localhost:8080/api/de_thi');
    if (!response.ok) {
      throw new Error('Lỗi khi lấy dữ liệu từ API');
    }
    
    // Kiểm tra nếu phản hồi không phải là JSON
    const contentType = response.headers.get('content-type');
    if (!contentType || !contentType.includes('application/json')) {
      throw new Error('Dữ liệu từ API không phải là JSON');
    }

    // Đảm bảo rằng body của phản hồi chưa được đọc
    const responseData = await response.json();
    
    // Sử dụng dữ liệu trả về ở đây
    this.deThi = responseData;
  } catch (error) {
    console.error(error);
  }

}
  }
};
</script>

<style>
</style>  -->


<template>
  <div id="app">
    <Exam :examData="examData" />
  </div>
</template>

<script>
import Exam from './Exam.vue';

export default {
  components: {
    Exam
  },
  data() {
    return {
      examData: {}
    };
  },
  created() {
    // Fetch data from API
    this.fetchExamData();
  },
  methods: {
    async fetchExamData() {
      try {
        const response = await fetch('http://localhost:8080/api/de_thi');
        if (!response.ok) {
          throw new Error('Failed to fetch data from API');
        }
        this.examData = await response.json();
      } catch (error) {
        console.error(error);
      }
    }
  }
};
</script>  

<style>

</style>