<template>
  <div class="row">
    <div class="btn-contaner">
      <el-popover placement="top" trigger="hover" :width="180">
        <img
          class="react-icon"
          :class="{ 'react-icon-on': i == liked }"
          @click="onReact(i)"
          v-for="(url, i) in icons"
          :src="url"
        />
        <template #reference>
          <el-button
            class="btn"
            :class="{ 'btn-on': liked != -1 }"
            type="primary"
            @click="onReact(-1)"
            size="small"
          >
            <div class="btn-icon" v-if="liked != -1">
              {{ iconsName[liked] }}
              <img :src="icons[liked]" />
            </div>
            <template v-else>Like</template>
          </el-button>
        </template>
      </el-popover>
      <el-button class="btn" size="small" @click="comment = !comment">{{ cmtCount }} Comments</el-button>
    </div>
    <div class="reacts">
      <div class="react" v-for="(v, i) in reaction">
        {{ v }}
        <img :src="icons[i]" />
      </div>
    </div>
  </div>
  <div v-if="comment">
    <hr />
    <comments :postId="id"></comments>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Comments from '../Post/Comments.vue'

const react_t = ['like', 'love', 'haha', 'wow', 'sad', 'angry']

export default {
  components: { Comments },
  props: ['id', 'initReaction', 'cmtCount'],
  data() {
    return {
      comment: false,
      choice: 0,
      iconsName: ['Like', 'Love', 'Haha', 'Wow', 'Sad', 'Angry'],
      liked: -1,
      reaction: [],
    }
  },
  computed: {
    ...mapState('reaction', react_t),
  },
  methods: {
    ...mapActions({ getReactionType: 'reaction/getReactionType', react: 'reaction/react', getReaction: 'reaction/getReaction' }),
    onReact(n) {
      if (n == -1 && this.liked != -1)
        this.liked = -1
      else if (n != -1)
        this.liked = n
      else
        this.liked = 0
    },
  },
  created() {
    this.reaction = this.initReaction
    this.getReactionType(this.id).then(t => {
      this.liked = react_t.findIndex((v) => v == t)
      this.$watch('liked', (newVal, oldVal) => {
        let t
        if (newVal == -1) t = 'del'
        else t = react_t[newVal]
        this.react({ postId: this.id, type: t }).then(() =>
          this.getReaction(this.id).then(r => this.reaction = r)
        )
      })
    })
    this.icons = [
      this.like, this.love, this.haha, this.wow, this.sad, this.angry
    ]
  }
}
</script>
<style lang="scss" scoped>
@import "@/style.scss";

.reacts {
  display: inline-flex;
  margin-bottom: 0;
  opacity: 0.75;
}
.react {
  display: flex;
  justify-content: flex-end;
  font-size: 0.9em;
  font-weight: lighter;
  width: 40px;
}
img {
  margin-left: 4px;
  width: 1em;
}
.btn-contaner {
  width: 50%;
  display: flex;
}

.row {
  display: flex;
  // justify-content: flex-start;
  justify-content: space-between;
  align-items: center;
  margin: 0 $p4 0 $p4;
}

.btn-icon {
  display: flex;
  justify-content: center;
}

.btn {
  color: var(--fg);
  font-weight: 300;
  width: 50%;
  /* width: max(25%, 100px); */
}
.btn,
.btn:focus {
  background: none;
}

.btn-on,
.btn-on:focus {
  background: #2080ffb0;
}

.react-icon {
  width: 1.8em;
  opacity: 0.7;
}
.react-icon:hover {
  opacity: 0.9;
  transform: scale(1.3);
}
.react-icon-on {
  width: 1.8em;
  opacity: 1;
  border: 1px solid lightcoral;
  border-radius: 100px;
}
</style>
