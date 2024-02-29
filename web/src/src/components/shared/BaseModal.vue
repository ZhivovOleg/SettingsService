<template>
  <teleport to="body">
    <div class="overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 max-h-full bg-black bg-opacity-50 flex h-full">
      <div class="relative p-4 w-full max-h-full flex justify-center items-center">
        <div class="relative bg-white rounded-lg shadow">
          <button @click="closeModal"  type="button" class="absolute top-3 end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center ">
            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
            </svg>
            <span class="sr-only" >Close modal</span>
          </button>
          <div class="p-4 md:p-5 min-w-80 min-h-40 max-h-[90vh] overflow-auto">
            <slot></slot>
          </div>
        </div>
      </div>
    </div>

  </teleport>
</template>

<script setup lang="ts">
import {onMounted, onUnmounted} from "vue";

const emit = defineEmits<{
  (event: 'close'): void
}>()
function closeModal() {
  emit('close')
}

function triggerEspace (e: KeyboardEvent) {
  if (e.key === 'Escape') {
    closeModal()
  }
}

onMounted(() => {
  document.addEventListener("keydown", triggerEspace)
})

onUnmounted(() => {
  document.removeEventListener("keydown", triggerEspace)
})
</script>

<style scoped>

</style>
