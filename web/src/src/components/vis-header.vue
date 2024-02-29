<script setup lang="ts">

</script>

<template>
  <div class="header pl-4 pr-4">
    <div class="content">
      <img alt="герб" src="@/assets/logo.svg"/>
    </div>
  </div>
</template>

<style scoped lang="postcss">
.header {
  background: var(--header);
}
.content {
  @apply p-2 flex flex-row items-center justify-between
}
</style>
