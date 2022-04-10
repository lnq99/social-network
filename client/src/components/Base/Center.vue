<template>
  <div class="center">
    <slot></slot>
  </div>
</template>

<style scoped>
.center {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
