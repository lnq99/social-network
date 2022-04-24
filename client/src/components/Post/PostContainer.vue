<template>
  <scroll-container :items="posts" :load="getMorePosts">
    <template v-slot="slotProps">
      <post :key="slotProps.item" :id="slotProps.item"></post>
    </template>
  </scroll-container>
</template>

<script>
import Post from './Post.vue'

export default {
  components: { Post },
  props: ['posts', 'getMorePosts'],
}
</script>
