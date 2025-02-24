import {createNamespacedHelpers} from "vuex";
const { mapState:mapStateConst} = createNamespacedHelpers('const')
export const GetDomainMixin = {
    methods: {
         getDomain(domainStr,protocol="http") {
            if (domainStr === undefined) {
                if (protocol==="http") {
                    return this.const.domain
                } else if (protocol==="ws") {
                    return this.const.domain.replace("http", "ws")
                }
            } else {
                return null
            }
        },
    },
    computed: {
    ...mapStateConst({
        'const': state => state.const,
    }),
    },
};
