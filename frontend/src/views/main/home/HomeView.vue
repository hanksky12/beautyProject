<template>
    <div>
        <bread-crumb :props_items="items"/>
        <query-table />
        <br>
    </div>
</template>

<script>
import Breadcrumb from '@/components/Breadcrumb.vue'
import QueryTable from './components/QueryTable.vue';

export default {
    components: {
        'bread-crumb':Breadcrumb,
      'query-table':QueryTable,
    },
    data() {
      return {
        items: [
          {
            text: '首頁',
             active: true
          },
        ]
      }
    }
}
</script>

