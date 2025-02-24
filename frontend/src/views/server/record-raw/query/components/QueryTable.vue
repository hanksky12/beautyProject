<template>
    <base-table
        :columns="columns"
        :api_path="api_path"
        :row_style="rowStyle"
    />

</template>

<script>
import BaseTable from "@/components/BaseTable.vue";

export default {
  components: {
    'base-table':BaseTable,
  },
  data () {
    return {
      api_path: "/server/record/raw-info",
      columns: [
                {
                  field: 'state',
                  checkbox: true,
                },
                {
                    field: 'id',
                    title: 'id',
                    visible: false
                },
                {
                  field: 'hardware_name',
                  title: '硬體名稱',
                  sortable: true
                },
                {
                  field: 'percent',
                  title: '運作率',
                  sortable: true
                },
                {
                  field: 'time',
                  title: '時間',
                  sortable: true
                },
      ],
      rowStyle() {
          return { classes: "bg-type-yellow" }
    },
    }
  },
}
</script>
<style>
.responsive-image {
  max-width: 70%; /* 圖片最大寬度為父容器寬度 */
  height: auto; /* 保持原始高度寬高比 */
}
</style>
