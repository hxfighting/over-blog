<template>
    <div>
        <card>
            <div class="search-con search-con-top">
                <Input @on-change="handleClear" clearable placeholder="输入友联名称搜索" class="hx_input"
                       v-model="listData.name"/>
                <Button @click="handleSearch" class="search-btn" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;搜索
                </Button>
                <Button @click="modal_add=true" class="link_add_button" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;新增友联
                </Button>
            </div>
            <Table :data="tableData" :columns="tableColumns1" stripe :loading="loading"></Table>
            <div style="margin: 10px;overflow: hidden">
                <div style="float: right;">
                    <Page :total="total" show-total :page-size="listData.pageSize" @on-change="changePage"></Page>
                </div>
            </div>
            <Modal v-model="modal_add" scrollable :mask-closable="false" :closable="false" title="新增友联" @on-ok="addLink('formCustom')"
                   :loading="button_loading" @on-cancel="modalCancel('formCustom')">
                <Form ref="formCustom" :model="formData" :rules="ruleValidate" :label-width="80">
                    <FormItem label="友联URL" prop="url">
                        <Input v-model="formData.url" placeholder="请输入友联URL"></Input>
                    </FormItem>
                    <FormItem label="友联名称" prop="name">
                        <Input v-model="formData.name" placeholder="Enter something..."></Input>
                    </FormItem>
                    <FormItem label="友联描述" prop="description">
                        <Input v-model="formData.description" type="textarea" :rows="4" placeholder="Enter something..."></Input>
                    </FormItem>
                    <FormItem label="友联排序" prop="order">
                        <Input v-model="formData.order" placeholder="Enter something..."></Input>
                    </FormItem>
                    <FormItem label="是否显示">
                        <i-switch v-model="is_show_bool" size="large" @on-change="handleSwitch">
                            <span slot="open">On</span>
                            <span slot="close">Off</span>
                        </i-switch>
                    </FormItem>
                </Form>
            </Modal>
            <Modal v-model="modal_update" scrollable :mask-closable="false" :closable="false" title="编辑友联" @on-ok="updateLink('formCustomTwo')"
                   :loading="button_loading" @on-cancel="modalCancel('formCustomTwo')">
                <Form ref="formCustomTwo" :model="formData" :rules="ruleValidate" :label-width="80">
                    <FormItem label="友联URL" prop="url">
                        <Input v-model="formData.url" placeholder="请输入友联URL"></Input>
                    </FormItem>
                    <FormItem label="友联名称" prop="name">
                        <Input v-model="formData.name" placeholder="Enter something..."></Input>
                    </FormItem>
                    <FormItem label="友联描述" prop="description">
                        <Input v-model="formData.description" type="textarea" :rows="4" placeholder="Enter something..."></Input>
                    </FormItem>
                    <FormItem label="友联排序" prop="order">
                        <Input v-model="formData.order" placeholder="Enter something..."></Input>
                    </FormItem>
                    <FormItem label="是否显示">
                        <i-switch v-model="is_show_bool" size="large" @on-change="handleSwitch">
                            <span slot="open">On</span>
                            <span slot="close">Off</span>
                        </i-switch>
                    </FormItem>
                </Form>
            </Modal>
        </card>
    </div>
</template>

