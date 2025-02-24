import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'

export const GetTokenApiMixin =   {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async getTokenApi() {
            const requestData= {
                "path":"/user/authentication/token",
                "method":"get",
                "body":{}
            }
            return await this.sendRequestToPostProcessing(requestData,false,false)
        },
    },
};