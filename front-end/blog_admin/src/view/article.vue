<template>
    <div>
        <card>
            <div class="search-con search-con-top">
                <Select v-model="listData.category_id" style="width:150px" clearable filterable placeholder="请选择文章分类">
                    <Option v-for="item in category" :value="item.id" :key="item.id">{{ item.title }}</Option>
                </Select>
                <Input clearable placeholder="输入文章名称搜索" class="hx_input"
                       v-model="listData.search"/>
                <Button @click="handleSearch" class="search-btn" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;搜索
                </Button>
                <Button @click="modalOpen" class="link_add_button" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;新增文章
                </Button>
            </div>
            <Table :data="tableData" :columns="tableColumns" :loading="loading" stripe></Table>
            <div style="margin: 10px;overflow: hidden">
                <div style="float: right;">
                    <Page :total="total" show-total :page-size="listData.pageSize" @on-change="changePage"></Page>
                </div>
            </div>
            <Modal v-model="articleModal"
                   :mask-closable="false"
                   :closable="false"
                   fullscreen
                   title="新增文章"
                   @on-ok="handleArticle('articleAdd')"
                   :loading="button_loading"
                   @on-cancel="modalCancel('articleAdd')"
            >
                <Form ref="articleAdd" :model="formData" :rules="ruleValidate" :label-width="80" inline>
                    <FormItem label="文章分类" style="width: 33%" prop="category_id">
                        <Select v-model="formData.category_id" :style="inputStyle" clearable filterable
                                placeholder="请选择文章分类">
                            <Option v-for="item in category" :value="item.id" :key="item.id">{{ item.title }}</Option>
                        </Select>
                    </FormItem>
                    <FormItem label="文章作者" style="width: 33%" prop="author">
                        <Input v-model="formData.author" placeholder="请输入文章作者" :style="inputStyle"></Input>
                    </FormItem>
                    <FormItem label="文章标题" style="width: 30%" prop="title">
                        <Input v-model="formData.title" placeholder="请输入文章标题" :style="inputStyle"></Input>
                    </FormItem>
                    <FormItem label="缩略图" style="width: 33%" prop="thumb">
                        <div class="demo-upload-list" v-for="item in uploadList">
                            <template v-if="item.status === 'finished'">
                                <img :src="item.url">
                                <div class="demo-upload-list-cover">
                                    <Icon type="ios-eye-outline" @click.native="handleView(item.name)"></Icon>
                                    <Icon type="ios-trash-outline" @click.native="handleRemove(item)"></Icon>
                                </div>
                            </template>
                            <template v-else>
                                <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
                            </template>
                        </div>
                        <Upload
                                ref="upload"
                                :show-upload-list="false"
                                :default-file-list="defaultImageList"
                                :on-success="handleSuccess"
                                :format="['jpg','jpeg','png','gif']"
                                :max-size="2048"
                                :on-format-error="handleFormatError"
                                :on-exceeded-size="handleMaxSize"
                                :before-upload="handleBeforeUpload"
                                :multiple="false"
                                type="drag"
                                :action="upload_url"
                                :headers="uploadHeader"
                                style="display: inline-block;width:58px;">
                            <div style="width: 58px;height:58px;line-height: 58px;">
                                <Icon type="ios-camera" size="20"></Icon>
                            </div>
                        </Upload>
                        <Modal title="缩略图" v-model="visible">
                            <img :src="imgName" v-if="visible" style="width: 100%">
                        </Modal>
                    </FormItem>
                    <FormItem label="文章标签" style="width: 33%" prop="tags">
                        <Select v-model="formData.tags" multiple clearable filterable :style="inputStyle" ref="tags_select">
                            <Option v-for="item in tags" :value="item.id" :key="item.id">{{ item.name }}</Option>
                        </Select>
                    </FormItem>
                    <FormItem label="关键词" style="width: 30%" prop="keywords">
                        <Input v-model="formData.keywords" value="admin" placeholder="请输入文章关键词"
                               :style="inputStyle"></Input>
                    </FormItem>
                    <FormItem label="是否原创" style="width: 33%">
                        <i-switch v-model="switchData.is_original" @on-change="switchOriginal"></i-switch>
                    </FormItem>
                    <FormItem label="是否置顶" style="width: 33%">
                        <i-switch v-model="switchData.is_top" @on-change="switchTop"></i-switch>
                    </FormItem>
                    <FormItem label="是否显示" style="width: 30%">
                        <i-switch v-model="switchData.is_show" @on-change="switchShow"></i-switch>
                    </FormItem>
                    <FormItem label="文章描述" style="width: 100%" prop="description">
                        <Input v-model="formData.description" type="textarea" :rows="6"
                               placeholder="请输入文章关键词,不输入默认文章前200个字" :style="inputStyle"></Input>
                    </FormItem>
                    <FormItem label="文章内容" style="width: 100%" prop="content_html">
                        <mavon-editor
                                v-model="formData.content_md"
                                style="min-height: 400px"
                                codeStyle="monokai-sublime"
                                ref=md
                                @imgAdd="imgAdd"
                                @change="contentChange"
                        />
                    </FormItem>
                </Form>
            </Modal>
        </card>
    </div>
