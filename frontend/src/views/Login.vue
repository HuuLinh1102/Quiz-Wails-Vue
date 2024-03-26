<template>
  <section class="relative isolate overflow-hidden bg-white px-6 py-24 sm:py-32 lg:px-8">
    <div
      class="absolute inset-0 -z-10 bg-[radial-gradient(45rem_50rem_at_top,theme(colors.indigo.100),white)] opacity-20" />
    <div
      class="absolute inset-y-0 right-1/2 -z-10 mr-16 w-[200%] origin-bottom-left skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 sm:mr-28 lg:mr-0 xl:mr-16 xl:origin-center" />
    <div class="mx-auto max-w-2xl lg:max-w-4xl">
      <img class="mx-auto h-12" src="./../assets/images/itclogo.webp" alt="" />
      <figure class="mt-10">
        <blockquote class="text-center text-xl font-semibold leading-8 text-gray-900 sm:text-2xl sm:leading-9">
          <p>Thông tin thí sinh</p>
        </blockquote>
        <figcaption class="mt-10">
          <img v-if="user.sex == 'Nam'" class="mx-auto h-10 w-10 rounded-full" src="./../assets/images/man.png"
            alt="" />
          <img v-else class="mx-auto h-10 w-10 rounded-full" src="./../assets/images/woman.png" alt="" />
          <info-line title="Số báo danh" :subTitle="user.idNumber || ''"></info-line>
          <info-line title="Họ và tên" :subTitle="user.fullName || ''"></info-line>
          <info-line title="Giới tính" :subTitle="user.sex || ''"></info-line>
          <info-line title="Ngày sinh" :subTitle="user.bod || ''"></info-line>
          <info-line title="Nơi sinh" :subTitle="user.pod || ''"></info-line>
          <info-line title="Phòng thi" :subTitle="user.room || ''"></info-line>
        </figcaption>
      </figure>
    </div>
    <div class="flex justify-center mt-6">
      <button
        class="flex justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        @click="lunch">
        Bắt đầu thi
      </button>
    </div>
  </section>
</template>

<script setup>
import { useRoute, useRouter } from "vue-router";
import { ref } from "vue";
import InfoLine from "../components/InfoLine.vue";

const router = useRouter();
const route = useRoute();
const user = ref({
  idNumber: "",
  fullName: "",
  sex: "Nữ",
  bod: "22/04/2002",
  pod: "Lâm Đồng",
  room: "TV1",
});

const init = () => {
  user.value = JSON.parse(route.params.user);
};

init();

function lunch() {
  router.push({
    name: "exam",
    params: {
      data: "A"
      // exam: JSON.stringify(exam),
    }
  })
}
</script>
