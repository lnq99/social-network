<template>
  <div style="margin: 10px; width: 100%">
    <el-input placeholder="Search..." v-model="searchQuery">
      <template #append>
        <el-button
          class="input-color"
          @click="$emit('search', searchQuery); searchQuery = '';"
          icon="el-icon-search"
        ></el-button>
      </template>
    </el-input>
    <div style="position: absolute; z-index: 1">
      <div class="content dropdown-content" v-for="user in searchedUsers" :key="user.name">
        <friend-item :friend="user"></friend-item>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import FriendItem from '../Relationship/FriendItem.vue'

export default {
  components: { FriendItem },
  data() {
    return {
      searchQuery: ''
    }
  },
  computed: {
    ...mapState('relationship', ['friends']),
    searchedUsers() {
      return this.friends.filter((user) => {
        if (this.searchQuery === '') return false
        return (
          user.name.toLowerCase().indexOf(this.searchQuery.toLowerCase()) != -1
        )
      })
    }
  },
}
</script>

<style lang="scss" scoped>
.content {
  // border-top: 1px solid #eee;
  background: var(--hl);
}
.dropdown-content {
  // display: none;
  min-width: 160px;
  box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
  z-index: 1;
}
</style>
