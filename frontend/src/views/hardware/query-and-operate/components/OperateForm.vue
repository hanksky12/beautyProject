<template>
    <div class="op">
    <base-slot-form :fields="fields" @submit="handleFormSubmit" >
      <template v-slot:title>
        ~@[操作]@~
      </template>
      <template v-for="field in fields"  v-slot:[field.slotName] >
            <component :is="field.component" :options_array="field.props.options_array" :rules="field.props.rules"></component>
      </template>
      <template v-slot:button>
        <execute-button />
        <clear-button />
      </template>
    </base-slot-form>
  </div>

</template>

<script>
import BaseSlotForm from '@/components/form/BaseSlotForm.vue';
import FunctionColumn from '@/components/form/column/FunctionColumn.vue';
import { createNamespacedHelpers } from 'vuex'
const { mapState:mapStateTable } = createNamespacedHelpers('table')
const { mapGetters:mapGetterForm } = createNamespacedHelpers('form')

import ClearButton from "@/components/form/button/ClearButton.vue";
import ExecuteButton from "@/components/form/button/ExecuteButton.vue";
import {GetOperateToSendRequestMixin} from "@/mixins/utils/get-operate-to-send-request/getOperateToSendRequest";
import {StartRecordApiMixin} from "@/mixins/api/hardware/startRecord";
import {StopRecordApiMixin} from "@/mixins/api/hardware/stopRecord";


export default {
  components: {
     'clear-button':ClearButton,
      'execute-button':ExecuteButton,
      'base-slot-form':BaseSlotForm,
      'function-operate-column':FunctionColumn,
  },
  mixins: [
    GetOperateToSendRequestMixin,
    StartRecordApiMixin,
    StopRecordApiMixin,
  ],
  data() {
      return {
          correspond_operate_obj :{
            '~@啟動@~': this.startRecordApi,
            '~@停止@~': this.stopRecordApi,
          },
      }
  },
  methods: {

    handleFormSubmit() {
        if (this.selected_array.length > 1 ) {
          this.showToast("只能選擇一筆!!!!", "warning");
          return
        }
        this.getOperateToSendRequest(
            this.getFormData,
            this.correspond_operate_obj,
            this.selected_array)
      }
  },
    computed:{
     ...mapStateTable({
        'selected_array': state => state.selected_array,
    }),
    ...mapGetterForm([
      'getFormData',
    ]),
    fields() {
     return [
      {
        slotName: 'function-operate-column',
        component: 'function-operate-column' ,
        props: {
            rules:"required",
            options_array:[{ text: '未选择', value: null }, ...Object.keys(this.correspond_operate_obj)]
        }
      },
      ]
    },
  },
    watch: {
        selected_array: function (newVal) {
            console.log('[Operate form] selected_array')
            console.log(`${JSON.stringify(newVal,null,2)}`)
        },
    }
}
</script>

<style scoped>
@import '@/components/style/op.css';
</style>