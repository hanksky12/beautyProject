export const moduleCrudOperate = {
    namespaced: true,
    state: {
        operate: "",
    },
       getters: {
        getOperate: state => {
            return state.operate
        }
      },

    mutations: {
        setOperate(state, data) {
            // console.log(`setQueryClick 參數 ${JSON.stringify(data)}`)
            state.operate = data
        },
        resetOperate(state) {
            state.operate = ""
        },
    }
}