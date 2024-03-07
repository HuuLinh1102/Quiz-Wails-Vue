<template>
  <div  v-if="examMetadata">
    <h1 class="text-center">{{ examMetadata.ten_de_thi }}</h1>
    <div class="" v-for="(question, index) in examQuestions" :key="index">
      <Question :question="question.noi_dung_cau_hoi" :choices="question.lua_chon" @answer-selected="updateSelectedAnswer(index, $event)" />
    </div>
    <!-- <div class="d-flex align-items-center flex flex-row justify-between">
      <div class="nav-time me-4">
        <span class="fw-bold fs-5"><i class="far fa-clock mx-2"></i><span id="timer">00:00:00</span></span>
      </div>
    </div>
    <div>
      <button @click="completeExam()" id="btn-nop-bai" class="btn btn-hero btn-primary" role="button"><i class="far fa-file-lines me-1"></i> Nộp bài</button>
    </div> -->
      <div class="container d-flex justify-content-between align-items-center py-2">
          <div class="nav-center nav-time me-4">
            <span class="fw-bold"><i class="far fa-clock mx-2"></i><span id="timer">00:00:00</span></span>

          </div>
          <div class="nav-right d-flex align-items-center">
              <button @click="completeExam()" id="btn-nop-bai" class="btn btn-hero btn-primary" role="button"><i class="far fa-file-lines me-1"></i> Nộp bài</button>
          </div>
      </div>
      <!-- <div class="text-center ">

        <p  v-if="showResults">{{ examResults }}</p>
      </div> -->
        <span class="score " v-if="showResults">{{ examResults }}</span>

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
      selectedAnswers: []
    };
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
<<<<<<< HEAD
      this.examQuestions[index] = reactive({ ...this.examQuestions[index], selectedAnswer: choice });
=======
     this.examQuestions[index] = reactive({
        ...this.examQuestions[index],
        selectedAnswer: choice
      });
>>>>>>> cfe27b5d9538112619093fb97d270a24898b6c55
    },
    completeExam() {
      const score = this.calculateScore();
      this.examResults = `Điểm số của bạn là: ${score}`;
      this.showResults = true;
    },
    calculateScore() {
      const correctAnswers = this.examQuestions.filter(question => question.dap_an_dung === question.selectedAnswer);
      const score = (correctAnswers.length / this.examQuestions.length) * 10;
      return score;
    }
  }
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

