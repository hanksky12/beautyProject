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
      api_path: "/user/record/raw-info",
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
                  field: 'mouse_action_name',
                  title: '滑鼠操作',
                  sortable: true
                },
                {
                  field: 'x',
                  title: 'X座標',
                  sortable: true
                },
                {
                  field: 'y',
                  title: 'Y座標',
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
