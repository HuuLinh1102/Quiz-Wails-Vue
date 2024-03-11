<<<<<<< HEAD

=======
// OKELA
>>>>>>> d269b33faecef6f5448f8f88ab2a2654c659ae97
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