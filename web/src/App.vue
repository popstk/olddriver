<template>
  <div id="app">
    <el-container class="container">
      <el-header>
          <el-menu :default-active="$route.path"
                   class="el-menu-demo"
                   mode="horizontal"
                   @select="handleSelect" router
                   background-color="#545c64"
                   text-color="#fff"
                   active-text-color="#ffd04b">
            <el-menu-item index="/">搜索</el-menu-item>
            <el-menu-item index="/Manage">管理</el-menu-item>
            <el-menu-item index="/Discovery">发现</el-menu-item>
          </el-menu>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </div>
</template>

<script>
  import Router from "./router/index"
  export default {
    name: 'App',
    data() {
      return {
      }
    },
    methods: {
      handleSelect(key, keyPath) {
        // eslint-disable-next-line no-console
        console.log(key, keyPath)
      }
    },
    router: Router
  }
</script>

<style>
  #app {
    font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB',
    'Microsoft YaHei', SimSun, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  .container {
    width: 85%;
    margin: 0 auto;
  }
</style>
