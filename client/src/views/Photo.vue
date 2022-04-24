<template>
  <short-info :id="id">
    <template v-slot="slotProps">
      <el-avatar class="ava" :size="60" :src="slotProps.avatars"></el-avatar>
      <h2>{{ slotProps.uname }} - Photos</h2>
      <br />
    </template>
  </short-info>
  <grid v-if="loaded" :items="photos">
    <template v-slot="slotProps">
      <el-image
        class="preview"
        style="height: 100%"
        fit="cover"
        :src="slotProps.item.url"
        :preview-src-list="[slotProps.item.url]"
        lazy
      >
      </el-image>
    </template>
  </grid>
</template>

<script>
import { mapActions } from 'vuex'
import { squarePreviewDiv } from '@/utils.js'

export default {
  data() {
    return { loaded: false }
  },
  methods: {
    ...mapActions({ getPhotos: 'photo/getPhotosOfProfile' })
  },
  created() {
    this.id = this.$route.params.id
    this.getPhotos(this.id).then(res => {
      this.photos = res
      this.loaded = true
    })
  },
  mounted() {
    squarePreviewDiv()
  },
}
</script>
