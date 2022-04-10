<template>
  <div
    v-infinite-scroll="loadMore"
    infinite-scroll-disabled="disabled"
    infinite-scroll-immediate="false"
  >
    <template v-for="(item, index) in items">
      <slot :item="item"></slot>
    </template>
    <p v-if="loading">Loading...</p>
    <p v-if="noMore">No more</p>
  </div>
</template>

<script>
export default {
  // props: ['items'],
  props: {
    items: {
      type: Array
    },
    load: {
      default: () => { },
      type: Function
    }
  },
  data() {
    return {
      count: this.items.length,
      loading: false,
    }
  },
  computed: {
    noMore() {
      return this.count >= this.items.length
    },
    disabled() {
      return this.loading || this.noMore
    },
  },
  methods: {
    loadMore() {
      // setTimeout(() => {
      //   this.count += 2
      //   this.loading = false
      // }, 1000)
      this.load()
      this.loading = true
      this.count += 5
      this.loading = false
    },
  },
}
</script>
