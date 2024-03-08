<template>
  <div  v-if="examMetadata">
    <h1 class="text-center">{{ examMetadata.ten_de_thi }}</h1>
    <div class="" v-for="(question, index) in examQuestions" :key="index">
      <Question :question="question.noi_dung_cau_hoi" :choices="question.lua_chon" @answer-selected="updateSelectedAnswer(index, $event)" />
    
    </div>
      <div class="container d-flex justify-content-between align-items-center py-2">
          <div class="nav-center nav-time me-4">
            <span class="fw-bold fs-5"><i class="far fa-clock mx-2"></i><span id="timer">{{ timer }}</span></span>
          </div>

          <div class="nav-right d-flex align-items-center">
              <button @click="completeExam()" id="btn-nop-bai" class="btn btn-hero btn-primary" role="button"><i class="far fa-file-lines me-1"></i> Nộp bài</button>
          </div>
      </div>
        <span class="score "v-if="showResults">{{ examResults }}</span>
      </div>
</template>

<script>
import { reactive, ref } from 'vue';
import Question from './Question.vue';

export default {
  components: {
    Question
  },
  props: {
    examData: Object,
    required: true
  },
  data() {
    return {
      showResults: false,
      examResults: null,
      selectedAnswers: [],

      timer: '00:00:00',
      timeRemaining: 60 * 30, // 30 minutes in seconds
      intervalId: null,
    };
  },
  mounted() {
    this.startTimer();
  },
  computed: {
    examMetadata() {
      return this.examData.metadata;
    },
    examQuestions() {
      return this.examData.danh_sach_cau_hoi.map(question => ({
        ...question,
        selectedAnswer: null // Initialize selectedAnswer for each question
      }));
    }
  },
  methods: {
    updateSelectedAnswer(index, choice) {
     this.examQuestions[index] = reactive({
        ...this.examQuestions[index],
        selectedAnswer: choice
      });
    },
    completeExam() {
    this.$swal.fire({
      title: "<p class='fs-3 mb-0'>Bạn có chắc chắn muốn nộp bài ?</p>",
      html: "<p class='text-muted fs-6 text-start mb-0'>Khi xác nhận nộp bài, bạn sẽ không thể sửa lại bài thi của mình. Chúc bạn may mắn!</p>",
      icon: "info",
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: "Vâng, chắc chắn!",
      cancelButtonText: "Huỷ",
    }).then((result) => {
      if (result.isConfirmed) {
        const score = this.calculateScore();
        this.examResults = `Điểm số của bạn là: ${score}`;
        clearInterval(this.intervalId);
        this.showResults = true;
      }
    });
    },

    calculateScore() {
      const correctAnswers = this.examQuestions.filter(question => question.dap_an_dung === question.selectedAnswer);
      const score = (correctAnswers.length / this.examQuestions.length) * 10;
      return score;
    },

    // Define the startTimer method
    startTimer() {
      this.intervalId = setInterval(() => {
        if (this.timeRemaining === 0) {
          clearInterval(this.intervalId);
          this.completeExam();
          // Time's up. Handle it here.
        } else {
          this.timeRemaining--;
          this.timer = this.formatTime(this.timeRemaining);
        }
      }, 1000);
    },

    formatTime(timeInSeconds) {
      const minutes = Math.floor(timeInSeconds / 60);
      const seconds = timeInSeconds % 60;
      return `${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
    },
  },
};
</script>

<style>
.score {
  font-size: 24px;
  font-weight: bold;
  color: rgb(233, 17, 17);
  display: flex;
  justify-content: center;
}
</style>

