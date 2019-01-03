<template>
    <card>
        <div class="search-con search-con-top">
            <Select v-model="listData.type" style="width:150px;padding-right: 5px" clearable filterable
                    placeholder="请选择配置类型">
                <Option v-for="item in type" :value="item.id" :key="item.id">{{ item.title }}</Option>
            </Select>
            <Button @click="handleSearch" class="search-btn" type="primary">
                <Icon type="search"/>&nbsp;&nbsp;搜索
            </Button>
            <Button @click="configModal=true" class="link_add_button" type="primary">
                <Icon type="search"/>&nbsp;&nbsp;新增配置
            </Button>
        </div>
        <Table :data="tableData" :columns="tableColumns" :loading="loading" stripe></Table>
        <div style="margin: 10px;overflow: hidden">
            <div style="float: right;">
                <Page :total="total" show-total :page-size="listData.pageSize" @on-change="changePage"></Page>
            </div>
        </div>
        <Modal v-model="configModal"
               :mask-closable="false"
               :closable="false"
               title="新增配置"
               @on-ok="handleConfig('configAdd')"
               :loading="button_loading"
               @on-cancel="modalCancel('configAdd')"
        >
            <Form ref="configAdd" :model="formData" :rules="ruleValidate" :label-width="80">
                <FormItem label="配置类型" prop="type">
                    <RadioGroup v-model="formData.type">
                        <Radio label="1">社交地址</Radio>
                        <Radio label="2">footer内容</Radio>
                        <Radio label="3">其他配置</Radio>
                    </RadioGroup>
                </FormItem>
                <FormItem label="配置标题" prop="title">
                    <Input v-model="formData.title" placeholder="请输入配置标题"></Input>
                </FormItem>
                <FormItem label="配置key" prop="name">
                    <RadioGroup v-model="formData.name" v-if="formData.type=='1'">
                        <Radio label="qq">QQ</Radio>
                        <Radio label="twitter">twitter</Radio>
                        <Radio label="weibo">微博</Radio>
                        <Radio label="git">码云</Radio>
                        <Radio label="github">github</Radio>
                    </RadioGroup>
                    <RadioGroup v-model="formData.name" v-else-if="formData.type=='2'">
                        <Radio label="famous_remark">名言</Radio>
                        <Radio label="blog_related">博客相关</Radio>
                        <Radio label="blogger_collection">博主常逛</Radio>
                        <Radio label="copyright">版权信息</Radio>
                    </RadioGroup>
                    <Input v-model="formData.name" placeholder="请输入配置key" v-else></Input>
                </FormItem>
                <FormItem label="配置value" prop="val">
                    <Input v-model="formData.val" type="textarea" :rows="6"
                           placeholder="请输入配置value"></Input>
                </FormItem>
            </Form>
        </Modal>
        <Modal v-model="configUpdateModal"
               :mask-closable="false"
               :closable="false"
               title="修改配置"
               @on-ok="updateConfig('configUpdate')"
               :loading="button_loading"
               @on-cancel="modalCancel('configUpdate')"
        >
            <Form ref="configUpdate" :model="formData" :rules="ruleValidate" :label-width="80">
                <FormItem label="配置类型" prop="type">
                    <RadioGroup v-model="formData.type">
                        <Radio label="1">社交地址</Radio>
                        <Radio label="2">footer内容</Radio>
                        <Radio label="3">其他配置</Radio>
                    </RadioGroup>
                </FormItem>
                <FormItem label="配置标题" prop="title">
                    <Input v-model="formData.title" placeholder="请输入配置标题"></Input>
                </FormItem>
                <FormItem label="配置key" prop="name">
                    <RadioGroup v-model="formData.name" v-if="formData.type=='1'">
                        <Radio label="qq">QQ</Radio>
                        <Radio label="twitter">twitter</Radio>
                        <Radio label="weibo">微博</Radio>
                        <Radio label="git">码云</Radio>
                        <Radio label="github">github</Radio>
                    </RadioGroup>
                    <RadioGroup v-model="formData.name" v-else-if="formData.type=='2'">
                        <Radio label="famous_remark">名言</Radio>
                        <Radio label="blog_related">博客相关</Radio>
                        <Radio label="blogger_collection">博主常逛</Radio>
                        <Radio label="copyright">版权信息</Radio>
                    </RadioGroup>
                    <Input v-model="formData.name" placeholder="请输入配置key" v-else></Input>
                </FormItem>
                <FormItem label="配置value" prop="val">
                    <Input v-model="formData.val" type="textarea" :rows="6"
                           placeholder="请输入配置value"></Input>
                </FormItem>
            </Form>
        </Modal>
    </card>
</template>

