<template>
  <el-input
    type="textarea"
    :autosize="{ minRows: 6, maxRows: 10 }"
    placeholder="Please input"
    v-model="model.content"
  >
  </el-input>

  <div class="row">
    <el-tag
      :key="tag"
      v-for="tag in tags"
      closable
      :disable-transitions="false"
      @close="handleTagClose(tag)"
    >
      {{ tag }}
    </el-tag>
    <el-input
      class="input-new-tag"
      v-if="tagVisible"
      v-model="tagValue"
      ref="saveTagInput"
      size="mini"
      @keyup.enter.native="handleTagConfirm"
      @blur="handleTagConfirm"
    >
    </el-input>
    <el-button v-else class="button-new-tag" size="small" @click="showTagInput"
      >+ New Tag</el-button
    >
  </div>

  <div class="atch-group center">
    <el-radio-group v-model="atchRadio" size="mini">
      <el-radio-button label="0">No attachment</el-radio-button>
      <el-radio-button label="1">Photo Link</el-radio-button>
      <el-radio-button label="2">Photo Upload</el-radio-button>
      <el-radio-button label="3" disabled>Other</el-radio-button>
    </el-radio-group>
  </div>

  <div v-if="atchRadio == 1">
    <el-input placeholder="Photo URL" v-model="model.atchUrl" clearable>
    </el-input>
  </div>
  <div v-else-if="atchRadio == 2">
    <el-upload
      class="upload center"
      drag
      :on-success="handlePhotoSuccess"
      :on-preview="handlePhotoPreview"
      :on-remove="handlePhotoRemove"
      :file-list="fileList"
      list-type="picture"
      limit="1"
      :disabled="!showUpload"
    >
      <i class="el-icon-upload"></i>
      <div class="el-upload__text">
        Drop file here or <em>click to upload</em>
      </div>
      <template #tip>
        <div class="el-upload__tip">
          jpg/png files with a size less than 500kb
        </div>
      </template>
    </el-upload>
  </div>

  <center>
    <el-button
      class="btn-post center"
      type="primary"
      size="small"
      @click="onCreatePost"
      >Post</el-button
    >
  </center>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  data() {
    return {
      model: {
        content: '',
        atchType: '',
        // atchId: '',
        atchUrl: 'https://loremflickr.com/320/180/all',
      },
      tags: [],
      atchRadio: 1,
      showUpload: true,
      fileList: [],
      tagVisible: false,
      tagValue: ''
    }
  },
  methods: {
    ...mapActions({ postPost: 'post/post' }),
    onCreatePost() {
      if (this.atchRadio == 1 && this.model.atchUrl != '') {
        this.model.atchType = 'photo'
      } else {
        this.model.atchType = 'none'
      }
      this.postPost({
        ...this.model,
        tags: this.tags.join('-')
      }).then(() => {
        this.$emit('posted')
      })
    },

    handlePhotoRemove(file, fileList) {
      console.log(file, fileList)
      this.showUpload = true
    },
    handlePhotoCardPreview(file) {
      this.dialogImageUrl = file.url
      this.dialogVisible = true
    },
    handlePhotoSuccess(file, fileList) {
      console.log(file, fileList)
      this.showUpload = false
    },

    handleTagClose(tag) {
      this.tags.splice(this.tags.indexOf(tag), 1)
    },
    showTagInput() {
      this.tagVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.tag.focus()
      })
    },
    handleTagConfirm() {
      let tagValue = this.tagValue
      if (tagValue) {
        this.tags.push(tagValue)
      }
      this.tagVisible = false
      this.tagValue = ''
    }
  },
};
</script>

<style scoped>
.atch-group {
  margin: 20px 0 20px;
}
.btn-post {
  margin-top: 20px;
  width: min(40%, 100px);
}
</style>
