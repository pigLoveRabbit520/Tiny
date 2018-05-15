<template>
  <div class="dashboard-container">
      <el-row>
          <h2>爬虫执行日志</h2>
          <el-col :span="14">
              <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                        :data="logs"
                        style="width: 100%">
                  <el-table-column
                          type="index"
                          label="#"
                          width="180">
                  </el-table-column>
                  <el-table-column
                          prop="create_datetime"
                          label="执行时间"
                          width="200">
                  </el-table-column>
                  <el-table-column
                          prop="task_rate"
                          :formatter="getRate"
                          label="完成度">
                  </el-table-column>
                  <el-table-column
                          fixed="right"
                          label="操作"
                          width="100">
                      <template slot-scope="scope">
                          <el-button @click="viewDetail(scope.row.detail)" type="text" size="small">查看详情</el-button>
                      </template>
                  </el-table-column>
              </el-table>
          </el-col>
      </el-row>

      <el-row>
          <h2>推送日志</h2>
          <el-col :span="14">
              <el-table v-loading.body="listLoading" element-loading-text="Loading" border stripe
                        :data="pushLogs"
                        style="width: 100%">
                  <el-table-column
                          type="index"
                          label="#"
                          width="180">
                  </el-table-column>
                  <el-table-column
                          prop="push_datetime"
                          label="推送时间"
                          width="200">
                  </el-table-column>
                  <el-table-column
                          prop="title"
                          label="文章标题">
                  </el-table-column>
                  <el-table-column
                          prop="name"
                          label="公众号">
                  </el-table-column>
              </el-table>
          </el-col>
      </el-row>

      <el-dialog title="任务详情" :visible.sync="dialogTableVisible">
          <el-table :data="logDetail">
              <el-table-column property="name" label="公众号" width="200"></el-table-column>
              <el-table-column property="getHistoryUrl" label="获取历史页" width="200">
                  <template slot-scope="scope">
                      <i v-if="scope.row.getHistoryUrl" class="el-icon-check bg-success"></i>
                      <i v-else class="el-icon-close bg-danger"></i>
                  </template>
              </el-table-column>
              <el-table-column property="getMsgList" :formatter="getSuccessOrFail" label="获取历史页消息">
                  <template slot-scope="scope">
                      <i v-if="scope.row.getMsgList" class="el-icon-check bg-success"></i>
                      <i v-else class="el-icon-close bg-danger"></i>
                  </template>
              </el-table-column>
              <el-table-column property="addArticlesNum" label="抓取文章数量"></el-table-column>
          </el-table>
      </el-dialog>
  </div>
</template>

<script>
    import { mapGetters } from 'vuex'
    import { getList } from '@/api/dashboard'

    export default {
        data() {
            return {
                list: null,
                listLoading: true,
                logs: [],
                accountsMap: [],
                dialogTableVisible: false,
                logDetail: [],
                pushLogs: []
            }
        },
        created() {
            this.fetchData()
        },
        methods: {
            getRate(row, col) {
                return row[col.property] + "%";
            },
            getSuccessOrFail(row, col) {
                return row[col.property] ? '成功' : '失败';
            },
            fetchData() {
                this.listLoading = true
                getList().then(response => {
                    this.logs = response.data['logs']
                    this.pushLogs = response.data['pushLogs']
                    let accounts = response.data['accounts']
                    for (let account of accounts) {
                        this.accountsMap[account['account_id']] = account['name']
                    }
                    this.listLoading = false
                })
            },
            viewDetail(detail) {
                let logData = JSON.parse(detail);
                let data = [];
                if (Object.keys(logData).length > 0) {
                    for (let accountId in logData) {
                        let item = logData[accountId]
                        data.push({
                            name: this.accountsMap[accountId],
                            getHistoryUrl: item['getHistoryUrl'],
                            getMsgList: item['getMsgList'],
                            addArticlesNum: item['addArticles'] ? item['addArticles']['numReal'] : 0,
                        })
                    }
                    this.logDetail = data
                    this.dialogTableVisible = true
                }
            }
        }
    }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .dashboard {
    &-container {
      margin: 30px;
    }
  }
  .bg-success {
      color: #67c23a;
  }
  .bg-danger {
      color: #f56c6c;
  }
</style>
