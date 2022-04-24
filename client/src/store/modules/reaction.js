import axios from 'axios'

export default {
  namespaced: true,
  state: {
    like: '/assets/like.svg',
    love: '/assets/love.svg',
    haha: '/assets/haha.svg',
    wow: '/assets/wow.svg',
    sad: '/assets/sad.svg',
    angry: '/assets/angry.svg',
  },
  actions: {
    async getReaction(_, postId) {
      return axios({ url: `/react/${postId}` })
    },
    async getReactionType(_, postId) {
      return axios({ url: `/react/u/${postId}` }).then((r) => r.data)
    },
    async react(_, { postId, type }) {
      return axios({
        method: 'PUT',
        url: `/react/${postId}/${type}`,
      })
    },
  },
}
