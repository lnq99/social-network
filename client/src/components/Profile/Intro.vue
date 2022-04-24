<template>
  <h2>
    Intro
    <el-button v-if="editable" class="btn-right" @click="isEdit = true">
      Edit
    </el-button>
  </h2>
  <p v-for="i in introContent.split('\n')">{{ i }}</p>
  <el-dialog
    v-if="editable"
    title="Edit Intro"
    v-model="isEdit"
    width="40%"
    center
  >
    <el-input
      type="textarea"
      :autosize="{ minRows: 2, maxRows: 4 }"
      v-model="introContent"
    >
    </el-input>
    <template #footer>
      <span class="dialog-footer">
        <el-button type="primary" @click="onSaveIntro">Save</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script>
import { ref } from 'vue'
import { mapActions } from 'vuex'

export default {
  props: ['intro', 'editable'],
  data() {
    return {
      isEdit: false,
      introContent: ref(''),
    }
  },
  methods: {
    ...mapActions({ saveIntro: 'profile/saveIntro' }),
    onSaveIntro() {
      this.saveIntro(this.introContent).then(() => {

        this.isEdit = false
        console.log(this.intro)
      })
    },
  },
  created() {
    if (this.intro.length == 0) {
      this.introContent = 'Works at ...\nStudied at \nLives in ...\nFrom ...'
    } else {
      this.introContent = this.intro
    }
    this.introContent = this.intro.replaceAll('\\n', '\n')
  },
}
</script>

<style scoped>
.btn-right {
  float: right;
  margin-top: -6px;
  padding-top: 4px;
  padding-bottom: 4px;
}
p {
  margin-bottom: 0px;
}
</style>
