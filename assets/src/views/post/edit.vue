<template>
    <div class="app-container">
        <textarea rows="10"></textarea>
    </div>
</template>

<script>
import { getList } from '@/api/post'

export default {
    data() {
        return {
            listLoading: true,
            postsTotal: 0,
            pageSize: 15,
            currentPage: 1,
            posts: [],
            cates: [],
        }
    },
    created() {
        document.title = "编辑文章"
    },
    methods: {
    }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

