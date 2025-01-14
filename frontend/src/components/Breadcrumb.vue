<template>
    <div>
     <b-breadcrumb :items="items" class="bread"></b-breadcrumb>
    </div>

</template>

<script>
export default {
    props: {
        props_items: {
            type: Array,
            required: true,
        }
    },

    computed: {
        items() {
            const additionalItem = {
                text: '目前位置==>',
                active: true
          }
            return [additionalItem, ...this.props_items];
        }
    },
}
</script>

<style scoped>
.bread {
  padding: 10px 0;
    font-size: 15px;
    font-weight: 600;
  background-color: rgba(119, 146, 175, 0.85);
}

</style>