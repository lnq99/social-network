<template :key="id" class="row" style="flex: 1">
  <div class="desc">{{ data.mutual }} mutual friends</div>

  <div class="tag-btn" v-if="data.type == 'friend'">
    <el-tag size="small" type="success">Friend</el-tag>
    <div>
      <el-dropdown size="mini">
        <span class="el-dropdown-link">
          More<i class="el-icon-arrow-down el-icon--right"></i>
        </span>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item @click="onChangeRelationship('unfriend')"
              >Unfriend</el-dropdown-item
            >
            <el-dropdown-item @click="onChangeRelationship('block')" divided
              >Block</el-dropdown-item
            >
          </el-dropdown-menu>
        </template>
      </el-dropdown>
    </div>
  </div>
  <div class="tag-btn" v-else-if="data.type == 'request'">
    <el-tag size="small" type="warning">Request</el-tag>
    <div>
      <el-button
        class="btn"
        size="mini"
        type="primary"
        @click="onChangeRelationship('accept')"
        >Accept</el-button
      >
      <el-button
        class="btn"
        size="mini"
        type="info"
        @click="onChangeRelationship('delete')"
        >Delete</el-button
      >
    </div>
  </div>
  <div class="tag-btn" v-else-if="data.type == 'follow'">
    <el-tag size="small" type="info">Followed</el-tag>
    <el-button
      class="btn"
      size="mini"
      type="info"
      plain
      @click="onChangeRelationship('unfollow')"
      >Unsend</el-button
    >
  </div>
  <div class="tag-btn" v-else-if="data.type == 'block'">
    <el-tag size="small" type="danger">Block</el-tag>
    <el-button
      class="btn"
      size="mini"
      type="danger"
      plain
      @click="onChangeRelationship('unblock')"
      >Unblock</el-button
    >
  </div>
  <div class="tag-btn" v-else>
    <span></span>
    <el-button
      class=""
      size="mini"
      type="success"
      plain
      @click="onChangeRelationship('request')"
      >Send friend request</el-button
    >
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  props: ['id', 'user'],
  data() {
    return { data: {} }
  },
  methods: {
    ...mapActions({
      getMutualAndType: 'relationship/getMutualAndType',
      changeRelationship: 'relationship/changeRelationship'
    }),
    onChangeRelationship(type) {
      // console.log(this.id, type)
      this.changeRelationship([this.id, type]).then(
        _ => this.load()
      )
    },
    load() {
      this.getMutualAndType(this.id).then(r => {
        this.data = {
          type: r.type,
          mutual: r.mutual.length
        }
      })
    }
  },
  created() {
    if (!this.user) {
      this.load()
    } else {
      this.data = this.user
    }
  }
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
  /* margin-left: 24px; */
  font-size: 0.8em;
  opacity: 0.7;
  text-align: right;
  width: max(100px, 12%);
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
