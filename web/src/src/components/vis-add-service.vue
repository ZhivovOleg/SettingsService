<template>
  <div class="p-4 flex gap-4 sticky top-0 bg-white border-b-2">
    <BaseInput v-model="model" @keydown.enter="emit('add-settings')"/>
    <BaseButton text="Добавить" :disabled="!model" @click="emit('add-settings')"/>
  </div>
</template>

<script setup lang="ts">
import BaseButton from "@/components/shared/BaseButton.vue";
import BaseInput from "@/components/shared/BaseInput.vue";


const model = defineModel()
const emit = defineEmits<{
  (event: 'add-settings'): void
}>()
</script>
