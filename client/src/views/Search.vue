<template>
  <div class="search">
    <h2>Search</h2>
    <el-row>
      <search-box @search="search"></search-box>
    </el-row>
    <hr />
    <br />
    <h3 v-if="searchQuery">Result for "{{ searchQuery }}"</h3>
    <search-container :items="items"></search-container>
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import SearchContainer from '../components/Search/SearchContainer.vue'
import SearchBox from '../components/Search/SearchBox.vue'

export default {
  components: { SearchContainer, SearchBox },
  data() {
    return {
      items: [],
      searchQuery: '',
      item: [
        { id: 1, mutual: 8, type: 'friend' },
        { id: 8, mutual: 2, type: 'request' },
        { id: 3, mutual: 1, type: 'follow' },
        { id: 4, mutual: 3, type: 'block' },
        { id: 2, mutual: 0, type: '' },
      ]
    }
  },
  methods: {
    ...mapActions({ searchProfile: 'profile/searchProfile' }),
    search(searchQuery) {
      if (searchQuery.length < 2) return
      this.searchProfile(searchQuery).then(res => {
        this.searchQuery = searchQuery
        this.items = res
      })
    }
  }
}
</script>

<style scoped>
.search {
  margin-left: 8vw;
  margin-right: 8vw;
}
</style>