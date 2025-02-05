import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'
export const StopRecordApiMixin = {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async stopRecordApi(selections) {
            const requestData = {
                "path": "/hardware/status",
                "method": "post",
                "body": {
                    "state": "stop",
                    "hardware": selections[0].hardware_name,
                }
            }
            await this.sendRequestToPostProcessing(requestData, false)
        }
    },
}