<template>
  <div>
    <div>
      <h2 class="text-center p-3">Bài thi Ứng dụng Công nghệ thông tin</h2>

      <div class="nav-center">
        <span class="fw-bold fs-6 m-4">Thí sinh: {{ studentName }}</span>
      </div>

      <Question v-for="(question, index) in questions" :key="index" :question="question" :questionNumber="index + 1"
        @answer-selected="updateSelectedAnswer(index, $event)" />

      <div class="container d-flex justify-content-between align-items-center py-2">
        <div class="nav-center nav-time me-4">
          <span class="fw-bold fs-5"><i class="far fa-clock mx-2"></i><span id="timer">{{ timer }}</span></span>
        </div>

        <div class="nav-right d-flex align-items-center">
          <button @click="completeExam()" id="btn-nop-bai" class="btn btn-hero btn-primary" role="button"><i
              class="far fa-file-lines me-1"></i> Nộp bài</button>
        </div>
      </div>

      <span class="score" v-if="showResults">{{ examResults }}</span>
    </div>
  </div>
</template>

<script>
import { reactive, ref } from 'vue';
import Question from '../components/Question.vue';
// import Infomation from './views/components/Infomation.vue';
import { LoadQuestions } from '../../wailsjs/go/main/App';

const examMetadata = ref({
  exam_id: "",
  name_exam: "Bài thi ứng dụng Công nghệ thông tin",
  exam_time: 30,
  exam_start: "2024-03-25T00:00:00",
  exam_end: "2024-03-25T00:30:00",
  easy_quest: 1,
  medium_quest: 2,
  hard_quest: 3,
})

export default {
  components: {
    Question,
    // Infomation
  },
  props: {
    examData: Object,
    required: true
  },
  data() {
    return {
      questions: [], // Mảng các câu hỏi được tạo từ func LoadQuestions
      showResults: false,
      examResults: null,
      selectedAnswers: [],
      studentName: 'Nguyễn Hữu Tài Linh',
      examStarted: false,

      timer: '00:00:00',
      timeRemaining: 60 * 30, // 30 minutes in seconds
      intervalId: null,
    };
  },
  mounted() {
    this.startTimer();

    this.loadQuestions();
  },
  computed: {
    // examMetadata() {
    //   return this.examData.metadata;
    // },
    // examQuestions() {
    //   return this.examData.danh_sach_cau_hoi.map(question => ({
    //     ...question,
    //     selectedAnswer: null // Initialize selectedAnswer for each question
    //   }));
    // }
  },
  methods: {
    async loadQuestions() {
      try {
        const questions = await LoadQuestions();
        this.questions = questions;
        console.log('Câu hỏi:', questions);
      } catch (error) {
        console.error('Lỗi rồi, không thể load câu hỏi từ file:', error);
      }
    },
    // updateStudentName(name) {
    //   this.studentName = name;
    // },
    startExam() {
      this.examStarted = true;
    },

    updateSelectedAnswer(questionIndex, selectedAnswer) {
    this.selectedAnswers[questionIndex] = selectedAnswer;
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
      const correctAnswers = this.questions.filter((question, index) => this.selectedAnswers[index] === question.answer);
      let score = (correctAnswers.length / this.questions.length) * 10;
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
