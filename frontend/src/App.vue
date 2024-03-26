<<<<<<< HEAD
// OKELA
=======
>>>>>>> da1697b (fix)
<template>
  <div id="app">
    <Exam :examData="examData" />
  </div>
</template>

<script>
import ExamVue from './Exam.Vue';

export default {
  components: {
    ExamVue
  },
  data() {
    return {
      examData: {} // Khởi tạo examData là một đối tượng trống
    };
  },
  created() {
    // Gọi phương thức loadQuestions từ app.go để tải dữ liệu câu hỏi
    this.loadQuestions();
  },
  methods: {
    async loadQuestions() {
      try {
        // Gọi phương thức LoadQuestions từ app.go thông qua Wails
        const response = await this.$wails.Call("LoadQuestions");
        // Gán dữ liệu câu hỏi nhận được từ app.go vào examData
        this.examData = { danh_sach_cau_hoi: response };
      } catch (error) {
        console.error('Error loading questions:', error);
      }
    }
  }
};
</script>

<style>

</style>
