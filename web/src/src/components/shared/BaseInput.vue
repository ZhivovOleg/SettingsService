<template>
  <input v-model="model"
         id="first_name"
         class="bg-gray-50
         border
         border-gray-300
         text-gray-900 t
         w-text-sm
         rounded-l
         rounded-r
         focus:ring-select-color
         focus:border-select-color
         block
         w-full
         p-2.5
         focus:outline-none"
         placeholder="Имя сервиса"
         required/>

</template>
<script setup lang="ts">
const model = defineModel()

</script>

<style scoped lang="scss">

</style>
