import {createNamespacedHelpers} from "vuex";
const { mapState:mapStateConst} = createNamespacedHelpers('const')
import {GetCookieMixin} from "./getCookie";
import {GetDomainMixin} from "./getDomain";

export const CreateRequestMixin = {
    mixins: [GetCookieMixin,GetDomainMixin],
    methods: {
        createRequest(requestData, protocol="http") {
            console.log("createRequest")
            const url = this.getDomain(requestData.domain, protocol) + requestData.path
            const headers = this.$_getHeaders(requestData)
            return this.$_createRequest(url, headers, requestData.method, requestData.body)
        },
        $_getHeaders: function (baseRequestData) {
             let headers = this.const.headers
             if (baseRequestData.method === 'post' || baseRequestData.method === 'put'  || baseRequestData.method === 'delete')
             {
                 headers['X-CSRF-TOKEN'] = this.getCookie("csrf_access_token")
             }
             return headers
        },
        $_createRequest(url,headers,method,body) {
            let init = {}
            if (method === "get") {
              url = url + this.$_getGetParams(body)
              init={
                  method: method,
                  headers: new Headers(headers),
                  // credentials: 'include'
              }
            } else {
              init={
                  method: method,
                  headers: new Headers(headers),
                  body: JSON.stringify(body),
                  // credentials: 'include'
              }
            }
            console.log(`url=${JSON.stringify(url,null,2)}`)
            console.log(`method=${JSON.stringify(method,null,2)}`)
            console.log(`headers=${JSON.stringify(headers,null,2)}`)
            console.log(`body=${JSON.stringify(body,null,2)}`)
            return new Request(url,init)


      },
      $_getGetParams: function (body) {
            let params = "?"
            for (const key in body) {
                params += `${key}=${body[key]}&`
            }
            return params
      }

    },
    computed: {
    ...mapStateConst({
        'const': state => state.const,
    }),
    },
};