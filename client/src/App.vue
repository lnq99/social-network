<template>
  <div id="app" :class="this.$store.state.theme.isDark ? 'dark' : 'light'">
    <router-view />
  </div>
</template>

<style lang="scss">
@import "./style.scss";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--text);
  background: var(--bg);
  height: 100%;
  width: 100%;
  display: flex;
  // letter-spacing: 0.1px;
  word-spacing: 0.1px;
  line-height: 120%;
}

.input .el-textarea__inner {
  border: none;
  background-color: var(--bg);
  color: var(--text);
  opacity: 0.8;
}

.el-tree-node__content {
  color: var(--text) !important;
  padding-top: $p1 !important;
  padding-bottom: $p1 !important;
  height: initial !important;
  .card {
    background-color: var(--bg) !important;
  }
}

.el-tree-node {
  :hover,
  :focus {
    background-color: initial !important;
  }
  white-space: normal !important;
}

.el-tree__empty-block {
  min-height: 20px;
}

.ava {
  opacity: 0.9;
}
</style>
