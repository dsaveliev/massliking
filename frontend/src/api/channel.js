import auth from './auth'
import router from '../router'

const API_URL = process.env.API_URL

export default {

  isEmptyChannelLists (context) {
    return context.$store.state.instagramList.some((instagram, _index, _array) => {
      !instagram.channels
    })
  },

  updateAllChannels (context) {
    context.$store.state.instagramList.forEach((instagram, _index, _array) => {
      this.updateAccountChannels(context, instagram, false)
    })
  },

  updateAccountChannels (context, instagram, redirect) {
    var url = API_URL + 'instagram/' + instagram.id + '/channels'
    context.$http.get(url, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('updateChannelList', { instagram: instagram, channelList: response.body })
        context.$forceUpdate()
        if (redirect) { router.push({path: redirect}) }
      }, error => {
        console.log(error)
        // TODO: Переделать на правильную ошибку.
        // if (error.status === 400) { context.errorInvalidCredentials = true }
      })
  },

  createChannel (context, instagram, form, redirect) {
    var url = API_URL + 'instagram/' + instagram.id + '/channels'
    context.$http.post(url, form, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('addChannel', { instagram: instagram, channel: response.body })
        if (redirect) { router.push({path: redirect}) }
      }, error => {
        console.log(error)
      })
  },

  editChannel (context, instagram, channel, form, redirect) {
    var url = API_URL + 'instagram/' + instagram.id + '/channels/' + channel.id
    context.$http.put(url, form, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('updateChannel', { instagram: instagram, channel: response.body })
        if (redirect) { router.push({path: redirect}) }
      }, error => {
        console.log(error)
      })
  },

  deleteChannel (context, instagram, channel, redirect) {
    var url = API_URL + 'instagram/' + instagram.id + '/channels/' + channel.id
    context.$http.delete(url, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('deleteChannel', { instagram: instagram, channel: channel })
        if (redirect) { router.push({path: redirect}) }
      }, error => {
        console.log(error)
      })
  },

  startChannel (context, instagram, channel) {
    var url = API_URL + 'instagram/' + instagram.id + '/channels/' + channel.id + '/start'
    context.$http.get(url, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('updateChannel', { instagram: instagram, channel: response.body })
        context.$forceUpdate()
      }, error => {
        console.log(error)
      })
  },

  stopChannel (context, instagram, channel) {
    var url = API_URL + 'instagram/' + instagram.id + '/channels/' + channel.id + '/stop'
    context.$http.get(url, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('updateChannel', { instagram: instagram, channel: response.body })
        context.$forceUpdate()
      }, error => {
        console.log(error)
      })
  }

}
