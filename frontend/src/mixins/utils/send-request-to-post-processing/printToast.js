export const PrintToastMixin = {
    inject: ['showToast'],
    methods: {
        printToast(responseData) {
            const message = this.$_createMessage(responseData);
            const level = this.$_createLevel(responseData);
            console.log(message)
            console.log("結果: " + responseData.is_success)
            console.log("code: " + responseData.code)
            console.log(`訊息:${JSON.stringify(message,null,10)}`)
            this.showToast(message, level)
        },
        $_createLevel(responseData) {
            if (responseData.is_success) {
                return (responseData.code === 200) ? 'success' : 'warning'
            } else {
                return 'danger'
            }
        },
        $_createMessage(responseData) {
            let message = ''
            if (responseData.is_success && responseData.code === 200) {
                message = responseData.data
                if ('failed_list' in message||'success_list' in message){
                    message = this.$_replaceObjectKeys(message);
                    message = this.$_replaceObjectValues(message);
                }
            }
            else {
                message = responseData.message
            }
            return message
        },
        $_replaceObjectKeys(obj) {
            const keyMap = {
                        "failed_list": "失敗清單",
                        "success_list": "成功清單"
                    }
            return Object.fromEntries(
                Object.entries(obj).map(([oldKey, value]) => [keyMap[oldKey] || oldKey, value]));
        },
        $_replaceObjectValues: function (message) {
            return {
                ...message,
                "失敗清單": this.$_replaceKeyInList(message["失敗清單"], "upload_result_id", "上傳結果ID"),
                "成功清單": this.$_replaceKeyInList(message["成功清單"], "upload_result_id", "上傳結果ID")
            };
        },
        $_replaceKeyInList(list, oldKey, newKey) {
            return list.map(item => {
                if (item[oldKey] !== undefined) {
                    item[newKey] = item[oldKey];
                    delete item[oldKey];
                }
                return item;
            });
        }
    },
};
