export const moduleForm = {
    namespaced: true,
    state: {
        form_data: {},
    },
    getters: {
        getFormData: state => {
          //篩選裡面的 null 不加入
            let temp = {}
            for (let key in state.form_data) {
                if (state.form_data[key] != null && state.form_data[key] != '') {
                    temp[key] = state.form_data[key]
                }
            }
            // console.log(`原始form: ${JSON.stringify(state.form_data)}`)
            // console.log(`篩選form: 參數 ${JSON.stringify(temp)}`)
            return temp
        }
      },
    mutations: {
        resetFormData(state) {
            state.form_data = {}
        },
        resetAndKeepPartFormData(state, payload) {
            // keep payload part of form data
            let temp = {}
            for (let key in state.form_data) {
                if (payload.includes(key)) {
                    temp[key] = state.form_data[key]
                }
            }
            state.form_data = temp
        },
        setColumn(state, payload) {
            // console.log(`setColumn 參數 ${JSON.stringify(payload)}`)
            state.form_data = {...state.form_data, ...payload};
        },
    },
}
