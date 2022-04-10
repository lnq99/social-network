<template>
  <card :style="cls" class="center hover">
    <router-link :key="$route.fullPath" :to="link">
      <slot></slot>
    </router-link>
  </card>
</template>

<script>
export default {
  props: ['link', 'cls'],
}
</script>