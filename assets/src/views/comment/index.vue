<template>
    <div class="app-container">

        <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                :data="comments"
                style="width: 60%">
            <el-table-column
                    type="selection"
                    width="33">
            </el-table-column>
            <el-table-column
                    prop="author"
                    label="作者">
                <template slot-scope="scope">
                    <el-button type="text" style="font-weight: bold;">{{ scope.row.author }}</el-button><br>
                    <a :href="'mailto:' + scope.row.mail" target="_blank">{{ scope.row.mail }}</a><br>
                    <span style="color: #999;">{{ scope.row.ip }}</span>
                </template>
            </el-table-column>
            <el-table-column
                    prop="text"
                    label="内容">
                <template slot-scope="scope">
                    <div class="comment-date">
                        {{ scope.row.created | formatTime }} 于 <el-button type="text">{{ scope.row.title }}</el-button><br>
                    </div>
                    <span style="color: #999;">{{ scope.row.text }}</span>
                </template>
            </el-table-column>
        </el-table>

        <div style="margin-top: 10px; margin-left: -10px">
            <el-pagination
                    background
                    layout="prev, pager, next"
                    @current-change="handleCurrentChange"
                    :page-size="pageSize"
                    :current-page.sync="currentPage"
                    :total="commentsTotal">
            </el-pagination>
        </div>
    </div>
</template>

<script>
import { getList } from '@/api/comment'
import { formatTime, parseTime } from '@/utils/index'

export default {
    data() {
        return {
            listLoading: true,
            commentsTotal: 0,
            pageSize: 15,
            currentPage: 1,
            comments: [],
        }
    },
    created() {
        document.title = "评论管理"
        this.getComments()
    },
    methods: {
        getComments(params) {
            this.listLoading = true
            getList(params).then(response => {
                this.comments = response.data.comments
                this.commentsTotal = response.data.total
                this.listLoading = false
            }).then(() => {
                this.listLoading = false
            })
        },
        handleCurrentChange(val) {
            this.getComments({
                page: val
            })
        }
    },
    filters: {
        formatTime: formatTime
    }
}
</script>

<style scoped>
.line{
    text-align: center;
}
.comment-date {
    font-size: .92857em;
    color: #999;
}
</style>

