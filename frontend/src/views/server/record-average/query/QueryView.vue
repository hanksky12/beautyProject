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
            text: '資料查詢',
            active: true
          },
          {
            text: '伺服器',
            active: true
          },
          {
            text: '平均查詢',
            active: true
          }
        ]
      }
    }
}
</script>

<style scoped>

</style>