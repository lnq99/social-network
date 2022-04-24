import axios from 'axios'

export default {
  namespaced: true,
  state: {
    feed: [],
    lim: 5,
    off: 0,
    // TODO: feed idea: from-to new Date().toISOString(),
  },
  mutations: {
    loadFeed(state, payload) {
      let data = payload
      state.feed = [...new Set([...state.feed, ...data])]
      // state.feed.sort().reverse()
    },
  },
  actions: {
    async post(_, postBody) {
      let options = {
        method: 'POST',
        url: `/post`,
        data: postBody,
      }
      // console.log(postBody)
      return axios(options)
    },
    async delete(_, id) {
      let options = {
        method: 'DELETE',
        url: `/post/${id}`,
      }
      return axios(options)
    },
    async getPost(_, id) {
      return axios({ url: `/post/${id}` })
    },
    async getPostsOfProfile(_, profileId) {
      return axios({ url: `/post/u/${profileId}` })
    },
    async getFeed({ state, commit, rootState }, payload) {
      let options = {
        method: 'GET',
        url: `/feed?lim=${state.lim}&off=${state.off}`,
      }
      return axios(options).then((data) => {
        state.off += state.lim
        commit('loadFeed', data)
      })
    },
  },
  getters: {
    feed(state) {
      return state.feed
    },
  },
}
