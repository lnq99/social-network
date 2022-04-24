<template>
  <div class="commentinput">
    <el-avatar class="ava" :size="30" :src="avatars"></el-avatar>
    <el-input
      class="input"
      type="textarea"
      :autosize="{ minRows: 1, maxRows: 4 }"
      placeholder="Write a comment..."
      v-model="cmtModel"
    >
    </el-input>
    <div>
      <el-button
        class="input-color"
        @click="
          $emit('cmt', cmtModel);
          cmtModel = '';
        "
        icon="el-icon-s-promotion"
        size="small"
      ></el-button>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import { ref } from 'vue'

export default {
  data() {
    return {
      cmtModel: ref(''),
    }
  },
  computed: {
    ...mapState('profile', ['avatars'])
  },
}
</script>
<style scoped>
.commentinput {
  display: flex;
  align-content: center;
  justify-content: center;
  margin: 18px;
  margin-bottom: 0;
}

.ava {
  float: left;
}

.input {
  margin-left: 10px;
  width: 90%;
  height: initial !important;
  /* background: var(--hl); */
}
.input .el-textarea__inner {
  background-color: #888;
}
</style>
