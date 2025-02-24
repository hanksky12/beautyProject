import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'

export const GetHardwareNameApiMixin = {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async getHardwareNameApi() {
            const requestData = {
                "path": "/server/hardware/info",
                "method": "get",
                "body": {}
            };
            return await this.sendRequestToPostProcessing(requestData, false, false);
        },
    },
};