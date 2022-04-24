<template>
  <slot v-if="loaded" :uname="data.name" :avatars="data.avatars"></slot>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  props: ['id'],
  data() {
    return {
      loaded: false
    }
  },
  methods: {
    ...mapActions({ getProfileShort: 'profile/getProfileShort' }),
  },
  created() {
    this.getProfileShort(this.id).then(data => {
      this.data = data
      this.loaded = true
    })
  }
}
</script>
