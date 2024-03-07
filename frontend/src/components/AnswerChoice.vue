<template>
  <div class="answer-choice">
    <input type="radio" :name="'choice-' + cauHoiId" :id="'choice-' + id" :value="id" />
    <label :for="'choice-' + id">{{ noiDung }}</label>
  </div>
</template>

<script>
export default {
  props: {
    noiDung: {
      type: String,
      required: true
    },
    id: {
      type: Number,
      required: true
    },
    cauHoiId: {
      type: Number,
      required: true
    }
  }
}
</script>
