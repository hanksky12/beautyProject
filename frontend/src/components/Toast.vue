<template>
  <div>
    <b-toast
            id="b-toast-id"
            :title="title"
            :variant="variant"
            :solid="true"
            :append-toast="true"
            :autoHideDelay= "20000"
            toaster="b-toaster-bottom-left">
      <pre>{{ formattedMsg }}</pre>
    </b-toast>


  </div>
</template>

<script>
  export default {
      //提示訊息
      data() {
          return {
              title: "",
              msg: "",
              variant: "",
          }
      },
    methods: {
      show(content, level='info') {
       this.makeTitle(level)
       this.makeNodesMsg(content)
       this.makeVariant(level)
        this.$bvToast.show('b-toast-id')

      },

      makeTitle(level) {
        let title = null
        if (level  === "info"){title = '请求资讯'}
        else if (level === "success"){title = '请求成功'}
        else if (level === "danger"){title = '请求失败'}
        else if (level === "warning"){title = '请求警告'}
        this.title = title
      },
      makeNodesMsg(content) {
        this.msg = content
      },
      makeVariant(level){
        //info, success, warning, danger
          this.variant = (level === "info") ?'primary' :level
        }
    },
    computed: {
      formattedMsg() {
          return this.msg
      },
    },
  }
</script>