<template>
  <div v-if="examMetadata">
    <h1>{{ examMetadata.ten_de_thi }}</h1>
    <div v-for="(question, index) in examQuestions" :key="index">
      <Question :question="question.noi_dung_cau_hoi" :choices="question.lua_chon" @answer-selected="updateSelectedAnswer(index, $event)" />
    </div>
    <button @click="completeExam">Hoàn thành</button>
    <p v-if="showResults">{{ examResults }}</p>
  </div>
</template>

<script>
import { ref } from 'vue';
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
      this.$set(this.examQuestions[index], 'selectedAnswer', choice);
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
