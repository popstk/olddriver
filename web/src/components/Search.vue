<template>
<el-main class="container">
    <el-row>
        <el-col :span="24">
        <el-input placeholder="请输入内容" v-model="keyword" class="input-with-select" @keyup.enter.native="dosearch">
            <el-select v-model="select" slot="prepend" placeholder="请选择">
            <el-option label="琉璃神社" value="hacg"></el-option>
            <el-option label="桃花岛" value="taohua"></el-option>
            </el-select>
            <el-button slot="append" icon="el-icon-search" @click="dosearch"></el-button>
        </el-input>
        </el-col>
    </el-row>

    <el-row>
        <el-table :data="tableData" v-loading="loading">
        <el-table-column label="标题">
            <template slot-scope="props">
            <a :href="props.row.url" target="_blank">{{props.row.title}}</a>
            </template>
        </el-table-column>
        <el-table-column label="站点" prop="site" width="120px">
        </el-table-column>
        <el-table-column label="下载" width="150px">
            <template slot-scope="scope">
            <el-col :span="12" v-if="scope.row.magnets.length>0">
                <el-dropdown @command="handleMagnet">
                <span class="el-dropdown-link">
                    磁力
                    <i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item v-for="item in scope.row.magnets" :key="item" :command="item">{{item}}</el-dropdown-item>
                </el-dropdown-menu>
                </el-dropdown>
            </el-col>
            <el-col :span="12" v-if="scope.row.baidu.length>0">
                <el-dropdown @command="handlePan">
                <span class="el-dropdown-link">
                    度盘
                    <i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item v-for="item in scope.row.baidu" :key="item" :command="item">{{item}}</el-dropdown-item>
                </el-dropdown-menu>
                </el-dropdown>
            </el-col>
            </template>
        </el-table-column>
        </el-table>
    </el-row>
    </el-main>
</template>

<script>
import {SpiderClient} from 'backend_grpc_web_pb'
import {SearchRequest, SearchReply} from 'backend_pb'

export default {
  name: 'Search',
  data() {
    return {
      keyword: '',
      select: 'hacg',
      tableData: [],
      loading: false
    }
  },
  created: function() {
    this.client = new SpiderClient('/grpc')
  },
  methods: {
    dosearch() {
      const vm = this
      vm.loading = true
      
      var request = new SearchRequest()
      request.setType = this.select
      request.setKeyword = this.keyword

      var call = this.client.search(request,{}, function(err, response) {
        if (err) {
          vm.$notify.error({
            title: '错误',
            message: err
          })
          console.log("normal is ", err)
          vm.loading = false
          return
        }

        vm.tableData = response.data
        vm.loading = false
      })

      call.on('status', function(status) {
        console.log("status code is ", status.code)
        console.log("status details is ", status.details)
        console.log("status metadata is ", status.metadata)
      })
    },
    handleMagnet(text) {
      if (typeof text === 'object') {
        text = text.length > 0 ? text[0] : ''
      }
      if (text) {
        const prefix = 'magnet:?xt=urn:btih:'
        if (!text.startsWith(prefix)) {
          text = prefix + text
        }
        console.log(text)
        window.open(text)
      }
    },
    handlePan(text) {
      if (typeof text === 'object') {
        text = text.length > 0 ? text[0] : ''
      }
      if (text) {
        const prefix = 'http://pan.baidu.com/s/'
        if (!text.startsWith(prefix)) {
          text = prefix + text
        }
        console.log(text)
        window.open(text)
      }
    }
  }
}
</script>

<style scoped>
.el-dropdown-link {
  cursor: pointer;
  color: #409eff;
}

.el-icon-arrow-down {
  font-size: 12px;
}

.el-select {
  width: 130px;
}
</style>
