<template>
    <div>
        <el-input @keyup.enter.native="doSearch"
                  class="input-with-select"
                  placeholder="请输入内容"
                  v-model="keyword">
            <el-select placeholder="请选择" slot="prepend" v-model="select">
                <el-option label="琉璃神社" value="hacg"></el-option>
                <el-option label="桃花岛" value="taohua"></el-option>
            </el-select>
            <el-button @click="doSearch" icon="el-icon-search" slot="append"></el-button>
        </el-input>

        <el-table :data="tableData" height="500px" v-loading="loading">
            <el-table-column type="index"
                             width="80px"></el-table-column>
            <el-table-column label="标题">
                <template slot-scope="scope">
                    <el-link :href="scope.row.url"
                             :underline="false"
                             rel="noreferrer"
                             target="_blank"
                             type="primary">
                        {{scope.row.title}}
                    </el-link>
                </template>
            </el-table-column>
            <el-table-column label="发布时间"
                             prop="time"
                             :formatter="formatDayOnly"
                             width="120px"
                             sortable>
            </el-table-column>
            <el-table-column label="来源"
                             prop="tag"
                             width="120px">
            </el-table-column>
            <el-table-column label="下载" width="300px">
                <template slot-scope="scope">
                    <el-col :span="12" v-if="scope.row.magnet">
                        <el-dropdown @command="handleMagnet">
                            <span class="el-dropdown-link">
                                磁力<i class="el-icon-arrow-down el-icon--right"></i>
                            </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item :command="item"
                                                  :key="item"
                                                  v-for="item in scope.row.magnet">
                                    {{item}}</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                    <el-col :span="12" v-if="scope.row.baidu">
                        <el-dropdown @command="handlePan">
                            <span class="el-dropdown-link">
                                度盘<i class="el-icon-arrow-down el-icon--right"></i>
                            </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item
                                        :command="item"
                                        :key="item"
                                        v-for="item in scope.row.baidu">
                                    {{item}}</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                    <el-col :span="12" v-if="scope.row.link">
                        <el-dropdown @command="handleLink">
                            <span class="el-dropdown-link">
                                链接<i class="el-icon-arrow-down el-icon--right"></i>
                            </span>
                            <el-dropdown-menu slot="dropdown">
                                <el-dropdown-item :command="item"
                                                  :key="item" v-for="item in scope.row.link">
                                    {{item}}</el-dropdown-item>
                            </el-dropdown-menu>
                        </el-dropdown>
                    </el-col>
                </template>
            </el-table-column>
        </el-table>
        <div style="margin-top: 10px">
            <el-pagination
                    :current-page.sync="currentPage"
                    :hide-on-single-page="hidePagination"
                    :page-size="pageSize"
                    :total="total"
                    @current-change="handleCurrentChange"
                    background
                    layout="prev, pager, next, jumper">
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
        mounted: function (){
            Number.prototype.padLeft = function (base, chr) {
                const len = (String(base || 10).length - String(this).length) + 1;
                return len > 0 ? new Array(len).join(chr || '0') + this : this;
            };
        },
        methods: {
            doSearch() {
                this.currentPage = 1;
                this.handleCurrentChange()
            },
            formatDayOnly(row){
                let d = new Date(row.time);
                const dd = d.getDate().padLeft();
                const mm = (d.getMonth() + 1).padLeft();
                const yyyy = d.getFullYear();
                return yyyy + '-' + mm + '-' + dd
            },
            handleCurrentChange() {
                const vm = this;
                vm.loading = true;
                this.$axios.post('/v1/spider/items', {
                    type: this.select,
                    keyword: this.keyword,
                    page: this.currentPage,
                    pageSize: this.pageSize
                }).then(res => {
                    vm.loading = false;
                    vm.tableData = res.data.data;
                    vm.hidePagination = ((vm.tableData === undefined) || vm.tableData.length === 0);

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
</style>
