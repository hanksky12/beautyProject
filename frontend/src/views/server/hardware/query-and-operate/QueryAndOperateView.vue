<template>
    <div>
        <bread-crumb :props_items="items"/>
        <operate-form />
      <query-table />

    </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue'
import OperateForm from "./components/OperateForm.vue";
import QueryTable from './components/QueryTable.vue';


export default {
    components: {
            'bread-crumb':Breadcrumb,
            'operate-form':OperateForm,
            'query-table':QueryTable,

    },
    data() {
      return {
        items: [
          {
            text: '硬體操作',
             active: true
          },
          {
            text: '伺服器',
            active: true
          },
          {
            text: '操作',
            active: true
          }
        ]
      }
    }
}
</script>

<style scoped>

</style>