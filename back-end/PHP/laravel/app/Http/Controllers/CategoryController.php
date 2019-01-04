<?php

namespace App\Http\Controllers;

use App\Http\Models\Category;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

class CategoryController extends Controller
{
    //文章分类type
    private const ARTICLE_TYPE = 1;

    //获取分类列表
    public function list(Category $category): JsonResponse
    {
        $list = $category->getTree();
        if ($list->isNotEmpty())
        {
            foreach ($list as $item)
            {
                $item->expand = true;
            }
            return renderSuccess('获取分类列表成功', $list);
        }
        return renderError('暂无分类列表数据');
    }

    //删除分类
    public function destroy(Request $request, Category $category): JsonResponse
    {
        $data = $this->validate($request, ['id' => 'required|integer']);
        $exist_category = $category->find($data['id']);
        if ($exist_category->type != self::ARTICLE_TYPE)
        {
            return renderError('不能删除此分类');
        }
        $res = $exist_category->delete();
        return $res ? renderSuccess('删除分类成功')
            : renderError('删除分类失败,请稍后再试！');
    }

    //添加分类
    public function store(Request $request, Category $category): JsonResponse
    {
        $data = $this->validate($request, ['pid' => 'required|integer|min:0', 'title' => 'required|min:2|max:20|unique:category,title']);
        $category->pid = $data['pid'];
        $category->title = $data['title'];
        $res = $category->save();
        return $res ? renderSuccess('添加分类成功')
            : renderError('添加分类失败,请稍后再试！');
    }

    //修改分类
    public function update(Request $request, Category $category): JsonResponse
    {
        $id = $request->id;
        $data = $this->validate($request, ['id' => 'required|integer', 'title' => 'required|min:2|max:20|unique:category,title,' . $id]);
        $exist_category = $category->find($id);
        $exist_category->title = $data['title'];
        $res = $exist_category->save();
        return $res ? renderSuccess('修改分类成功')
            : renderError('修改分类失败,请稍后再试！');
    }
}
