<template>
  <div class="newsfeed">
    <new-post></new-post>
    <post-container :posts="posts" :getMorePosts="getFeed"></post-container>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import NewPost from '@/components/Post/NewPost.vue'
import PostContainer from '@/components/Post/PostContainer.vue'

export default {
  components: {
    NewPost,
    PostContainer,
  },
  computed: {
    ...mapGetters({ posts: 'post/feed' })
  },
  methods: {
    ...mapActions({ getFeed: 'post/getFeed' })
  },
  created() {
    this.getFeed()
  }
}
</script>

<style scoped>
.newsfeed {
  display: flex;
  flex-direction: column;
  align-items: center;
  background: var(--bg0);
  margin-left: 6vw;
  margin-right: 6vw;
}
</style>
