import {createNamespacedHelpers} from "vuex";
const { mapState:mapStateConst} = createNamespacedHelpers('const')
export const GetDomainMixin = {
    methods: {
         getDomain(domainStr) {
            if (domainStr === undefined) {
                return this.const.domain
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
