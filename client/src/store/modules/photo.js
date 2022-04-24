import axios from 'axios'

export default {
  namespaced: true,
  actions: {
    async getPhoto(_, id) {
      let options = {
        method: 'GET',
        url: `/photo/${id}`,
      }
      return axios(options)
    },
    async getPhotosOfProfile(_, profileId) {
      let options = {
        method: 'GET',
        url: `/photo/u/${profileId}`,
      }
      return axios(options)
    },
  },
}
