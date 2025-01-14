import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'

export const RegisterApiMixin =   {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async registerApi(form_data) {
            console.log(form_data)
            const requestData= {
                "path":`/user/user`,
                "method":"post",
                "body":{
                    "user_name": form_data.user_name,
                    "user_password": form_data.user_password,
                }
            }
            return await this.sendRequestToPostProcessing(requestData,false)
        },
    },
};