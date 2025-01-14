<template>
  <div>
    <b-modal  id="modal-prevent-closing" ref="modal" size="lg" :title="title"  @hidden="resetModal" @ok="handleOk">
        <pre>{{ content }}</pre>
    </b-modal>
  </div>
</template>


<script>
export default {
    //確認是否要執行
    inject: ['showToast'],
    data() {
        return {
            title: null,
            content: null,
            op_func: null,
            op_data: null,
        };
    },
    methods: {
        show(title, content, op_func, op_data) {
            //op_data不一定是selected_array 所以不從state取
            this.title = title;
            this.op_func = op_func;
            this.op_data = op_data;
            this.content = content;
            this.$refs.modal.show()
        },
        resetModal() {
            this.$nextTick(() => {
            this.title = '';
            this.content ='';
          });
        },
        handleOk(bvModalEvent) {
            bvModalEvent.preventDefault()
            this.$nextTick(() => {
                this.$bvModal.hide('modal-prevent-closing')
            })
            console.log('Confirm Ok')
            this.showToast("開始执行操作");
            this.op_func(this.op_data)
        },
    },
    watch: {
        props_op_func: function (newVal) {
            this.op_func = newVal
        },
        props_op_data: function (newVal) {
            this.op_data = newVal
        }
    }
}
</script>