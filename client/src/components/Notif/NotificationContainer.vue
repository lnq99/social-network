<template>
  <scroll-container :items="notif">
    <template v-slot="slotProps">
      <notif-item :notif="slotProps.item"></notif-item>
    </template>
  </scroll-container>
</template>

<script>
import { mapGetters } from 'vuex'
import NotifItem from './NotifItem.vue'

export default {
  components: { NotifItem },
  computed: {
    ...mapGetters({ notif: 'notif/notif' }),
  },
  created() {
    this.$store.dispatch('notif/getNotif')
    console.log(this.notif)
  }
}
</script>
