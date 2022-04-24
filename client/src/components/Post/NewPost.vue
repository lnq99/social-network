<template>
  <card class="new-post card-hl">
    <div class="new-post-content">
      <el-avatar class="ava" :size="40" :src="avatars"></el-avatar>
      <div class="input">
        <span @click="onCreatePost = true"> Create a post </span>
      </div>
    </div>
  </card>
  <el-dialog
    title="Create Post"
    v-model="onCreatePost"
    width="max(40%, 400px)"
    center
  >
    <post-editor @posted="onCreatePost = false"></post-editor>
  </el-dialog>
</template>

<script>
import { mapState } from 'vuex'
import PostEditor from "./PostEditor.vue"
export default {
  components: { PostEditor },
  data() {
    return {
      onCreatePost: false,
    }
  },
  computed: {
    ...mapState('profile', ['avatars'])
  },
};
</script>

<style lang="scss" scoped>
@import "@/style.scss";

.new-post {
  width: 100%;
  margin-bottom: 10px;
}
.new-post-content {
  padding: $p4;
  display: flex;
  align-content: center;
}
.input {
  margin-left: $p4;
  background: var(--hl);
  border: none;
  display: flex;
  flex-direction: column;
  justify-content: center;
  flex: 1;
}
.input:hover {
  background: var(--hl2);
}
</style>
