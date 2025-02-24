import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'

export const GetMouseActionNameApiMixin = {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async getMouseActionNameApi() {
            const requestData = {
                "path": "/user/mouse/info",
                "method": "get",
                "body": {}
            };
            return await this.sendRequestToPostProcessing(requestData, false, false);
        },
    },
};