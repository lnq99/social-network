<template>
  <link-card
    class="hover p18"
    :link="{ name: 'Profile', params: { id: friend.id } }"
  >
    <div class="row">
      <span>
        <el-avatar class="ava" :size="40" :src="friend.avatars"></el-avatar>
      </span>
      <div class="name">{{ friend.name }}</div>
    </div>
  </link-card>
</template>

<script>
export default {
  props: ['friend'],
}
</script>

<style scoped>
.row {
  display: flex;
}
.ava {
  margin: -10px 20px -10px -10px;
}
.name {
  white-space: nowrap;
  overflow: hidden;
}
</style>
