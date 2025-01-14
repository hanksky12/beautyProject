export const moduleQuery = {
    namespaced: true,
    state: {
        is_query_click: false,
    },
    mutations: {
        setQueryClick(state, data) {
            // console.log(`setQueryClick 參數 ${JSON.stringify(data)}`)
            state.is_query_click = data
        }
    }
}