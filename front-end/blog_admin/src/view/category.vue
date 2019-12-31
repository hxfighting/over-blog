<template>
    <div>
        <card>
            <div class="search-con search-con-top">
                <Button @click="addCategory(0)" class="search-btn" type="primary">
                    <Icon type="search"/>&nbsp;&nbsp;新增一级分类
                </Button>
            </div>
            <Tree :data="categoryData" :render="renderContent"></Tree>
            <div v-if="loading" class="demo-spin-container">
                <Spin fix></Spin>
            </div>
        </card>
    </div>
</template>

<script>
    import {getCategoryList,addCategory,updateCategory,deleteCategory} from '../api/category'
    export default {
        name: 'category',
        data() {
            return {
                loading:false,
                addLoading:false,
                categoryData: [],
                title:'',
                buttonProps: {
                    type: 'default',
                    size: 'small',
                }
            }
        },
        methods: {
            renderContent (h, { root, node, data }) {
                if(data.pid===0 && data.type===1){
                    return h('span', {
                        style: {
                            display: 'inline-block',
                            width: '100%'
                        }
                    }, [
                        h('span', [
                            h('Icon', {
                                props: {
                                    type: 'ios-paper-outline'
                                },
                                style: {
                                    marginRight: '8px'
                                }
                            }),
                            h('span', data.title)
                        ]),
                        h('span', {
                            style: {
                                display: 'inline-block',
                                float: 'right',
                                marginRight: '32px'
                            }
                        }, [
                            h('Button', {
                                style: {
                                    marginRight: '2px'
                                },
                                props: {
                                    type: 'primary',
                                    size: 'small',
                                    icon: 'md-brush'
                                },
                                attrs: {
                                    title: '编辑分类'
                                },
                                on: {
                                    click: () => {
                                        this.updateCategory(data)
                                    }
                                }
                            }),
                            h('Button', {
                                style: {
                                    marginRight: '2px'
                                },
                                attrs: {
                                    title: '添加子分类'
                                },
                                props: Object.assign({}, this.buttonProps, {
                                    icon: 'ios-add',
                                    type: 'success'
                                }),
                                on: {
                                    click: () => {
                                        this.addCategory(data.id)
                                    }
                                }
                            }),
                            h('Poptip', {
                                props: {
                                    confirm: true,
                                    title: '你确定要删除吗?'
                                },
                                attrs: {
                                    title: '删除分类'
                                },
                                on: {
                                    'on-ok': () => {
                                        this.remove(data)
                                    }
                                }
                            }, [
                                h('Button', {
                                    props: {
                                        type: 'error',
                                        size: 'small',
                                        icon: 'ios-remove'
                                    }
                                })
                            ])
                        ])
                    ]);
                }else {
                    return h('span', {
                        style: {
                            display: 'inline-block',
                            width: '100%'
                        }
                    }, [
                        h('span', [
                            h('Icon', {
                                props: {
                                    type: 'ios-paper-outline'
                                },
                                style: {
                                    marginRight: '8px'
                                }
                            }),
                            h('span', data.title)
                        ]),
                        h('span', {
                            style: {
                                display: 'inline-block',
                                float: 'right',
                                marginRight: '32px'
                            },
                        }, [
                            h('Button', {
                                style: {
                                    marginRight: '2px'
                                },
                                props: {
                                    type: 'primary',
                                    size: 'small',
                                    icon: 'md-brush'
                                },
                                on: {
                                    click: () => {
                                        this.updateCategory(data)
                                    }
                                }
                            }),
                            h('Poptip', {
                                props: {
                                    confirm: true,
                                    title: '你确定要删除吗?'
                                },
                                attrs: {
                                    title: '删除分类'
                                },
                                on: {
                                    'on-ok': () => {
                                        this.remove(data)
                                    }
                                }
                            }, [
                                h('Button', {
                                    props: {
                                        type: 'error',
                                        size: 'small',
                                        icon: 'ios-remove'
                                    }
                                })
                            ])
                        ])
                    ]);
                }
            },
            append (data) {
                const children = data.children || [];
                children.push({
                    title: 'appended node',
                    expand: true
                });
                this.$set(data, 'children', children);
            },
            remove (data) {
                deleteCategory({id:data.id}).then(res=>{
                    let re = res.data;
                    if(re.code===200){
                        this.$Message.success(re.msg);
                        this.getCategoryList();
                    }else {
                        this.$Message.error(re.msg)
                    }
                })
            },
            getCategoryList(){
                this.loading = true;
                getCategoryList().then(res=>{
                    this.loading = false;
                    let data = res.data;
                    if(data.code===200){
                        this.categoryData = data.data;
                    }else {
                        this.$Message.error(data.msg)
                    }
                })
            },
            addCategory(pid){
                this.$Modal.confirm({
                    loading: this.addLoading,
                    render: (h) => {
                        return h('Input', {
                            props: {
                                autofocus: true,
                                placeholder: '请输入分类名称'
                            },
                            on: {
                                input: (val) => {
                                    this.title = val;
                                }
                            }
                        })
                    },
                    onOk: () => {
                        if(this.title==='' || this.title===undefined){
                            this.$Message.error('请输入分类名称!');
                            return;
                        }
                        this.addLoading = true;
                        addCategory({pid: pid, title: this.title}).then(res => {
                            this.addLoading = false;
                            this.$Modal.remove();
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg);
                                this.getCategoryList();
                            } else {
                                this.$Message.error(data.msg);
                            }
                        })
                    }
                })
            },
            updateCategory(data){
                this.title = data.title
                this.$Modal.confirm({
                    loading: this.addLoading,
                    render: (h) => {
                        return h('Input', {
                            props: {
                                value: this.title,
                                autofocus: true,
                                placeholder: '请输入分类名称'
                            },
                            on: {
                                input: (val) => {
                                    this.title = val;
                                }
                            }
                        })
                    },
                    onOk: () => {
                        if(this.title==='' || this.title===undefined){
                            this.$Message.error('请输入分类名称!');
                            return;
                        }
                        this.addLoading = true;
                        updateCategory({id: data.id, title: this.title}).then(res => {
                            this.addLoading = false;
                            this.$Modal.remove();
                            let data = res.data;
                            if (data.code == 200) {
                                this.$Message.success(data.msg);
                                this.getCategoryList();
                            } else {
                                this.$Message.error(data.msg);
                            }
                        })
                    }
                })
            }
        },
        created() {
            this.getCategoryList();
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