<script>
    import {addConfig, getList,updateConfig,deleteConfig} from "../api/webConfig";

    export default {
        name: 'webConfigPage',
        data() {
            const validateRadio = (rule, value, callback) => {
                if (this.formData.type == '') {
                    callback(new Error('请选择配置类型'))
                } else {
                    callback()
                }
            }
            const validateName = (rule, value, callback) => {
                if (this.formData.name == '') {
                    callback(new Error('请输入配置key'))
                } else {
                    callback()
                }
            }
            return {
                tableData: [],
                loading: false,
                type: [
                    {id: 1, title: '社交地址'},
                    {id: 2, title: 'footer内容'},
                    {id: 3, title: '其他配置'},
                ],
                button_loading: true,
                configModal: false,
                configUpdateModal:false,
                total: 0,
                inputStyle: {
                    width: '300px'
                },
                listData: {
                    type: '',
                    pageSize: 10,
                    pageNum: 1,
                },
                formData: {
                    title: '',
                    name: '',
                    val: '',
                    type: '',
                    id:''
                },
                ruleValidate: {
                    type: [
                        {validator: validateRadio, trigger: 'blur'}
                    ],
                    title: [
                        {required: true, message: '请输入配置标题', trigger: 'blur'},
                        {type: 'string', min: 2, max: 100, message: '配置标题在2到100个字符之间', trigger: 'change'},
                    ],
                    name: [
                        {validator: validateName, trigger: 'blur'}
                    ],
                    val: [
                        {required: true, message: '请输入配置val', trigger: 'blur'},
                        {type: 'string', min: 1, message: '配置val最少2个字符', trigger: 'change'},
                    ]
                },
                tableColumns: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '配置类型',
                        key: 'type',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.type === 1 ? 'primary' : row.type === 2 ? 'success' : 'error';
                            const text = row.type === 1 ? '社交地址' : row.type === 2 ? 'footer内容' : '其他配置';
                            return h('Tooltip', {
                                props: {
                                    content: text,
                                    placement: 'top'
                                }
                            }, [
                                h('Tag', {
                                    props: {
                                        type: 'dot',
                                        color: color
                                    }
                                }, text)
                            ]);
                        }
                    },
                    {
                        title: '配置名称',
                        key: 'title',
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.title;
                            return h('span', {}, text);
                        }
                    },
                    {
                        title: '配置key',
                        key: 'category',
                        tooltip: true,
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.name;
                            return h('span', {}, text);
                        }
                    },
                    {
                        title: '配置val',
                        key: 'val',
                        tooltip: true
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            let data = params.row;
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
                                            this.handleUpdateConfigModal(data);
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
                                            this.deleteConfig(data.id)
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
            modalCancel(name) {
                this.formData.name = '';
                this.$refs[name].resetFields();
            },
            changePage(page) {
                this.listData.pageNum = page;
                this.getConfigList();
            },
            //处理button的loading状态
            dealButtonLoading() {
                let _this = this
                setTimeout(function () {
                    _this.button_loading = false;
                    _this.$nextTick(() => {
                        _this.button_loading = true;
                    });
                }, 500);
            },
            handleSearch(){
                this.getConfigList();
            },
            handleConfig(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        addConfig(this.formData).then(res => {
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg)
                                this.configModal = false;
                                this.modalCancel(name)
                                this.getConfigList()
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg);
                            }
                        }).catch(err => {
                            this.dealButtonLoading();
                            this.$Message.error('服务器开小差了!')
                        })
                    } else {
                        this.dealButtonLoading();
                        this.$Message.error('验证失败!')
                    }
                })
            },
            getConfigList() {
                this.loading = true;
                getList(this.listData).then(res => {
                    this.loading = false;
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
                    this.loading = false;
                    this.$Message.error('服务器开小差了!')
                })
            },
            handleUpdateConfigModal(data){
                this.formData.id = data.id;
                this.formData.type = data.type.toString();
                this.formData.name = data.name;
                this.formData.val = data.val;
                this.formData.title = data.title;
                this.configUpdateModal = true;
            },
            updateConfig(name){
                updateConfig(this.formData).then(res=>{
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg)
                        this.configUpdateModal = false;
                        this.modalCancel(name)
                        this.getConfigList()
                    } else {
                        this.dealButtonLoading();
                        this.$Message.error(data.msg);
                    }
                }).catch(err=>{
                    this.dealButtonLoading();
                    this.$Message.error('服务器开小差了!');
                })
            },
            deleteConfig(id){
                this.loading = true;
                deleteConfig({id}).then(res=>{
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg)
                        this.getConfigList()
                    } else {
                        this.$Message.error(data.msg);
                    }
                }).catch(err=>{
                    this.loading = false;
                    this.$Message.error('服务器开小差了!');
                })
            }
        },
        created() {
            this.getConfigList()
        }
    }
</script>

<style scoped>
    .search-con {
        padding: 10px 0;
    }

    .link_add_button {
        float: right;
    }
</style>
