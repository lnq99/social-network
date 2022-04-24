<template>
  <scroll-container :items="items">
    <template v-slot="slotProps">
      <search-item
        :key="slotProps.item.id"
        :user="slotProps.item"
      ></search-item>
    </template>
  </scroll-container>
</template>

<script>
import SearchItem from './SearchItem.vue'

export default {
  components: { SearchItem },
  props: ['items'],
}
</script>

<style scoped>
</style>
