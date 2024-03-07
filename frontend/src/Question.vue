<template>
  <div>
    <p>{{ question }}</p>
    <ul>
      <li v-for="(choice, index) in choices" :key="index" @click="selectAnswer(choice)">
        <input type="radio" :value="choice" v-model="selectedAnswer">{{ choice }}
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
