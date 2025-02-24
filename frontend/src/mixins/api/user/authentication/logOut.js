import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'

export const LogOutApiMixin =   {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async logOutApi(form_data) {
            console.log(form_data)
            const requestData={
                "path":`/user/authentication/logout`,
                "method":"get",
            }
            return await this.sendRequestToPostProcessing(requestData,false)
        },
    },
};