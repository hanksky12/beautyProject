import { SendRequestToPostProcessingMixin } from '@/mixins/utils/send-request-to-post-processing/sendRequestToPostProcessing.js'

export const LogInApiMixin =   {
    mixins: [SendRequestToPostProcessingMixin],
    methods: {
        async logInApi(form_data) {
            console.log(form_data)
            const requestData= {
                "path":"/user/authentication/login",
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