<template>
  <div class="shadow-md rounded-5 bg-white question">
    <p>{{ question.content }}</p>
    <ul>
      <li v-for="(choice, index) in question.choices" :key="index" @click="selectAnswer(choice)">
        <input v-model="selectedAnswer" :value="choice.content" type="radio" :id="`choice-${index}`"/>
        <label :for="`choice-${index}`">{{ choice.content }}</label>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    question: Object // Thay đổi kiểu dữ liệu của question thành Object để phù hợp với dữ liệu truyền từ component cha
  },
  data() {
    return {
      selectedAnswer: null
    };
  },
  methods: {
    selectAnswer(choice) {
      this.selectedAnswer = choice.content;
      this.$emit('answer-selected', choice.content);
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
