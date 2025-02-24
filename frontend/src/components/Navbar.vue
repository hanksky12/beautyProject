<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="primary">
<!--        primary, success, info, warning, danger, dark, light-->
      <b-navbar-brand><router-link to="/">[美麗後台]</router-link></b-navbar-brand>
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item-dropdown text="[硬體操作]" right>
              <b-dropdown-item disabled>伺服器硬體</b-dropdown-item>
              <b-dropdown-item ><router-link to="/server-hardware-query-and-operate">操作</router-link></b-dropdown-item>
              <b-dropdown-divider />
              <b-dropdown-item disabled>使用者滑鼠</b-dropdown-item>
              <b-dropdown-item ><router-link to="/user-mouse-action-query-and-operate">操作</router-link></b-dropdown-item>
          </b-nav-item-dropdown>
          <b-nav-item-dropdown text="[資料查詢]" right>
            <b-dropdown-item disabled>伺服器查詢</b-dropdown-item>
            <b-dropdown-item ><router-link to="/server-average-record-query">[平均查詢]</router-link></b-dropdown-item>
            <b-dropdown-item ><router-link to="/server-raw-record-query">[原始查詢]</router-link></b-dropdown-item>
            <b-dropdown-divider />
            <b-dropdown-item disabled>使用者查詢</b-dropdown-item>
            <b-dropdown-item ><router-link to="/user-raw-record-query">[原始查詢]</router-link></b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>

        <!-- Right aligned nav items -->
        <b-navbar-nav class="ml-auto">
          <b-nav-item-dropdown right>
            <template v-slot:button-content>
              <em>帳號</em>
            </template>
            <b-dropdown-item v-if="!isAuthenticated"><router-link to="/login">登入</router-link></b-dropdown-item>
            <b-dropdown-item v-if="isAuthenticated"  @click="logout">登出</b-dropdown-item>
            <b-dropdown-item ><router-link to="/personal-register">注册</router-link></b-dropdown-item>
          </b-nav-item-dropdown>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
  </div>
</template>
<script>
import {LogOutApiMixin} from "@/mixins/api/user/authentication/logOut";

export default {
  mixins: [LogOutApiMixin],
  computed: {
    isAuthenticated() {
        return this.$store.state.authenticate.isAuthenticated;
    },
    isDevelopmentOrRoot() {
        // return  this.$store.state.authorization.level === 'root';
        return process.env.NODE_ENV === 'development' || this.$store.state.authorization.level === 'root';
    },
  },
  methods: {
    async logout() {
      await this.logOutApi()
      this.$store.commit("authenticate/setAuthenticated", false);
      this.$store.commit("authorization/setAuthorization", "");
      this.$router.push('/login'); // 跳转到登录页面
    },
  },
};
</script>


