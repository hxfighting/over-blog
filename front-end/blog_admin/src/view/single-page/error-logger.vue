<template>
    <div>
        <card>
            <Button @click="exportData" type="primary" style="margin: 0 10px 10px 0;">导出日志记录</Button>
            <Button @click="deleteErrorLog" v-if="deleteButtonShow" type="error" style="margin: 0 10px 10px 0;">删除
            </Button>
            <Table border
                   ref="table"
                   :columns="columns"
                   :data="tableData"
                   :loading="loading"
                   @on-selection-change="tableSelectChange"
            ></Table>
            <div style="margin: 10px;overflow: hidden">
                <div style="float: right;">
                    <Page :total="total" show-total :page-size="listData.pageSize" @on-change="changePage"></Page>
                </div>
            </div>
        </card>
    </div>
</template>

<script>
    import dayjs from 'dayjs'
    import {mapMutations} from 'vuex'
    import {getErrorList, deleteErrorLog} from '../../api/data'

    export default {
        name: 'error_logger_page',
        data() {
            return {
                listData: {
                    pageSize: 10,
                    pageNum: 1
                },
                deleteButtonShow: false,
                flag: false,
                loading: false,
                tableData: [],
                ids: [],
                total: 0,
                columns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        type: 'index',
                        title: '序号',
                        width: 100,
                    },
                    {
                        key: 'type',
                        title: '类型',
                        width: 100,
                        render: (h, {row}) => {
                            return (
                                < div >
                                < icon
                            size = {16}
                            type = {row.type === 'ajax' ? 'md-link' : 'md-code-working'} > < /icon>
                                < /div>
                        )
                        }
                    },
                    {
                        key: 'code',
                        title: '编码',
                        render: (h, {row}) => {
                            return (
                                < span > {row.code === 0 ? '-' : row.code} < /span>
                        )
                        }
                    },
                    {
                        key: 'mes',
                        title: '信息'
                    },
                    {
                        key: 'url',
                        title: 'URL'
                    },
                    {
                        key: 'time',
                        title: '时间',
                        render: (h, {row}) => {
                            if (this.flag) {
                                return h('span', {}, row.created_at);
                            } else {
                                return (
                                    < span > {dayjs(row.time
                            ).
                                format('YYYY-MM-DD HH:mm:ss')
                            }<
                                /span>
                            )
                            }
                        },
                        sortable: true
                    }
                ]
            }
        },
        computed: {
            errorList() {
                this.getErrorList()
            }
        },
        methods: {
            ...mapMutations([
                'setHasReadErrorLoggerStatus'
            ]),
            exportData() {
                this.$refs.table.exportCsv({
                    filename: '错误日志.csv'
                })
            },
            changePage(page) {
                this.listData.pageNum = page;
                this.getErrorList();
            },
            getErrorList() {
                this.loading = true;
                getErrorList(this.listData).then(res => {
                    this.loading = false;
                    this.flag = true;
                    let data = res.data;
                    if (data.code == 200) {
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                    } else {
                        this.tableData = [];
                        this.total = 0;
                        this.$Message.error(data.msg);
                    }
                }).catch(err => {
                    this.tableData = this.$store.state.app.errorList;
                })
            },
            tableSelectChange(val) {
                this.ids = [];
                for (let va of val) {
                    this.ids.push(va.id)
                }
                this.deleteButtonShow = val.length > 0;
            },
            deleteErrorLog() {
                this.$Modal.confirm({
                    title: '删除错误日志',
                    content: '<p>你确定删除这些日志吗?</p>',
                    loading: true,
                    onOk: () => {
                        deleteErrorLog({ids: this.ids}).then(res => {
                            this.$Modal.remove();
                            this.deleteButtonShow = false;
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg);
                                this.getErrorList();
                            } else {
                                this.$Message.error(data.msg);
                            }
                        }).catch(err => {
                            this.$Modal.remove();
                            this.deleteButtonShow = false;
                            this.tableData = [];
                            this.total = 0;
                            this.$Message.error('服务器错误');
                        })
                    }
                });
            }
        },
        activated() {
            this.setHasReadErrorLoggerStatus()
            this.getErrorList();
        },
        mounted() {
            this.setHasReadErrorLoggerStatus()
        }
    }
</script>

<style>

</style>
