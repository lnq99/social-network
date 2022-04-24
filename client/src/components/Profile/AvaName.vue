<template>
  <div class="avaname">
    <el-avatar class="ava" :size="200" :src="avatar"></el-avatar>
    <h2 class="name">{{ name }}</h2>
  </div>
</template>

<script>

export default {
  props: ['name', 'avatar'],
}
</script>

<style scoped>
.ava {
  margin: 20px;
}
.name {
  text-align: center;
}
.avaname {
  display: flex;
  flex-direction: column;
  text-align: center;
}
</style>
