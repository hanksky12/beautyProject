export const moduleAuthorization = {
  namespaced: true,
  state: {
    level: "",
  },
  mutations: {
    setAuthorization(state, data) {
        console.log(`[Vuex] setAuthorization 參數 ${JSON.stringify(data)}`)
        state.level = data
    }
  }
}