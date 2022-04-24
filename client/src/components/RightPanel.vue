<template>
  <el-row>
    <search-box></search-box>
  </el-row>
  <hr />
  <friend-list></friend-list>
</template>

<script>
import FriendList from './Relationship/FriendList.vue'
import SearchBox from './Search/SearchBox.vue'

export default {
  components: { FriendList, SearchBox },
}
</script>
