class ResponseData {
  constructor(data, message, code, is_success = true) {
    this.data = data;
    this.message = message;
    this.is_success = is_success;
    this.code = code
  }
}

export const moduleAjax = {
  namespaced: true,
  state: {
      responseData: null
  },

  mutations: {
    setResponseData(state, responseData) {
          state.responseData = responseData
        },
  },
  actions: {
    async send({ commit,dispatch }, request) {
        try{
          // console.log("send")
          let response = await fetch(
              request,
              {credentials: 'include'});
            console.log(response);
          let jsonResponse = await response.json();
          await dispatch("processStatus",jsonResponse);
        }catch (error) {
            console.error('Error fetching data:', error);
            let responseData=   new ResponseData({}, error.message, 0,false)
            commit('setResponseData', responseData);
        }
    },
    async processStatus({dispatch},response) {
      console.log("processStatus")
      console.log(response)
      if (response.code >= 200 && response.code < 300) {
        console.log("http 200-300")
          await dispatch("successHandling",response)
      } else {
        console.log("非 http 200")
         await dispatch("failHandling",response)
      }
    },

    async successHandling({commit},response) {
      console.log("successHandling")
      let responseData=   new ResponseData(
          response.data,
          JSON.stringify(response.message),
          response.code,
          true)
      commit('setResponseData', responseData);
      },

    async failHandling({commit},response) {
        let message = "預設錯誤訊息"
        if (Object.prototype.hasOwnProperty.call(response, "msg")) {
          console.log("解析msg")
          message = response.msg
        } else if (Object.prototype.hasOwnProperty.call(response, "message")) {
          console.log("解析message")
          message = response.message
        } else if (Object.prototype.hasOwnProperty.call(response, "messages")) {
          console.log("解析messages")
          message = response.messages.json
        } else {
          console.log("解析errors.json ")
          message = response.errors.json
        }
        let responseData=   new ResponseData(
            {},
            message,
            response.code,
            false)
        commit('setResponseData', responseData);
      }
  }
}