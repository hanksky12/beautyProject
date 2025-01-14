export const GetOperateToSendRequestMixin = {
    inject: ['showToast', 'showConfirmModal'],
    methods: {
        getOperateToSendRequest(getFormData, correspond_operate_obj, selected_array) {
        console.log('Operate onSubmit~~~~')
        const {operate_name,operate_func} = this.$_getOperate(getFormData, correspond_operate_obj)
        if (selected_array.length === 0) {
            this.showToast("没有选择资料!!", "warning");
            return
        }
        if (operate_func === null||operate_func === undefined) {
            this.showToast("没有选择功能!!", "warning");
            return
        }
        this.showConfirmModal(
            "操作确认",
            `是否执行${operate_name} 共 ${selected_array.length} 筆?`,
            operate_func,
            selected_array)
        },
        $_getOperate(getFormData, correspond_operate_obj) {
            console.log('getOperate')
            const operate_name = getFormData["function_operate"]
            const operate_func = (operate_name===null) ? null : correspond_operate_obj[operate_name]
            return {operate_name,operate_func}
        },
    },
};
