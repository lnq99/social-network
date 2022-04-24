import axios from 'axios'

export default {
  namespaced: true,
  actions: {
    async getCmtTree(_, postId) {
      return axios({ url: `/cmt/${postId}` })
    },
    async comment(_, cmtBody) {
      let options = {
        method: 'POST',
        url: `/cmt`,
        data: cmtBody,
      }
      return axios(options)
    },
  },
}
