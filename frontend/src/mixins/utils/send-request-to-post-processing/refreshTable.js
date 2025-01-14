export const RefreshTableMixin = {
    methods: {
        refreshTable() {
          this.$store.commit("query/setQueryClick", true);
          console.log("refreshTable");
        },
    }
};