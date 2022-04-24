import axios from 'axios'

export default {
  namespaced: true,
  state: {
    notif: [],
  },
  actions: {
    getNotif({ state }) {
      return axios({ url: '/notif' }).then((data) => {
        state.notif = data
      })
    },
  },
  getters: {
    notif(state) {
      return state.notif
    },
  },
}
