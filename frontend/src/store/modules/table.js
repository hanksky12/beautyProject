export const moduleTable = {
  namespaced: true,
  state: {
    selected_array: [],
  },
  mutations: {
    setTableSelected(state, data) {
        // console.log(`[Vuex] setTableSelected 參數 ${JSON.stringify(data)}`)
        state.selected_array = data
    }
  }
}