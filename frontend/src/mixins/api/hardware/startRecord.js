import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'
export const StartRecordApiMixin = {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async startRecordApi(selections) {
            console.log("startRecordApi")
            console.log(selections)
            const requestData = {
                "path": "/hardware/status",
                "method": "post",
                "body": {
                    "state": "start",
                    "hardware": selections[0].hardware_name,
                }
            }
            await this.sendRequestToPostProcessing(requestData, false)
        }
    },

}