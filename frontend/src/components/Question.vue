<template>
  <div class="shadow-md rounded-lg bg-white p-4 border border-gray-200 question">
    <p>{{ questionNumber }}. {{ question.content }}</p>

    <ul>
      <li v-for="(choice, index) in question.choices" :key="index" @click="selectAnswer(choice, index)">
        <input v-model="selectedAnswer" :value="choice.content" class="mr-5" type="radio" :id="`choice - ${index}`" />
        <label class="mt-2" :for="`choice - ${index}`">{{ answerLabels[index] }}. {{ choice.content }}</label>
      </li>
    </ul>
  </div>

</template>

<script>
export default {
  props: {
    question: Object, // Thay đổi kiểu dữ liệu của question thành Object để phù hợp với dữ liệu truyền từ component cha
    questionNumber: Number
  },
  data() {
    return {
      selectedAnswer: null,
      answerLabels: ['A', 'B', 'C', 'D']
    };
  },
  methods: {
    selectAnswer(choice, index) {
      this.selectedAnswer = choice.content;
      this.$emit('answer-selected', choice.content);
    },
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

input[type="radio"] {
  margin-right: 10px;
}
</style>
