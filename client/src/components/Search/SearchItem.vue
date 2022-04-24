<template>
  <card class="search-item row box-shadow">
    <div class="row">
      <link-card :link="{ name: 'Profile', params: { id: user.id } }">
        <short-info :id="user.id">
          <template v-slot="slotProps">
            <template class="row hover p18">
              <span>
                <el-avatar
                  class="ava"
                  :size="40"
                  :src="slotProps.avatars"
                ></el-avatar>
              </span>
              <div class="name">{{ slotProps.uname }}</div>
            </template>
          </template>
        </short-info>
      </link-card>
    </div>

    <relation :id="user.id" :user="user"></relation>
  </card>
</template>

<script>
import Relation from '../Relationship/Relation.vue'
export default {
  components: { Relation },
  props: ['user'],
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
  width: 140px;
  text-align: left;
  white-space: nowrap;
  overflow: hidden;
}
.desc {
  margin-left: 24px;
  font-size: 0.8em;
  opacity: 0.7;
}
.search-item {
  margin-bottom: 18px;
  padding-right: 18px;
  /* justify-content: space-between; */
}
.btn {
  width: 80px;
}
.el-tag {
  margin-left: 10px;
  width: 64px;
  text-align: center;
  opacity: 0.8;
}
.tag-btn {
  display: flex;
  flex: 1;
  justify-content: space-between;
}
</style>
