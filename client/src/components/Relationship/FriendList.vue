<template>
  <scroll-container :items="friends" class="friend-list">
    <template v-slot="slotProps">
      <friend-item :friend="slotProps.item"></friend-item>
    </template>
  </scroll-container>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import FriendItem from '../Relationship/FriendItem.vue'

export default {
  components: { FriendItem },
  computed: {
    ...mapState(['id']),
    ...mapState('relationship', ['friends']),
  },
  methods: {
    ...mapActions({ getFriends: 'relationship/getFriends' })
  },
  created() {
    this.getFriends(this.id).then((res) => {
      // this.friends = res
    })
  }
}
</script>

<style scoped>
.friend-list {
  height: 90vh;
  overflow: auto;
}
</style>
