<template>
  <!-- <div id="app">
    <Information />
    <Exam :examData="examData" />
    <ImportFile @exam-data-imported="onExamDataImported" />
  </div> -->

  <router-view></router-view>
</template>

<script>
// import Exam from './views/Exam.vue';
// // import Information from './Information.vue';

// export default {

//   components: {
//     Exam
//   },
//   data() {
//     return {
//       examData: {}
//     };
//   },

//   created() {
//     // Fetch data from API
//     this.fetchExamData();
//   },
//   methods: {
//     async fetchExamData() {
//       try {
//         const response = await fetch('http://localhost:8080/api/de_thi');
//         if (!response.ok) {
//           throw new Error('Failed to fetch data from API');
//         }
//         this.examData = await response.json();
//       } catch (error) {
//         console.error(error);
//       }
//     },

//     onExamDataImported(examData) {
//       this.examData = examData;
//     }
//   }
// };
</script>

<style></style>