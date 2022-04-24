<template v-if=loaded>
  <el-avatar class="ava" :size="200" :src="avatar"></el-avatar>
  <h2 class="name">{{ name }}</h2>

  <link-card :cls="cls" link="/"> News </link-card>
  <link-card :cls="cls" :link="{ name: 'Profile', params: { id } }">
    Profile
  </link-card>
  <link-card :cls="cls" :link="{ name: 'Photo', params: { id } }">
    Photo
  </link-card>
  <link-card :cls="cls" link="/notif">
    <!-- <el-badge :value="100" :max="10" class="item"> -->
    <p>Notification</p>
    <!-- </el-badge> -->
  </link-card>
  <link-card :cls="cls" link="/search"> Search </link-card>
  <link-card :cls="cls" link="/logout"> Logout </link-card>
  <div class="center">
    <el-switch
      v-model="theme"
      class="switch"
      active-color="#111"
      inactive-color="#aaa"
      active-text="Dark"
      inactive-text="Light"
    >
    </el-switch>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  data() {
    return {
      loaded: false,
      cls: {
        'font-size': '1.2em',
        'font-weight': 300,
        'letter-spacing': '1px',
        'word-spacing': '1px',
        'padding': '18px',
      }
    }
  },
  computed: {
    theme: {
      get() {
        return this.$store.state.theme.isDark
      },
      set(value) {
        this.$store.dispatch('theme/switchTheme', value)
      },
    },
    ...mapState({ id: 'id' }),
    ...mapState('profile', { name: 'name', avatar: 'avatarl' }),
  },
  created() {
    this.loaded = true
  }
}
</script>

<style scoped>
.ava {
  margin-top: 20px;
}
.name {
  margin-bottom: 4vh;
}
.switch {
  margin-top: 4vh;
  margin-left: -10px;
}
p {
  margin: 0 10px 0 10px;
}
</style>
