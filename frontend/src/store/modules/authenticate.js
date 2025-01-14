export const moduleAuthenticate = {
  namespaced: true,
  state: {
    isAuthenticated: false,
  },
  mutations: {
    setAuthenticated(state, data) {
        console.log(`[Vuex] setAuthenticated 參數 ${JSON.stringify(data)}`)
        state.isAuthenticated = data
    }
  }
}