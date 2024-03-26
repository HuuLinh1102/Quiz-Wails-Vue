<template>
  <div class="mt-4 flex items-center justify-center space-x-3 text-base">
    <div class="text-gray-600 text-left">{{title}}</div>
    <svg
      viewBox="0 0 2 2"
      width="3"
      height="3"
      aria-hidden="true"
      class="fill-gray-900"
    >
      <circle cx="1" cy="1" r="1" />
    </svg>
    <div class="font-semibold text-gray-900 text-right" >{{ subTitle }}</div>
  </div>
</template>

<script setup>
const props = defineProps({
    title: {
        type: String,
        default: ''
    },
    subTitle: {
        type: String,
        default: ''
    }
})
</script>