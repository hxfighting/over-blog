<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;

class UploadController extends Controller
{
    //允许上传的图片格式
    private $allow_image = ['png','jpg','jpeg','gif'];
    //图片最大尺寸2M
    private $max_size = 2048;


    //图片上传
    public function upload(Request $request)
    {
        $this->validate($request,['file'=>'required|image']);
        $file = $request->file('file');
        if($file->isValid()){
            $ext = strtolower($file->getClientOriginalExtension());
            if(!in_array($ext,$this->allow_image,true)){
                return renderError('上传图片允许的格式只能是:gif、png、jpg、jpeg！');
            }
            $file_size = $file->getSize()/1024;
            if($file_size>$this->max_size){
                return renderError('上传图片最多2M！');
            }
            $disk = Storage::disk('qiniu');
            $image_key = md5(file_get_contents($file->getPathname()));
            if($disk->exists($image_key)){
                return renderSuccess('上传图片成功', $image_key);
            }else{
                $picture_path = $disk->putFileAs('',$file,$image_key);
                return $picture_path ? renderSuccess('上传图片成功', $picture_path)
                    : renderError('上传图片失败,请稍后再试');
            }
        }else{
            return renderError('上传图片失败,请稍后再试');
        }
    }
}
