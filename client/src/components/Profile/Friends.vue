<template>
  <h2>
    Friends
    <a v-if="loaded" class="btn-right">See more ({{ friends.length }})</a>
  </h2>
  <div v-if="loaded" class="card-inner">
    <grid :items="friends.slice(0, 5)">
      <template v-slot="slotProps">
        <el-avatar
          shape="square"
          :size="70"
          :src="slotProps.item.avatars"
        ></el-avatar>
        <p class="friend-name">{{ slotProps.item.name }}</p>
      </template>
    </grid>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  props: ['id'],
  data() {
    return { loaded: false }
  },
  methods: {
    ...mapActions({ getFriends: 'relationship/getFriends' })
  },
  created() {
    this.getFriends(this.id).then((res) => {
      this.friends = res
      this.loaded = true
    })
  }
}
</script>

<style scoped>
.friend-name {
  margin-top: 0;
  margin-bottom: 0;
  font-size: 0.8em;
  /* white-space: nowrap;
  max-width: 80px;
  overflow: hidden;
  text-overflow: ellipsis; */
}
.btn-right {
  float: right;
  font-weight: lighter;
  font-size: 0.8em;
  margin-top: 0.1em;
}
</style>
