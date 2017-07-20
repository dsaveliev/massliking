import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export const store = new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  state: {
    user: {},
    instagramList: []
  },
  getters: {
    findInstagram: (state, getters) => (id) => {
      return state.instagramList.find(instagram => instagram.id.toString() === id.toString())
    },
    instagram: (state, getters) => (id) => {
      return getters.findInstagram(id)
    },
    findChannel: (state, getters) => (instagramId, id) => {
      var instagram = getters.findInstagram(instagramId)
      return instagram.channels.find(channel => channel.id.toString() === id.toString())
    },
    channel: (state, getters) => (instagramId, id) => {
      return getters.findChannel(instagramId, id)
    }
  },
  mutations: {
    updateUser (state, payload) {
      state.user = { ...state.user, ...payload.user }
    },
    deleteUser (state) {
      state.user = {}
    },
    updateInstagramList (state, payload) {
      payload.instagramList.forEach((instagram, _index, _account) => {
        var index = state.instagramList.findIndex(i => { return i.id === instagram.id })
        if (instagram.channels === undefined || instagram.channels === null) {
          instagram.channels = []
        }
        if (index === -1) {
          state.instagramList.push(instagram)
        }
        else {
          state.instagramList[index] = { ...state.instagramList[index], ...instagram }
        }
      })
    },
    updateInstagram (state, payload) {
      var index = state.instagramList.findIndex(i => { return i.id === payload.instagram.id })
      state.instagramList = Object.assign([], state.instagramList, {[index]: payload.instagram})
    },
    addInstagram (state, payload) {
      state.instagramList = [ ...state.instagramList, payload.instagram ]
    },
    deleteInstagram (state, payload) {
      var index = state.instagramList.indexOf(payload.instagram)
      state.instagramList.splice(index, 1)
    },
    deleteInstagramList (state) {
      state.instagramList = []
    },
    updateChannelList (state, payload) {
      var instagram = state.instagramList.find(instagram => instagram.id === payload.instagram.id)
      if (!instagram.channel) {
        Vue.set(instagram, 'channels', payload.channelList)
      }
      else {
        payload.channelList.forEach((channel, _index, _account) => {
          var index = instagram.channels.findIndex(c => { return c.id.toString() === channel.id.toString() })
          if (index === -1) {
            instagram.channels.push(channel)
          }
          else {
            instagram.channels[index] = { ...instagram.channels[index], ...channel }
          }
        })
      }
    },
    addChannel (state, payload) {
      var instagram = state.instagramList.find(instagram => instagram.id === payload.instagram.id)
      instagram.channels = [ ...instagram.channels, payload.channel ]
    },
    updateChannel (state, payload) {
      var instagram = state.instagramList.find(instagram => instagram.id === payload.instagram.id)
      var index = instagram.channels.findIndex(c => { return c.id.toString() === payload.channel.id.toString() })
      instagram.channels = Object.assign([], instagram.channels, {[index]: payload.channel})
    },
    deleteChannel (state, payload) {
      var instagram = state.instagramList.find(instagram => instagram.id === payload.instagram.id)
      var index = instagram.channels.findIndex(c => { return c.id.toString() === payload.channel.id.toString() })
      instagram.channels.splice(index, 1)
    }
  }
})
