<template>
    <el-table
            :data="mdata"
            v-loading="loading"
            stripe
            style="width: 100%">
        <el-table-column prop="spider" label="爬虫" width="230"></el-table-column>
        <el-table-column prop="start_time" :formatter="makeStartTime" label="开始时间" width="200"></el-table-column>
        <el-table-column :formatter="duration" label="耗时" width="120"></el-table-column>
        <el-table-column prop="status" label="状态" width="120"></el-table-column>
        <el-table-column label="操作">
            <template slot-scope="scope">
                <el-button v-if="scope.row.status != 'finished'" type="text" disabled>正在运行</el-button>
                <el-button v-else @click.native.prevent="startSpider(scope.row)" type="text">立即启动</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script>
    /* eslint-disable */
    export default {
        name: 'Manage',
        data() {
            return {
                loading: false,
                mdata: []
            }
        },
        methods: {
            startSpider(row) {
                const vm = this
                this.$axios
                    .post('/spider/' + row.spider)
                    .then(response => {
                        vm.refresh()
                        vm.$notify({
                            title: row.spider,
                            message: response.data,
                            type: 'success'
                        })
                    })
                    .catch(error => {
                        console.log(error)
                        vm.$notify.error({
                            title: '错误',
                            message: error
                        })
                    })
            },
            setJobsDetail(row) {
                console.log(row)
                this.gridData = row.jobs
            },
            refresh() {
                const vm = this
                vm.loading = true
                vm.$axios
                    .get('/spider/')
                    .then(response => {
                        vm.mdata = response.data
                        vm.loading = false
                    })
                    .catch(err => {
                        // eslint-disable-next-line no-console
                        console.log(err)
                        vm.loading = false
                    })
            },
            duration(row, column) {
                if (!row['end_time'] || row['status'] !== 'finished') {
                    return NaN
                }
                const sec = parseInt((new Date(row['end_time']) - new Date(row['start_time'])) / 1000.0);
                let hours = Math.floor(sec / 3600);
                let minutes = Math.floor((sec - hours * 3600) / 60);
                let seconds = sec - hours * 3600 - minutes * 60;
                if (hours < 10) hours = '0' + hours
                if (minutes < 10) minutes = '0' + minutes
                if (seconds < 10) seconds = '0' + seconds
                return hours + ':' + minutes + ':' + seconds
            },
            makeStartTime(row, column) {
                return new Date(row['start_time']).toLocaleString()
            }
        },
        created() {
            this.refresh()
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .item {
        margin-top: 10px;
        margin-right: 40px;
    }
</style>
