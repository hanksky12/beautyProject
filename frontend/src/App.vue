<template>
  <div id="app">
      <toast ref="toast"/>
      <confirm-modal ref="confirmModal"/>
      <navbar />
      <router-view/>
  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import Toast from "@/components/Toast.vue";
import ConfirmModal from "@/components/modal/ConfirmModal.vue";
import Swal from "sweetalert2";
export default {
  components: {
    'confirm-modal':ConfirmModal,
    'toast':Toast,
    'navbar':Navbar,
  },
  provide() {
      return{
          showToast: this.showToast,
          showConfirmModal: this.showConfirmModal,
          showSwAl: this.showSwAl,
      }
  },

  methods: {
    showToast(content, level) {
      this.$refs.toast.show(content, level)
    },
    showConfirmModal(title, content, op_func, op_data) {
      this.$refs.confirmModal.show(title, content, op_func, op_data)
    },
    showSwAl(title, content) {
         Swal.fire({
          title: title,
          html: content,
          confirmButtonText: '關閉',
          width: 600,
          position: 'center-end',
        });
    }
  },
};
</script>


<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #7792af;
}

nav {
  padding: 50px;
}

nav a {
  font-weight: bold;
  color: #e5998f;
}

nav a.router-link-exact-active {
  color: #bdc0e3;
}

.bg-type-yellow {
    background-color: #FFFFBB !important;
}

.bg-type-purple {
    background-color: #CCCCFF !important;
}

.bg-type-red {
    background-color: #FFB7DD !important;
}

.bg-type-green {
    background-color: #00FF99 !important;
}

.bg-type-bronze {
    background-color: #e5a469 !important;
}

.bg-type6 {
    background-color: #FFFFFF !important;
}

.result-table {
  text-align: center;
  padding: 10px 0;
  background-color: #ffffff;
  color: #d96767;
}
</style>

