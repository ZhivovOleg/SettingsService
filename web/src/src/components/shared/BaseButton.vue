<template>
  <button class="
  hover:enabled:text-white
  hover:enabled:border-transparent
  text-white
  font-semibold
  py-2
  px-4
  border
  rounded
  disabled:bg-opacity-50
  disabled:cursor-not-allowed
  disabled:text-opacity-50
  disabled:shadow-none
" :class="[
    color === 'warning' ? 'bg-red-500' : 'bg-select-color',
    color === 'warning' ? 'border-red-500' : 'border-select-color',
    color === 'warning' ? 'hover:enabled:bg-red-600' : 'hover:enabled:bg-hover-color',
    ]" :disabled="disabled">
    {{ text }}
  </button>
</template>

<script setup lang="ts">
interface PropsType {
  text: string
  disabled?: boolean
  color?: 'warning' | 'default'
}

withDefaults(defineProps<PropsType>(), {
  disabled: false,
  color: 'default'
})
</script>

<style lang="scss" scoped>

</style>