<script>
    import {getLinkList, addLink, deleteLink, updateLink} from '../api/link'

    export default {
        name: 'linkPage',
        data() {
            const validateOrderCheck = (rule, value, callback) => {
                if (value === '' || isNaN(value) || !(value%1 === 0)) {
                    callback(new Error('请输入数字排序,排序只能是整数!'));
                }else {
                    callback();
                }
            };
            return {
                tableData: [],
                searchValue: '',
                loading: false,
                total: 0,
                modal_add: false,
                modal_update:false,
                button_loading: true,
                listData: {
                    name: '',
                    pageSize: 10,
                    pageNum: 1,
                },
                is_show_bool: true,
                formData: {
                    name: '',
                    url: '',
                    order: '',
                    is_show: 1,
                    description: '',
                    id:''
                },
                tableColumns1: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '友联排序',
                        key: 'order'
                    },
                    {
                        title: '友联URL',
                        key: 'url',
                        render: (h, params) => {
                            const row = params.row;
                            return h('a', {
                                attrs: {
                                    title: row.name,
                                    href: row.url,
                                    target: '__blank'
                                },
                            }, row.url);
                        }
                    },
                    {
                        title: '友联名称',
                        key: 'name'
                    },
                    {
                        title: '友联描述',
                        key: 'description'
                    },
                    {
                        title: '是否显示',
                        key: 'is_show',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.is_show === 1 ? 'success' : 'error';
                            const text = row.is_show === 1 ? '√' : '×';
                            const title = row.is_show === 1 ? '显示' : '不显示';

                            return h('Tag', {
                                props: {
                                    color: color
                                },
                                attrs: {
                                    title: title,
                                },
                            }, text);
                        }
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
                                    attrs: {
                                        title: '编辑',
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.updateLinkModal(params.row)
                                        }
                                    }
                                }, '编辑'),
                                h('Poptip', {
                                    props: {
                                        confirm: true,
                                        title: '你确定要删除吗?'
                                    },
                                    on: {
                                        'on-ok': () => {
                                            let data = params.row;
                                            this.deleteLink(data.id)
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
                ],
                ruleValidate: {
                    url: [
                        {required: true, message: '请输入友联URL', trigger: 'blur'},
                        {type: 'url', message: '请输入正确的URL地址', trigger: 'change'},
                    ],
                    name: [
                        {required: true, message: '请输入友联名称', trigger: 'blur'},
                        {type: 'string', min: 2, max: 30, message: '友联名称在2到30个字符之间', trigger: 'change'},
                    ],
                    description: [
                        {required: true, message: '请输入友联描述', trigger: 'blur'},
                        {type: 'string', min: 2, max: 50, message: '友联描述在2到50个字符之间', trigger: 'change'}
                    ],
                    order: [
                        {validator: validateOrderCheck, trigger: 'change'}
                    ]
                }
            }
        },
        methods: {
            handleSwitch(value){
                this.formData.is_show = value?1:0
            },
            modalCancel(name) {
                this.$refs[name].resetFields();
            },
            changePage(page) {
                this.listData.pageNum = page;
                this.getLinkList();
            },
            handleClear(e) {
                if (e.target.value === '') this.insideTableData = this.value
            },
            handleSearch() {
                this.listData.pageNum = 1;
                this.getLinkList();
            },
            //处理button的loading状态
            dealButtonLoading(){
                let _this = this
                setTimeout(function () {
                    _this.button_loading = false;
                    _this.$nextTick(() => {_this.button_loading = true;});
                }, 500);
            },
            updateLinkModal(row){
                this.is_show_bool = row.is_show==1?true:false;
                this.formData.is_show = row.is_show;
                this.formData.name = row.name;
                this.formData.url = row.url;
                this.formData.description = row.description;
                this.formData.order = row.order;
                this.formData.id = row.id;
                this.modal_update = true;
            },
            getLinkList() {
                this.loading = true;
                getLinkList(this.listData).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                    } else {
                        this.$Message.error(data.msg);
                    }
                })
            },
            addLink(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        addLink(this.formData).then(res=>{
                            let data = res.data;
                            if (data.code == 200) {
                                this.modal_add = false;
                                this.modalCancel('formCustom');
                                this.$Message.success(data.msg);
                                this.getLinkList();
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg);
                            }
                        });
                    } else {
                        this.dealButtonLoading();
                    }
                })
            },
            updateLink(name){
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        updateLink(this.formData).then(res=>{
                            let data = res.data;
                            if (data.code == 200) {
                                this.modal_update = false;
                                this.modalCancel('formCustom');
                                this.$Message.success(data.msg);
                                this.getLinkList();
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg);
                            }
                        });
                    } else {
                        this.dealButtonLoading();
                    }
                })
            },
            deleteLink(id){
                this.loading = true;
                deleteLink({id}).then(res=>{
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getLinkList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created() {
            this.getLinkList()
        }
    }
</script>

<style scoped>
    .hx_input {
        display: inline-block;
        width: 200px !important;
        margin-left: 2px;
        margin-right: 3px;
    }

    .search-con {
        padding: 10px 0;
    }

    .link_add_button{
        float: right;
    }
</style>
