import { getCookie, overwriteCookie } from '@/utils.js'

export default {
  namespaced: true,
  state: {
    isDark: getCookie('dark') == 'true',
  },
  mutations: {
    switchTheme(state) {
      state.isDark = !state.isDark
    },
  },
  actions: {
    switchTheme({ state, commit }) {
      commit('switchTheme')
      overwriteCookie('dark', state.isDark)
    },
  },
}
