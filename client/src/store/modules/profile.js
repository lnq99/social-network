import axios from 'axios'

export default {
  namespaced: true,
  state: {
    id: 0,
    feed: [],
    shortProfiles: {},
    name: '',
    avatarl: '',
  },
  mutations: {
    cacheShortProfileArray(state, arr) {
      for (let p of arr) {
        state.shortProfiles[p.id] = p
      }
    },
    cacheShortProfile(state, p) {
      state.shortProfiles[p.id] = p
    },
    initProfile(state, profile) {
      state.id = profile.id
      state.name = profile.name
      state.avatarl = profile.avatarl
      state.avatars = profile.avatars
      state.intro = profile.intro
    },
  },
  actions: {
    async getProfile(_, id) {
      let options = {
        method: 'GET',
        // baseURL: '',
        url: `/profile/${id}`,
      }
      return axios(options).catch((err) => {
        console.log(err)
      })
    },
    async getProfileShort({ state, commit }, id) {
      let p = state.shortProfiles[id]
      if (p) {
        // console.log('cache hit', id)
        return p
      }
      let options = {
        method: 'GET',
        url: `/profile/short/${id}`,
      }
      return axios(options).then((data) => {
        // console.log('cache missed', id)
        commit('cacheShortProfile', data)
        return data
      })
    },
    async searchProfile(_, key) {
      let options = {
        method: 'GET',
        url: `/search?k=${key}`,
      }
      return axios(options).then((data) => {
        return data
      })
    },
    async saveIntro({ state }, intro) {
      state.intro = intro
      let options = {
        method: 'PATCH',
        url: `/profile/info`,
        data: {
          info: intro,
        },
      }
      return axios(options).then((data) => {
        return data
      })
    },
  },
  getters: {
    intro(state) {
      return state.intro
    },
  },
}
