import './axios.config.js'
import { createStore } from 'vuex'
import auth from './modules/auth.js'
import profile from './modules/profile.js'
import post from './modules/post.js'
import cmt from './modules/cmt.js'
import reaction from './modules/reaction.js'
import relationship from './modules/relationship.js'
import photo from './modules/photo.js'
import notif from './modules/notif.js'
import theme from './modules/theme.js'

const store = createStore({
  state: {
    id: 0,
  },
  mutations: {
    setRootId(state, id) {
      state.id = id
    },
  },
  modules: {
    auth,
    profile,
    post,
    cmt,
    reaction,
    relationship,
    photo,
    notif,
    theme,
  },
})

export default store
