<template>
  <el-tree v-if="loaded" node-key="id" :data="data" :props="tree" @node-click="replyTo">
    <template #default="{ node, data }">
      <div class="cmt">
        <short-info :id="data.userId">
          <template v-slot="slotProps">
            <span>
              <el-avatar class="ava" :size="30" :src="slotProps.avatars"></el-avatar>
            </span>
            <span>
              <card class="p12 box-shadow">
                <div class="cmt-header-r">
                  <span class="cmt-author">{{ slotProps.uname }}</span>
                  <time class="cmt-time">{{ data.created }}</time>
                </div>
                <div>{{ node.label }}</div>
              </card>
            </span>
          </template>
        </short-info>
      </div>
    </template>
  </el-tree>
  <el-tag
    class="tag"
    v-if="reply != 0"
    :key="reply"
    closable
    @close="reply = 0"
  >Reply to {{ reply }}</el-tag>
  <comment-input @cmt="postCmt" class="cmt-input"></comment-input>
</template>

<script>
import { mapActions } from 'vuex'
import CommentInput from './CommentInput.vue'

export default {
  props: ['postId'],
  components: { CommentInput },
  data() {
    return {
      loaded: false,
      reply: 0,
      tree: {
        children: 'children',
        label: 'content',
      },
    }
  },
  methods: {
    ...mapActions({ getCmtTree: 'cmt/getCmtTree', postComment: 'cmt/comment' }),
    replyTo(data) {
      this.reply = data.id
    },
    postCmt(content) {
      this.postComment({
        postId: this.postId,
        parentId: this.reply,
        content: content
      }).then(() => {
        this.reply = 0
        this.getCmtTree(this.postId).then(data => {
          this.data = data
          this.loaded = true
        })
      })
    }
  },
  created() {
    this.getCmtTree(this.postId).then(data => {
      this.data = data
      this.loaded = true
      // console.log(JSON.stringify(data, null, 2))
    })
  },
}
</script>

<style scoped>
.el-tree {
  background: none;
  padding: 0 18px 0 18px;
}
.el-tree-node {
  margin: 5px;
}
.cmt {
  display: flex;
  padding-bottom: 4px;
  opacity: 0.9;
}
.cmt-header-r {
  font-size: 0.9em;
  margin-bottom: 4px;
}
.cmt-author {
  font-weight: 600;
}
.cmt-time {
  padding-left: 12px;
  opacity: 0.8;
  font-size: 0.8em;
}
.ava {
  margin-right: 8px;
}
.card {
  display: inline-block;
}
.tag {
  margin: 12px 0 -12px 60px;
}
</style>