</template>

<script>
    import {getArticleList, uploadImage, addArticle, updateArticle, deleteArticle} from '../api/article'
    import config from '@/config'
    import {mavonEditor} from 'mavon-editor'
    import 'mavon-editor/dist/css/index.css'
    import {getToken} from "../libs/util";


    const {baseUrl, imageUrl,uploadUrl} = config
    export default {
        name: 'articlePage',
        components: {
            mavonEditor
        },
        data() {
            const validateCategory = (rule, value, callback) => {
                if (value == '') {
                    callback(new Error('请选择文章分类'))
                } else {
                    callback()
                }
            }
            const validateTag = (rule, value, callback) => {
                if (value.length == 0) {
                    callback(new Error('请选择文章标签'))
                } else {
                    callback()
                }
            }
            return {
                uploadHeader:{Authorization:getToken()},
                tableData: [],
                loading: false,
                category: [],
                button_loading: true,
                tags: [],
                articleModal: false,
                total: 0,
                inputStyle: {
                    width: '500px'
                },
                switchData:{
                    is_show:true,
                    is_top:false,
                    is_original:true
                },
                listData: {
                    category_id: '',
                    search: '',
                    pageSize: 10,
                    pageNum: 1,
                },
                formData: {
                    id: '',
                    category_id: '',
                    author: 'admin',
                    title: '',
                    thumb: '',
                    tags: [],
                    keywords: '',
                    description: '',
                    content_html: '',
                    content_md: '',
                    is_show:1,
                    is_original:1,
                    is_top:0
                },
                defaultImageList: [],
                imgName: '',
                visible: false,
                uploadList: [],
                upload_url: '',
                ruleValidate: {
                    category_id: [
                        {validator: validateCategory, trigger: 'blur'},
                    ],
                    author: [
                        {required: true, message: '请输入文章作者', trigger: 'blur'},
                        {type: 'string', min: 2, max: 20, message: '作者名字在2到20个字符之间', trigger: 'change'},
                    ],
                    title: [
                        {required: true, message: '请输入文章标题', trigger: 'blur'},
                        {type: 'string', min: 2, max: 100, message: '文章标题在2到100个字符之间', trigger: 'change'},
                    ],
                    thumb: [
                        {required: true, message: '请上传文章缩略图', trigger: 'blur'},
                        {type: 'url', message: '文章缩略图格式不正确', trigger: 'blur'},
                    ],
                    tags: [
                        {validator: validateTag, trigger: 'blur'},
                    ],
                    keywords: [
                        {required: true, message: '请输入文章关键词', trigger: 'blur'},
                        {type: 'string', min: 2, max: 30, message: '文章关键词在2到30个字符之间', trigger: 'change'},
                    ],
                    content_html: [
                        {required: true, message: '请输入文章内容', trigger: 'blur'},
                        {type: 'string', min: 2, message: '文章内容最少2个字符', trigger: 'change'},
                    ],
                },
                tableColumns: [
                    {
                        title: 'ID',
                        key: 'id'
                    },
                    {
                        title: '缩略图',
                        key: 'thumb',
                        render: (h, params) => {
                            const row = params.row;
                            return h('Tooltip', {
                                props: {
                                    placement: 'top',
                                    theme:'light'
                                },
                                style:{
                                    cursor:'pointer'
                                }
                            }, [
                                h('img', {
                                    style: {
                                        width: "50px",
                                        verticalAlign: "middle"
                                    },
                                    attrs: {
                                        src: row.thumb
                                    }
                                },),
                                h('img', {
                                    style: {
                                        width: "200px",
                                        verticalAlign: "middle"
                                    },
                                    slot:'content',
                                    attrs: {
                                        src: row.thumb
                                    }
                                },)
                            ])

                        }
                    },
                    {
                        title: '作者',
                        key: 'author',
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.author;
                            return h('span', {}, text);
                        }
                    },
                    {
                        title: '分类',
                        key: 'category',
                        tooltip: true,
                        render: (h, params) => {
                            const row = params.row;
                            const text = row.category.title;
                            return h('span', {}, text);
                        }
                    },
                    {
                        title: '标题',
                        key: 'title',
                        tooltip: true
                    },
                    {
                        title: '描述',
                        key: 'description',
                        tooltip: true
                    },
                    {
                        title: '标签',
                        key: 'tags',
                        render: (h, params) => {
                            let data = params.row;
                            return h('Poptip', {
                                props: {
                                    trigger: 'hover',
                                    title: data.tags.length + '个标签',
                                    placement: 'top'
                                }
                            }, [
                                h('Tag', data.tags.length),
                                h('div', {
                                    slot: 'content'
                                }, [
                                    h('ul', this.tableData[params.index].tags.map(item => {
                                        return h('li', {
                                            style: {
                                                textAlign: 'center',
                                                padding: '4px'
                                            }
                                        }, item.name)
                                    }))
                                ])
                            ]);
                        }
                    },
                    {
                        title: '点击数',
                        key: 'click',
                        sortable: true
                    },
                    {
                        title: '评论数',
                        key: 'comments_count',
                        sortable: true
                    },
                    {
                        title: '原创',
                        key: 'is_original',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.is_original === 1 ? 'success' : 'error';
                            const text = row.is_original === 1 ? '√' : '×';

                            return h('Tag', {
                                props: {
                                    color: color
                                }
                            }, text);
                        }
                    },
                    {
                        title: '置顶',
                        key: 'is_top',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.is_top === 1 ? 'success' : 'error';
                            const text = row.is_top === 1 ? '√' : '×';

                            return h('Tag', {
                                props: {
                                    color: color
                                }
                            }, text);
                        }
                    },
                    {
                        title: '显示',
                        key: 'is_show',
                        render: (h, params) => {
                            const row = params.row;
                            const color = row.is_show === 1 ? 'success' : 'error';
                            const text = row.is_show === 1 ? '√' : '×';

                            return h('Tag', {
                                props: {
                                    color: color
                                }
                            }, text);
                        }
                    },
                    {
                        title: '创建时间',
                        key: 'created_at',
                        sortable: true,
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
                                            this.handleUpdateArticleData(data);
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
                                            this.deleteArticle(data.id)
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
            switchShow(value){
                this.formData.is_show = value?1:0
            },
            switchTop(value){
                this.formData.is_top = value?1:0
            },
            switchOriginal(value){
                this.formData.is_original = value?1:0
            },
            modalCancel(name) {
                this.$refs.tags_select.clearSingleSelect();
                this.formData.description = '';
                this.formData.id = '';
                this.formData.content_html = '';
                this.formData.content_md = '';
                this.handleIviewUploadDefaultListBug();
                this.$refs[name].resetFields();
            },
            modalOpen(){
                this.$refs.tags_select.clearSingleSelect();
                this.formData.is_original = 1;
                this.formData.is_top = 0;
                this.formData.is_show = 1;
                this.switchData.is_original = true;
                this.switchData.is_top = false;
                this.switchData.is_show = true;
                this.articleModal = true;
            },
            changePage(page) {
                this.listData.pageNum = page;
                this.getArticleList();
            },
            handleSearch() {
                this.listData.pageNum = 1;
                this.getArticleList();
            },
            handleView(name) {
                this.imgName = imageUrl + name;
                this.visible = true;
            },
            handleRemove(file) {
                const fileList = this.$refs.upload.fileList;
                this.$refs.upload.fileList.splice(fileList.indexOf(file), 1);
            },
            //获取转换后的html内容
            contentChange(value, render) {
                this.formData.content_html = render;
            },
            //文章图片上传
            imgAdd(pos, $file) {
                let formdata = new FormData();
                formdata.append('file', $file);
                let token = getToken();
                uploadImage({formdata,token}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$refs.md.$img2Url(pos, imageUrl + data.data);
                    } else {
                        this.$Message.error(data.msg);
                    }
                })
            },
            //处理图片上传的结果
            handleSuccess(res, file) {
                if (res.code == 200) {
                    file.url = imageUrl + res.data;
                    file.name = res.data;
                    this.formData.thumb = imageUrl + res.data;
                } else {
                    this.$Message.error(res.msg)
                }
            },
            handleFormatError(file) {
                this.$Message.error(file.name + '格式错误，允许的格式有jpg、jpeg、png、gif');
            },
            handleMaxSize(file) {
                this.$Message.error('上传文件过大，最多2M！');
            },
            handleBeforeUpload() {
                this.uploadList.splice(0, 1);
                const check = this.uploadList.length < 2;
                if (!check) {
                    this.$Message.error('最多上传一张照片')
                }
                return check;
            },
            //处理iview的默认defaultlist bug
            handleIviewUploadDefaultListBug(data=null){
                if(data==null){
                    data = [];
                }
                setTimeout(()=> {
                    this.defaultImageList = data; //将获取到的值赋值到 defaultList 中
                    this.$nextTick(()=> { //赋值后马上更新
                        this.uploadList = this.$refs.upload.fileList;
                    });
                },500);
            },
            //处理编辑文章的数据
            handleUpdateArticleData(data) {
                this.handleIviewUploadDefaultListBug([
                    {
                        'name': 'bc7521e033abdd1e92222d733590f104',
                        'url': data.thumb
                    }
                ]);
                this.formData.id = data.id;
                this.formData.category_id = data.category_id;
                this.formData.author = data.author;
                this.formData.title = data.title;
                this.formData.content_md = data.content_md ? data.content_md : data.content_html;
                this.formData.content_html = data.content_html;
                this.formData.description = data.description;
                this.formData.keywords = data.keywords;
                this.formData.thumb = data.thumb;
                this.formData.is_original = data.is_original;
                this.formData.is_top = data.is_top;
                this.formData.is_show = data.is_show;
                this.switchData.is_original = data.is_original==1?true:false;
                this.switchData.is_top = data.is_top==1?true:false;
                this.switchData.is_show = data.is_show==1?true:false;
                for (let val of data.tags) {
                    this.formData.tags.push(val.id)
                }
                this.articleModal = true;
            },
            //获取文章列表
            getArticleList() {
                this.loading = true;
                getArticleList(this.listData).then(res => {
                    this.loading = false;
                    let data = res.data;
                    if (data.code == 200) {
                        this.tableData = data.data.list;
                        this.total = data.data.total;
                        this.category = data.data.category;
                        this.tags = data.data.tag;
                    } else {
                        this.tableData = [];
                        this.total = 0;
                        this.$Message.error(data.msg);
                    }
                }).catch(err=>{
                    this.loading = false;
                    this.tableData = [];
                    this.total = 0;
                    this.$Message.error('服务器错误');
                })
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
            handleArticle(name) {
                this.$refs[name].validate((valid) => {
                    if (valid) {
                        let pro = '';
                        if(this.formData.id!='' && !isNaN(this.formData.id)){
                            pro = updateArticle(this.formData);
                        }else {
                            delete this.formData.id;
                            pro = addArticle(this.formData);
                        }
                        pro.then(res => {
                            let data = res.data;
                            if (data.code == 200) {
                                this.articleModal = false;
                                this.$Message.success(data.msg);
                                this.getArticleList();
                            } else {
                                this.dealButtonLoading();
                                this.$Message.error(data.msg);
                            }
                        })
                    } else {
                        this.dealButtonLoading();
                        this.$Message.error('验证失败!')
                    }
                })
            },
            deleteArticle(id) {
                this.loading = true;
                deleteArticle({id}).then(res => {
                    let data = res.data;
                    if (data.code == 200) {
                        this.$Message.success(data.msg);
                        this.getArticleList();
                    } else {
                        this.loading = false;
                        this.$Message.error(data.msg);
                    }
                })
            }
        },
        created() {
            this.upload_url = uploadUrl;
            this.getArticleList();
        },
        mounted() {
            this.uploadList = this.$refs.upload.fileList;
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

    .link_add_button {
        float: right;
    }

    .demo-upload-list {
        display: inline-block;
        width: 60px;
        height: 60px;
        text-align: center;
        line-height: 60px;
        border: 1px solid transparent;
        border-radius: 4px;
        overflow: hidden;
        background: #fff;
        position: relative;
        box-shadow: 0 1px 1px rgba(0, 0, 0, .2);
        margin-right: 4px;
    }

    .demo-upload-list img {
        width: 100%;
        height: 100%;
    }

    .demo-upload-list-cover {
        display: none;
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        background: rgba(0, 0, 0, .6);
    }

    .demo-upload-list:hover .demo-upload-list-cover {
        display: block;
    }

    .demo-upload-list-cover i {
        color: #fff;
        font-size: 20px;
        cursor: pointer;
        margin: 0 2px;
    }
</style>
