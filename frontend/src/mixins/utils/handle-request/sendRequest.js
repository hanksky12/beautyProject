
import {createNamespacedHelpers} from "vuex";
const { mapState:mapStateAjax} = createNamespacedHelpers('ajax')
export const SendRequestMixin =  {
    methods: {
      async sendRequest(requestData) {
        try {
          console.log('送出請求')
          await this.$store.dispatch("ajax/send", requestData)
        } catch (error) {
          console.error('Error fetching data:', error);
        }
        return this.responseData
      },
    },
    computed: {
    ...mapStateAjax({
        'responseData': state => state.responseData,
    }),
    },

};
