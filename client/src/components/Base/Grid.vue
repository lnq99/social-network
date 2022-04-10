<template>
  <div class="container">
    <div class="grid-item center" v-for="(item, index) in items">
      <slot :item="item"></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: ['items'],
}
</script>

<style scoped>
.container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 0 2%;
}
.grid-item {
  margin-bottom: 6%;
}
</style>
