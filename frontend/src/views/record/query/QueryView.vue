<template>
    <div>
        <bread-crumb :props_items="items"/>
        <query-form />
        <query-table />
    </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue'
import QueryForm from './components/QueryForm.vue'
import QueryTable from './components/QueryTable.vue';

export default {
    components: {
             'bread-crumb':Breadcrumb,
            'query-form':QueryForm,
            'query-table':QueryTable,
    },
    data() {
      return {
        items: [
          {
            text: '紀錄操作',
             active: true
          },
          {
            text: '資料查詢',
            active: true
          },
          {
            text: '查詢',
            active: true
          }
        ]
      }
    }
}
</script>

<style scoped>

</style>