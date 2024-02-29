<template>
  <div class="flex justify-between p-2">
      <BaseButton class="bg-error-color"
                  color="warning"
                  text="Удалить"
                  @click="emit('deleteSettings')"/>
      <BaseButton text="Отменить"
                  :disabled="!isEdit"
                  @click="emit('onCancel')"/>
      <BaseButton :disabled="!isEdit"
                  text="Сохранить"
                  @click="emit('updateSettings')"/>
  </div>
</template>

<script setup lang="ts">
import BaseButton from "@/components/shared/BaseButton.vue";
interface PropsType {
  isEdit: boolean
}


defineProps<PropsType>()

const emit = defineEmits<{
  (event: 'onCancel'): void
  (event: 'updateSettings'): void
  (event: 'deleteSettings'): void
}>()

</script>

<style lang="scss" scoped>

</style>
