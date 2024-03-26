<template>
  <section
    class="relative isolate overflow-hidden bg-white px-6 py-24 sm:py-32 lg:px-8"
  >
    <div
      class="absolute inset-0 -z-10 bg-[radial-gradient(45rem_50rem_at_top,theme(colors.indigo.100),white)] opacity-20"
    />
    <div
      class="absolute inset-y-0 right-1/2 -z-10 mr-16 w-[200%] origin-bottom-left skew-x-[-30deg] bg-white shadow-xl shadow-indigo-600/10 ring-1 ring-indigo-50 sm:mr-28 lg:mr-0 xl:mr-16 xl:origin-center"
    />
    <div class="mx-auto max-w-2xl lg:max-w-4xl">
      <!--<img class="mx-auto h-12" src="https://doitsolutions.vn/wp-content/uploads/2023/05/logo.png" alt="" />-->
      <img class="mx-auto h-32" src="./../assets/images/itclogo.webp" alt="" />
      <figure class="mt-10">
        <blockquote
          class="text-center text-xl font-semibold leading-8 text-gray-900 sm:text-3xl sm:leading-9"
        >
          <p>TRUNG TÂM CÔNG NGHỆ THÔNG TIN</p>
          <p>TRƯỜNG ĐẠI HỌC ĐÀ LẠT</p>
        </blockquote>
        <figcaption class="mt-10">
          <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
            <!--<form class="space-y-6">-->
              <div>
                <label
                  for="email"
                  class="block text-sm font-medium leading-6 text-gray-900"
                  >Số báo danh</label
                >
                <div class="mt-2 mb-4">
                  <input
                    v-model="data.idnum"
                    id="idnum"
                    name="idnum"
                    type="text"
                    required=""
                    class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                  />
                </div>
              </div>
              <div>
                <button
                  class="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
                  @click="login"
                >
                  Tiếp theo
                </button>
              </div>
            <!--</form>-->
          </div>
          <div
            class="mt-4 flex items-center justify-center space-x-3 text-base"
          >
            <div class="font-semibold text-gray-900">
              Phần mềm thi trắc nghiệm
            </div>
            <svg
              viewBox="0 0 2 2"
              width="3"
              height="3"
              aria-hidden="true"
              class="fill-gray-900"
            >
              <circle cx="1" cy="1" r="1" />
            </svg>
            <div class="text-gray-600">2024 Copyright © ITC DLU</div>
          </div>
        </figcaption>
      </figure>
    </div>
  </section>
  <TransitionRoot appear :show="isOpen" as="template">
    <Dialog as="div" @close="closeDialog" class="relative z-10">
      <TransitionChild
        as="template"
        enter="duration-300 ease-out"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="duration-200 ease-in"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black/25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div
          class="flex min-h-full items-center justify-center p-4 text-center"
        >
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <DialogPanel
              class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
            >
              <DialogTitle
                as="h3"
                class="text-lg font-medium leading-6 text-gray-900"
              >
                {{result.Message}}
              </DialogTitle>
              <div class="mt-2">
                <p class="text-sm text-gray-500">
                Vui lòng kiểm tra lại thông tin số báo danh được thông báo, hoặc liên hệ Giám thị để biết thêm thông tin chi tiết.
                </p>
              </div>

              <div class="mt-4">
                <button
                  type="button"
                  class="inline-flex justify-center rounded-md border border-transparent bg-blue-100 px-4 py-2 text-sm font-medium text-blue-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                  @click="closeDialog"
                >
                  Thử lại
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<script setup>
import { reactive, ref } from "vue";
import { Login } from "../../wailsjs/go/main/App";
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  DialogDescription,
  TransitionRoot,
  TransitionChild,
} from "@headlessui/vue";

import {useRouter} from 'vue-router'

const router = useRouter()

const isOpen = ref(false);

const data = ref({
  idnum: "",
});
const result = ref({
  Message: "",
  Status: false
})

function closeDialog() {
  isOpen.value = false;
}

function login() {
 
  Login(data.value.idnum).then((res) => {
    if (res.Status){
      let user = {
        "idNumber":"T",
        "fullName": "Trần Thị Tường Vân",
        "sex": "Nữ",
        "bod":"22/04/2002", 
        "pod":"Lâm Đồng",
        "room": "TV1"
      }
      router.push({
        name: 'login',
        params:{
          user: JSON.stringify(user)
        }
        
      })
      
    }else{
      result.value = res
      isOpen.value = true
    }
    
  });
}
</script>
