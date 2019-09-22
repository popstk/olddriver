<template>
    <div>
        <el-input placeholder="请输入内容"
                  v-model="keyword"
                  class="input-with-select"
                  @keyup.enter.native="doSearch">
            <el-select v-model="select" slot="prepend" placeholder="请选择">
                <el-option label="琉璃神社" value="hacg"></el-option>
                <el-option label="桃花岛" value="taohua"></el-option>
            </el-select>
            <el-button slot="append" icon="el-icon-search" @click="doSearch"></el-button>
        </el-input>

        <el-table :data="tableData" v-loading="loading" height="500px">
            <el-table-column type="index"
                             width="80px"></el-table-column>
            <el-table-column label="标题">
                <template slot-scope="scope">
                    <el-link :href="scope.row.url"
                             type="primary"
                             target="_blank"
                             :underline="false"
                             rel="noreferrer">
                        {{scope.row.title}}
                    </el-link>
                </template>
            </el-table-column>
            <el-table-column label="站点" prop="tag" width="120px">
            </el-table-column>
            <el-table-column label="下载" width="300px">
                <template slot-scope="scope">
                    <el-col :span="12" v-if="scope.row.magnet">
                        <el-dropdown @command="handleMagnet">
                <span class="el-dropdown-link">
                    磁力<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item v-for="item in scope.row.magnet" :key="item" :command="item">{{item}}
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                    <el-col :span="12" v-if="scope.row.baidu">
                        <el-dropdown @command="handlePan">
                <span class="el-dropdown-link">
                    度盘<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item v-for="item in scope.row.baidu" :key="item" :command="item">{{item}}
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                    <el-col :span="12" v-if="scope.row.link">
                        <el-dropdown @command="handleLink">
                <span class="el-dropdown-link">
                    链接<i class="el-icon-arrow-down el-icon--right"></i>
                </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item v-for="item in scope.row.link" :key="item" :command="item">{{item}}
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                </template>
            </el-table-column>
        </el-table>
        <div style="margin-top: 10px">
            <el-pagination
                    background
                    @current-change="handleCurrentChange"
                    :current-page.sync="currentPage"
                    :hide-on-single-page="hidePagination"
                    :page-size="pageSize"
                    layout="prev, pager, next, jumper"
                    :total="total">
            </el-pagination>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'Search',
        data() {
            return {
                keyword: '',
                select: 'taohua',
                tableData: [],
                loading: false,
                total: 0,
                pageSize: 15,
                currentPage: 1,
                hidePagination: true,
            }
        },
        methods: {
            doSearch() {
                this.currentPage = 1
                this.handleCurrentChange()
            },
            handleCurrentChange() {
                const vm = this
                vm.loading = true
                this.$axios.post('/v1/spider/items', {
                    type: this.select,
                    keyword: this.keyword,
                    page: this.currentPage,
                    pageSize: this.pageSize
                }).then(res => {
                    vm.loading = false
                    vm.tableData = res.data.data
                    vm.hidePagination = ((vm.tableData === undefined) || vm.tableData.length === 0)

                    if (res.data.total) {
                        vm.total = res.data.total
                    } else {
                        vm.total = 1000
                    }
                }).catch(err => {
                    vm.$notify.error({
                        title: '错误',
                        message: err
                    });
                    vm.loading = false
                })
            },
            handleMagnet(text) {
                if (typeof text === 'object') {
                    text = text.length > 0 ? text[0] : ''
                }
                if (text) {
                    const prefix = 'magnet:?xt=urn:btih:';
                    if (!text.startsWith(prefix)) {
                        text = prefix + text
                    }
                    window.open(text)
                }
            },
            handlePan(text) {
                if (typeof text === 'object') {
                    text = text.length > 0 ? text[0] : ''
                }
                if (text) {
                    const prefix = 'http://pan.baidu.com/s/';
                    if (!text.startsWith(prefix)) {
                        text = prefix + text
                    }
                    window.open(text)
                }
            },
            handleLink(text) {
                window.open(text)
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
