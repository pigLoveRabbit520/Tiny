<template>
    <div class="app-container">
        <el-form ref="form" :model="form" :inline="true" label-width="130px">
            <el-form-item label="抓取文章时间点">
                <el-time-select
                        v-model="form.execTime"
                        :picker-options="{
                            start: '08:30',
                            step: '00:30',
                            end: '22:30'
                        }"
                        placeholder="选择时间">
                </el-time-select>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="updateSpiderExecTime">更新</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import { getData, updateExecTime } from '@/api/settings'

export default {
    data() {
        return {
            list: null,
            listLoading: true,
            form: {
                execTime: ''
            },
        }
    },
    created() {
        this.getSettings()
    },
    methods: {
        updateSpiderExecTime() {
            updateExecTime(this.form.execTime).then(response => {
                this.$message({
                    message: '更新执行时间成功',
                    type: 'success'
                });
            })
        },
        getSettings() {
            getData().then(response => {
                this.form = response.data
                this.previousExecTime= response.data.execTime
            })
        }
    },
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

