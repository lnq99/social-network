<template>
  <el-row class="container">
    <el-col :xs="0" :sm="6" :md="5" :lg="4" :xl="4" class="left">
      <left-panel></left-panel
    ></el-col>

    <el-col :xs="24" :sm="18" :md="14" :lg="16" :xl="16" class="main-col">
      <el-col
        :xs="{ span: 24, offset: 0 }"
        :sm="{ span: 24, offset: 0 }"
        :md="{ span: 24, offset: 0 }"
        :lg="{ span: 22, offset: 1 }"
        :xl="{ span: 18, offset: 3 }"
        class="main"
      >
        <router-view></router-view> </el-col
    ></el-col>

    <el-col :xs="0" :sm="0" :md="5" :lg="4" :xl="4" class="right">
      <right-panel></right-panel
    ></el-col>
  </el-row>
</template>

<script>
// @ is an alias to /src
import LeftPanel from '@/components/LeftPanel.vue'
import RightPanel from '@/components/RightPanel.vue'

export default {
  name: 'Main',
  components: {
    LeftPanel,
    RightPanel,
  },
}
</script>

<style scoped>
.left,
.right {
  height: 100vh;
  box-shadow: var(--sep-shadow);
  padding: min(10px, 0.5vw);
  margin: 0;
}

.container {
  width: 100vw;
  height: 100vh;
  overflow: hidden;
}
.main {
  width: 100%;
  margin-top: 10px;
}
.main-col {
  height: 100vh;
  overflow: auto;
}
</style>
