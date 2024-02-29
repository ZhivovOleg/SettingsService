<script setup lang="ts">
import VisHeader from "@/components/vis-header.vue";
import {Notifications} from "@kyvg/vue3-notification";
import VisService from "@/pages/vis-service.vue";
</script>

<template>
  <div class="app">
    <vis-header/>
    <vis-service/>
    <Notifications
        position="bottom left"
        dangerously-set-inner-html/>
  </div>
</template>
