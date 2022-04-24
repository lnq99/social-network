<template>
  <div v-if="loaded" class="profile">
    <card class="col1">
      <el-affix :offset="-0">
        <card class="col-card card-hl center">
          <center>
            <avaname :name="profile.name" :avatar="profile.avatarl"></avaname>
          </center>
        </card>
        <card class="col-card card-hl">
          <photos :id="profile.id"></photos>
        </card>
        <card class="col-card card-hl">
          <friends :id="profile.id"></friends>
        </card>
      </el-affix>
    </card>
    <card class="col2">
      <card class="col-card card-hl">
        <intro :intro="profile.intro" :editable="isOwnProfile"></intro>
      </card>
      <card :key="profile.id" v-if="!isOwnProfile" class="col-card card-hl row">
        <relation :id="profile.id"></relation>
      </card>
      <post-container :posts="posts"></post-container>
    </card>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Avaname from '@/components/Profile/AvaName.vue'
import Photos from '@/components/Profile/Photos.vue'
import Friends from '@/components/Profile/Friends.vue'
import Intro from '@/components/Profile/Intro.vue'
import PostContainer from '@/components/Post/PostContainer.vue'
import Center from '../components/Base/Center.vue'
import Relation from '@/components/Relationship/Relation.vue'

export default {
  name: 'profile',
  components: { PostContainer, Avaname, Photos, Intro, Friends, Center, Relation },
  data() {
    return { loaded: false, isOwnProfile: false }
  },
  computed: {
    ...mapState(['id']),
  },
  methods: {
    ...mapActions({ getProfile: 'profile/getProfile', getPosts: 'post/getPostsOfProfile' }),
    fetchData(id) {
      this.loaded = false
      this.getProfile(id).then(res => {
        this.profile = res
        this.getPosts(id).then(res => {
          this.posts = res
          this.loaded = true
          this.isOwnProfile = this.id == this.profile.id
        })
      })
    }
  },
  async beforeRouteUpdate(to, from) {
    this.fetchData(to.params.id)
  },
  created() {
    this.fetchData(this.$route.params.id)
  },
}
</script>

<style scoped>
.profile {
  width: 100%;
  display: flex;
}
.col1 {
  flex: 1;
}
.col2 {
  flex: 2;
  padding-left: 12px;
}
.col-card {
  margin-bottom: 12px;
  padding: 18px;
}
</style>
