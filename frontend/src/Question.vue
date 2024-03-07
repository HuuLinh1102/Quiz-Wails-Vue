<template>
  <div class="shadow-md rounded-5 bg-white question ">
    <p class="">{{ question }}</p>
    <ul>
      <li v-for="(choice, index) in choices" :key="index" @click="selectAnswer(choice)">
        <!-- <input type="radio" :value="choice" v-model="selectedAnswer">{{ choice }} -->

        <input v-model="selectedAnswer" :value="choice" type="radio"/> {{ choice }}
        
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    question: String,
    choices: Array
  },
  data() {
    return {
      selectedAnswer: null
    };
  },
  methods: {
    selectAnswer(choice) {
      this.selectedAnswer = choice;
      this.$emit('answer-selected', choice);
    }
  }
};
</script>

<style>
.question {
  margin: 10px;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 4px 6px 0 rgba(0, 0, 0, 0.1);
 }
</style>