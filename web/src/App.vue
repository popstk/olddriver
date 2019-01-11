<template>
  <div id="app">
    <el-container>
      <el-header class="container">
        <div>
          <el-menu :default-active="$route.path" class="el-menu-demo" mode="horizontal" @select="handleSelect" router>
            <el-menu-item index="/">搜索</el-menu-item>
            <el-menu-item index="/Manage">管理</el-menu-item>
            <el-menu-item index="/Discovery">发现</el-menu-item>
          </el-menu>
        </div>
      </el-header>
      <router-view></router-view>
    </el-container>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath)
    }
  }
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
  width: 1140px;
  margin: 0 auto;
}
</style>
