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
      api_path: "/server/hardware-specs/info",
      columns: [
                {
                  field: 'cpu_info',
                  title: 'CPU規格',
                },
                {
                  field: 'memory_info',
                  title: '記憶體規格',
                },
                {
                  field: 'disk_info',
                  title: '硬碟規格',
                },
                {
                  field: 'host_info',
                  title: '主機資訊',
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
