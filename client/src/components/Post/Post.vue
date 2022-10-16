<template>
  <card v-if="loaded" class="post card-hl">
    <div class="post-header">
      <div class="row">
        <short-info :id="data.userId">
          <template v-slot="slotProps">
            <el-avatar class="ava" :size="40" :src="slotProps.avatars"></el-avatar>
            <div class="post-header-r">
              <span class="post-author">{{ slotProps.uname }}</span>
              <time class="post-time">{{ data.created }}</time>
            </div>
          </template>
        </short-info>
      </div>
      <el-dropdown v-if="userId == data.userId" size="mini">
        <i class="el-icon-arrow-down el-icon--right"></i>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item disabled>Edit</el-dropdown-item>
            <el-dropdown-item @click="deletePost(id)" icon="el-icon-delete" divided>Delete</el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
    <p class="post-content">{{ data.content }}</p>
    <div v-if="(data.atchType = 'photo')" class="attach">
      <img :src="data.atchUrl" />
    </div>
    <hr />
    <react-cmt :id="id" :initReaction="data.reaction" :cmtCount="data.cmtCount"></react-cmt>
  </card>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import ReactCmt from './ReactCmt.vue'

export default {
  components: { ReactCmt },
  props: ['id'],
  data() {
    return {
      data: {},
      author: {},
      loaded: false,
    }
  },
  computed: {
    ...mapState({ userId: 'id' })
  },
  methods: {
    ...mapActions({ getPhoto: 'photo/getPhoto', getProfileShort: 'profile/getProfileShort' }),
    ...mapActions({ getPost: 'post/getPost', deletePost: 'post/delete' }),
  },
  created() {
    this.getPost(this.id).then(data => {
      if (!data) return
      this.data = data
      // console.log(this.data.reaction)

      this.loaded = true

      if (this.data.atchType === 'photo') {
        this.getPhoto(this.data.atchId).then(data => {
          this.data.atchUrl = data.url
        })
      }
    })
  }
}
</script>

<style lang="scss" scoped>
@import "@/style.scss";

.post {
  margin-bottom: 12px;
  padding-bottom: $p4;
}

.post-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  text-align: left;
  padding: $p4;
}

.post-header-r {
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding-left: $p4;
}

.post-author {
  font-weight: 900;
}

.post-time {
  font-size: 0.8em;
  font-weight: 100;
}

.post-content {
  margin: 0 $p4 $p4 $p4;
  text-align: left;
}

img {
  width: 100%;
  display: block;
  margin-bottom: 8px;
}
</style>
