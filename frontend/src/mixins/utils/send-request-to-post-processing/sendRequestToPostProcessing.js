import { SendRequestMixin } from "@/mixins/utils/handle-request/sendRequest";
import {CreateRequestMixin} from "@/mixins/utils/handle-request/create-request/createRequest";
import {PrintToastMixin} from "./printToast";
import {HandleTokenInvalidationMixin} from "./handleTokenInvalidation";
import {RefreshTableMixin} from "./refreshTable";

export const SendRequestToPostProcessingMixin =  {
    mixins: [
        SendRequestMixin,
        CreateRequestMixin,
        PrintToastMixin,
        HandleTokenInvalidationMixin,
        RefreshTableMixin
    ],
    methods: {
      async sendRequestToPostProcessing(requestData, isRefreshTable=true, isPrintToast=true) {
          console.log("sendRequestToPostProcessing")
          const request = this.createRequest(requestData)
          let responseData = await this.sendRequest(request)
          if (responseData.message === "Token has expired"){
              this.handleTokenInvalidation()
              responseData.message = "登入逾時，請重新登入"
          }
          if (isPrintToast){
                this.printToast(responseData)
          }
          if (isRefreshTable){
                this.refreshTable()
          }
          return responseData
      },
    },
};
