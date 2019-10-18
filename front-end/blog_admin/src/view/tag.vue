<template>
  <div>
    <card>
      <div class="search-con search-con-top">
        <Button @click="addTag" class="search-btn" type="primary">
          <Icon type="search"/>&nbsp;&nbsp;新增标签
        </Button>
      </div>
      <Table :loading="loading" :data="tableData" :columns="tableColumns" stripe></Table>
      <div style="margin: 10px;overflow: hidden">
        <div style="float: right;">
          <Page :total="total" :current="page.pageNum" :page-size="page.pageSize" @on-change="changePage"
                show-total></Page>
        </div>
      </div>
    </card>
  </div>
</template>

<script>
    import {getTagList, addTag, updateTag, deleteTag} from '../api/tag'

    export default {
        name: 'tagPage',
        data() {
            return {
                tableData: [],
                loading: false,
                addLoading: false,
                total: 0,
                tagNma: '',
                page: {
                    pageSize: 10,
                    pageNum: 1,
                },
                tableColumns: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '标签名称',
                        key: 'name'
                    },
                    {
                        title: '创建时间',
                        key: 'created_at'
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            let data = params.row;
                                            this.updateTag(data.id, data.name)
                                        }
                                    }
                                }, '编辑'),
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: '你确定要删除吗?',
                                        transfer: true
                                    },
                                    on: {
                                        'on-ok': () => {
                                            let data = params.row;
                                            this.deleteTag(data.id)
                                        }
                                    }
                                }, [
                                    h('Button', {
                                        props: {
                                            type: 'error',
                                            size: 'small'
                                        }
                                    }, '删除')
                                ])
                            ]);
                        }
                    }
                ]
            }
        },
        methods: {
            //获取标签列表
            getTagList() {
                this.loading = true;
                let page = this.page;
                getTagList(page).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                    } else {
                        this.$Message.error(data.msg);
                    }
                })
            },
            changePage(page) {
                this.page.pageNum = page;
                this.getTagList()
            },
            addTag() {
                this.$Modal.confirm({
                    loading: this.addLoading,
                    render: (h) => {
                        return h('Input', {
                            props: {
                                autofocus: true,
                                placeholder: '请输入标签名称'
                            },
                            on: {
                                input: (val) => {
                                    this.tagName = val;
                                }
                            }
                        })
                    },
                    onOk: () => {
                        if (this.tagName === undefined || this.tagName == '') {
                            this.$Message.error('请输入标签名称!');
                            return;
                        }
                        this.addLoading = true;
                        addTag({name: this.tagName}).then(res => {
                            this.addLoading = false;
                            this.$Modal.remove();
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg);
                                this.getTagList();
                            } else {
                                this.$Message.error(data.msg);
                            }
                        })
                    }
                })
            },
            updateTag(id, name) {
                this.tagName = name;
                this.$Modal.confirm({
                    loading: this.addLoading,
                    render: (h) => {
                        return h('Input', {
                            props: {
                                value: name,
                                autofocus: true,
                                placeholder: '请输入标签名称'
                            },
                            on: {
                                input: (val) => {
                                    this.tagName = val;
                                }
                            }
                        })
                    },
                    onOk: () => {
                        if (this.tagName === undefined || this.tagName == '') {
                            this.$Message.error('请输入标签名称!');
                            return;
                        }
                        this.addLoading = true;
                        updateTag({id: id, name: this.tagName}).then(res => {
                            this.addLoading = false;
                            this.$Modal.remove();
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg);
                                this.getTagList();
                            } else {
                                this.$Message.error(data.msg);
                            }
                        })
                    }
                })
            },
            deleteTag(id) {
                this.loading = true;
                deleteTag({id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getTagList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created() {
            this.getTagList()
        }
    }
</script>

<style type="text/less" scoped>
  .search-con {
    padding: 10px 0;

    .search {
      &-col {
        display: inline-block;
        width: 200px;
      }

      &-input {
        display: inline-block;
        width: 200px;
        margin-left: 2px;
      }

      &-btn {
        margin-left: 2px;
      }
    }
  }
</style>
