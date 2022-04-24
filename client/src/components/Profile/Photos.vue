<template>
  <h2>
    Photos
    <router-link :to="`/photo/${id}`">
      <span class="btn-right"> See more ({{ photos.length }}) </span>
    </router-link>
  </h2>
  <grid v-if="loaded" :items="photos.slice(0, 8)">
    <template v-slot="slotProps">
      <img class="preview" :src="slotProps.item.url" />
    </template>
  </grid>
</template>

<script>
import { mapActions } from 'vuex'
import { squarePreview } from '@/utils.js'

export default {
  props: ['id'],
  data() {
    return { photos: [], loaded: false }
  },
  methods: {
    ...mapActions({ getPhotos: 'photo/getPhotosOfProfile' }),
  },
  created() {
    this.getPhotos(this.id).then(res => {
      this.photos = res
      this.loaded = true
    })
  },
  mounted() {
    squarePreview()
  },
}
</script>

<style scoped>
.flex-container {
  display: flex;
  flex-flow: row wrap;
}
.btn-right {
  float: right;
  font-weight: lighter;
  font-size: 0.8em;
  margin-top: 0.1em;
}
</style>
